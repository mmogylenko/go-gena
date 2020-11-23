package main

import (
	"context"
	"log"
	"net/http"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
)

func main() {
	ctx, cancel = context.WithCancel(context.Background())
	http.HandleFunc("/calculate", CalculateHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
