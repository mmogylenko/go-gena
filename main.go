package main

import (
	"context"
	"log"
	"math"
	"net/http"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	x      float64
)

func main() {
	x = 0.001
	ctx, cancel = context.WithCancel(context.Background())
	http.HandleFunc("/calculate", start)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func calculate() float64 {
	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}
	return x
}
