package controller

import (
	"ReceiptJournal/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func registerAPIRoutes() {
	http.HandleFunc("/api/receipts", handleReceipts)
}

func handleReceipts(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		pusher.Push("/css/app.css", &http.PushOptions{
			Header: http.Header{"Content-Type": []string{"text/css"}},
		})
	}

	data := model.GetReceipts()
	//json.NewEncoder(w).Encode(data)
	dataJson, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-type", "application/json")
	w.Write(dataJson)
	fmt.Println("called")
}
