package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mmogylenko/flexmessage"
)

// CalculateHandler is One-shot math operation (square root)
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	var notify flexmessage.FlexMessage
	w.Header().Add("Content-Type", "application/json")

	result := SquareRoot()
	notify.Message(fmt.Sprintf("%f", result))

	_ = json.NewEncoder(w).Encode(notify.Compact())
}
