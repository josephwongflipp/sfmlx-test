package main

import (
	"fmt"
	"net/http"

	"github.com/wishabi/sfmlx-test/handlers"
)

func main() {
	fmt.Println("Starting up sfmlx-service...")
	http.HandleFunc("/api/v1/sfmlx", handlers.GetSFMLX)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
