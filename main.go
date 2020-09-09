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
)

func main() {
	ctx, cancel = context.WithCancel(context.Background())
	http.HandleFunc("/start", start)
	http.HandleFunc("/stop", stop)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func load(ctx context.Context) {
	done := false
	x := 0.0001
	go func() {
		// Keep doing math.
		for i := 0; !done; i++ {
			x += math.Sqrt(x)
		}
	}()

	// Wait util context is cancelled.
	<-ctx.Done()
	// Turn on closing flag.
	done = true
}
