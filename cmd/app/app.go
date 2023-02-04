package main

import (
	"fmt"
	"net/http"
	"os"
	"universal-search/internal/handlers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	http.HandleFunc("/search", handlers.SearchHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
