package payloads

type CreateGameParam struct {
	Id      int    `json:"id"`
	Mdate   string `json:"mdate"`
	Stadium string `json:"stadium"`
	Team1   string `json:"team1"`
	Team2   string `json:"team2"`
}
