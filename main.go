package main

// KEY: D9E597D88058F1EFA44434531CC57183
// https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=D9E597D88058F1EFA44434531CC57183&steamid=76561198096213883&include_appinfo=true&format=json
// https://cdn.cloudflare.steamstatic.com/steam/apps/3147870/library_600x900_2x.jpg

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed views/*
var views embed.FS

var t = template.Must(template.ParseFS(views, "views/*"))

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("./games.json")
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}

		data := Games{}
		err = json.Unmarshal(content, &data)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}
		if err := t.ExecuteTemplate(w, "index.html", data); err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	})

	server := http.Server{
		Addr:    ":80",
		Handler: router,
	}

	fmt.Println("Listening on 8000")
	server.ListenAndServe()
}
