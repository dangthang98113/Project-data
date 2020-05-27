package main

import (
	"fmt"
	"log"
	"net/http"

	router "./src/router"
)

func main() {
	fmt.Println("server start...!")
	log.Fatal(http.ListenAndServe(":8080", router.Router()))
}
