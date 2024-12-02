package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	HOST := os.Getenv("host")
	PORT := os.Getenv("port")

	addr := fmt.Sprint(HOST, ":", PORT)

	log.Print(HOST, PORT)
	
	log.Fatal(http.ListenAndServe(addr, nil))
}
