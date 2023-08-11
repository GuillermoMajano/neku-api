package main

import (
	"log"
	"net/http"
)

func server() {

	r := http.NewServeMux()

	r.HandleFunc("/", GetTags)

	http.ListenAndServe(":8080", r)

	log.Println("run server!!")

}
