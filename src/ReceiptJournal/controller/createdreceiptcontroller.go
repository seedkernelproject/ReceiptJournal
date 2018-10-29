package controller

import (
	"ReceiptJournal/model"
	"ReceiptJournal/viewmodel"
	"html/template"
	"log"
	"net/http"
)

type createdreceipt struct {
	createdreceiptTemplate *template.Template
}

func (c createdreceipt) registerRoutes() {
	http.HandleFunc("/createdreceipt", c.handleCreatedReceipt)
}

func (c createdreceipt) handleCreatedReceipt(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		pusher.Push("/css/app.css", &http.PushOptions{
			Header: http.Header{"Content-Type": []string{"text/css"}},
		})
	}
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/newreceipt", http.StatusMovedPermanently)
		model.SaveReceipt(CreatedReceipt)
		log.Println(CreatedReceipt)
	}

	vm := viewmodel.NewCreatedReceipt(CreatedReceipt)
	w.Header().Add("Content-type", "text/html")
	c.createdreceiptTemplate.Execute(w, vm)
}
