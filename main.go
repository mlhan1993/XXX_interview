package main

import (
	"fmt"
	"github.com/mlhan1993/league_interview/api"
	"github.com/mlhan1993/league_interview/pkg/matrix"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	fmt.Println("service started at port 8080")
	processor := matrix.Processor{}
	handlers := api.NewMatrixHandlers(&processor)
	http.HandleFunc("/echo", handlers.Echo)
	http.HandleFunc("/invert", handlers.Invert)
	http.HandleFunc("/flatten", handlers.Flatten)
	http.HandleFunc("/sum", handlers.Sum)
	http.HandleFunc("/multiply", handlers.Multiply)
	http.ListenAndServe(":8080", nil)
}
