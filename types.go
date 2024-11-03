package main

type Games struct {
	Game []Game `json:"games"`
}

type Game struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	BoxArt      string `json:"boxArt"`
	Link        string `json:"link"`
	IsCompleted bool   `json:"isCompleted"`
	Date        string `json:"completedOn"`
	Rating      int    `json:"rating"`
}
