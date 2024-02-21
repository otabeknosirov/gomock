package main

import (
	router2 "go-mock/router"
	"net/http"
)

func main() {
	router := router2.Router()

	http.ListenAndServe(":8081", router)
}
