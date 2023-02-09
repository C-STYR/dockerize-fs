package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../client/dist"))

	log.Println("server running!")
	log.Fatal(http.ListenAndServe(":9000", fs))
}
