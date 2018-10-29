package controller

import (
	"ReceiptJournal/model"
	"ReceiptJournal/viewmodel"
	"html/template"
	"net/http"
)

type clientReceipt struct {
	clientReceiptTemplate          *template.Template
	retrievedClientReceiptTemplate *template.Template
}

func (c clientReceipt) registerRoutes() {
	http.HandleFunc("/clientreceipt", c.handleClientReceipt)
	http.HandleFunc("/clientreceipt/", c.handleClientReceiptDetails)
}

func (c clientReceipt) handleClientReceipt(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		pusher.Push("/css/app.css", &http.PushOptions{
			Header: http.Header{"Content-Type": []string{"text/css"}},
		})
	}

	vm := viewmodel.NewClientReceipt()
	w.Header().Add("Content-type", "text/html")
	c.clientReceiptTemplate.Execute(w, vm)
}

func (c clientReceipt) handleClientReceiptDetails(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		pusher.Push("/css/app.css", &http.PushOptions{
			Header: http.Header{"Content-Type": []string{"text/css"}},
		})
	}
	retrievedClientReceipt := model.GetReceipt(r.URL.Path[15:])
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/clientreceipt", http.StatusMovedPermanently)
		model.UpdateReceiptPaymentStatus(retrievedClientReceipt)
	}
	vm := viewmodel.RetrievedClientReceipt(retrievedClientReceipt)
	w.Header().Add("Content-type", "text/html")
	c.retrievedClientReceiptTemplate.Execute(w, vm)

}
