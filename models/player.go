package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Player struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name    string    `gorm:"check:name,LENGTH(name)<15" json:"name"`
	Country string    `gorm:"check:country,LENGTH(country)=2" json:"country"`
	Score   int       `gorm:"score" json:"score"`
	Rank    int       `gorm:"-" json:"rank"`
}

func (p *Player) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return nil
}
