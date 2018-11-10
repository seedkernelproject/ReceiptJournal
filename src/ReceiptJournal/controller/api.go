package controller

import (
	"ReceiptJournal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func registerAPIRoutes() {
	http.HandleFunc("/api/receipts", handleReceipts)
	http.HandleFunc("/api/createreceipt", handleCreateReceipt)
	http.HandleFunc("/api/updatereceipt", handleUpdateReceipt)
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

func handleCreateReceipt(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var createdreceipt model.Receipt
		err := decoder.Decode(&createdreceipt)
		if err != nil {
			panic(err)
		}
		createdreceipt.InputDate = time.Now().String()
		fmt.Println(createdreceipt)
		model.SaveReceipt(createdreceipt)
	}
}

func handleUpdateReceipt(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var retrievedreceipt model.Receipt
		err := decoder.Decode(&retrievedreceipt)
		if err != nil {
			panic(err)
		}
		model.UpdateReceiptPaymentStatus(retrievedreceipt)
	}
}
