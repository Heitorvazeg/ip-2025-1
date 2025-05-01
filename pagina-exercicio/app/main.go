package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileserver := http.Serverfile(http.dir(""))
}