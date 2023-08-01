package controllers

type CreatePlayerRequest struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Score   int    `json:"score"`
}

type UpdatePlayerReq struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
