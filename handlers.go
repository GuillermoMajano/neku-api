package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetTags(w http.ResponseWriter, r *http.Request) {

	rg, err := http.Get("https://api.jikan.moe/v4/manga/1")

	if err != nil {
		log.Fatal("a error has ocurred")
	}
	var getTags Tags
	err = json.NewDecoder(rg.Body).Decode(&getTags)

	if err != nil {
		log.Fatal("error decoding data")
	}

	tj, _ := json.Marshal(getTags)

	w.Header().Add("content-type", "application/json")

	w.Write(tj)
}

func insertManga(w http.ResponseController, r *http.Request) {

}
