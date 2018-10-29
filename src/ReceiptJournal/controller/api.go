package controller

import (
	"ReceiptJournal/viewmodel"
	"encoding/json"
	"net/http"
)

func registerAPIRoutes() {
	http.HandleFunc("/api/cashflow", handleCashflowAPI)
}

func handleCashflowAPI(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		pusher.Push("/css/app.css", &http.PushOptions{
			Header: http.Header{"Content-Type": []string{"text/css"}},
		})
	}

	data := viewmodel.NewCashflow()
	json.NewEncoder(w).Encode(data)
	w.Header().Add("Content-type", "text/html")
}
