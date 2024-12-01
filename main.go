package main

import (
	"log"
	"net/http"
)

func main() {

	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}
