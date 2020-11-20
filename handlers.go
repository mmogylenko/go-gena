package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mmogylenko/flexmessage"
)

func start(w http.ResponseWriter, r *http.Request) {
	var notify flexmessage.FlexMessage
	w.Header().Add("Content-Type", "application/json")

	result := calculate()
	notify.Message(fmt.Sprintf("%f", result))

	_ = json.NewEncoder(w).Encode(notify.Compact())
}
