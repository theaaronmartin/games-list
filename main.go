package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

//go:embed views
var views embed.FS

var t = template.Must(template.ParseFS(views, "views/*"))

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("./static/json/games.json")
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}

		data := Games{}
		err = json.Unmarshal(content, &data)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}
		if err := t.ExecuteTemplate(w, "index.gohtml", data); err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	})

	router.HandleFunc("GET /{id}", func(w http.ResponseWriter, req *http.Request) {
		content, err := os.ReadFile("./static/json/games.json")
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}
		idString := req.PathValue("id")
		num, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		data := Games{}
		err = json.Unmarshal(content, &data)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}
		if err := t.ExecuteTemplate(w, "game.gohtml", data.Game[num-1]); err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	})

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	server.ListenAndServe()
}
