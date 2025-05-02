package main

import (
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("../static"))

	http.Handle("/", fileserver)

	if err := http.ListenAndServe(":8081", fileserver); err != nil {
		log.Fatal(err)
	}
}
