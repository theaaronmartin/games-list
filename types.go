package main

// type Steam struct {
// 	Count int    `json:"game_count"`
// 	Games []Game `json:"games"`
// }

// type Game struct {
// 	AppId int    `json:"appid"`
// 	Name  string `json:"name"`
// }

// type Response struct {
// 	Steam Steam `json:"response"`
// }

type Games struct {
	Game []Game `json:"games"`
}

type Game struct {
	Name        string `json:"name"`
	BoxArt      string `json:"boxArt"`
	Link        string `json:"link"`
	IsCompleted bool   `json:"isCompleted"`
}
