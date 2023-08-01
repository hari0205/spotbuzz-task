package controllers

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hari0205/spotbuzz-task/models"
	"github.com/hari0205/spotbuzz-task/setup"
	"gorm.io/gorm"
)

func CreatePlayer(ctx *gin.Context) {

	var createplayer CreatePlayerRequest
	if err := ctx.ShouldBindJSON(&createplayer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Bad input. Please check your input and try again",
		})
		return
	}

	tx := setup.GormDB.Create(&models.Player{
		Name:    createplayer.Name,
		Country: createplayer.Country,
		Score:   createplayer.Score,
	})

	if tx.Error != nil && tx.RowsAffected != 1 {
		var errmessage string
		if strings.Contains(tx.Error.Error(), "name") {
			errmessage = fmt.Sprintf("Name given must not have more than %d characters", 15)
		} else {
			errmessage = fmt.Sprintf("Country must exactly have %d characters", 2)
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error creating player",
			"error":   errmessage,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Player created",
		"data":    createplayer,
	})

}

func UpdatePlayer(ctx *gin.Context) {

	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Bad input. Please check id parameter and try again.",
		})
		return
	}

	var updateBody UpdatePlayerReq
	if err := ctx.ShouldBindJSON(&updateBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad input. Check inputs and try again.",
		})
		return

	}

	if err := setup.GormDB.Model(&models.Player{}).Where("id = ?", id).First(&models.Player{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Record not found.",
			})
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal server error.",
			})
			return
		}
	}

	updateprop := make(map[string]interface{})
	if updateBody.Name != "" {
		updateprop["name"] = updateBody.Name
	}
	if updateBody.Score != 0 {
		updateprop["score"] = updateBody.Score
	}

	res := setup.GormDB.Model(&models.Player{}).Where("id = ?", id).Updates(&updateprop)
	if res.Error != nil {

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   res.Error.Error(),
			"message": "Error updating player",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Player Updated!",
		"data":    res,
	})

}

func GetPlayers(ctx *gin.Context) {

	var players []models.Player

	// Descending order of players based on score and name.
	res := setup.GormDB.Order("score DESC").Order("name DESC").Find(&players)

	if res.RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No Players were found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": players,
	})
}

func GetRandomPlayers(ctx *gin.Context) {
	var player models.Player

	var count int64
	setup.GormDB.Model(&models.Player{}).Count(&count)

	if count == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No players found in the database",
		})
		return
	}

	// Offsetting to a number from less than total players
	offset := rand.Intn(int(count))
	setup.GormDB.Offset(offset).Limit(1).Find(&player)
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   player,
	})

}

func GetPlayerByRank(ctx *gin.Context) {
	val := ctx.Param("val")
	rank, err := strconv.Atoi(val)
	if err != nil || rank <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Invalid rank",
		})
		return
	}

	var player models.Player
	// Get score in descending order . First player has 1st rank. Second player has second rank... Offset to find the required rank
	if err := setup.GormDB.Order("score desc").Offset(rank - 1).First(&player).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Player not found",
		})
		return
	}

	player.Rank = rank
	ctx.JSON(http.StatusOK, player)
}
