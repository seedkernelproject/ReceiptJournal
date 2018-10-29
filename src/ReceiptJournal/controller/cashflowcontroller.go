package controller

import (
	"ReceiptJournal/model"
	"ReceiptJournal/viewmodel"
	"html/template"
	"net/http"
)

type cashflow struct {
	cashflowTemplate         *template.Template
	retrievedReceiptTemplate *template.Template
}

func (c cashflow) registerRoutes() {
	//	http.HandleFunc("/", c.handleCashflow)
	http.HandleFunc("/cashflow", c.handleCashflow)
	http.HandleFunc("/cashflow/", c.handleReceiptDetails)
}

func (c cashflow) handleCashflow(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		pusher.Push("/css/app.css", &http.PushOptions{
			Header: http.Header{"Content-Type": []string{"text/css"}},
		})
	}

	vm := viewmodel.NewCashflow()
	w.Header().Add("Content-type", "text/html")
	c.cashflowTemplate.Execute(w, vm)
}

func (c cashflow) handleReceiptDetails(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		pusher.Push("/css/app.css", &http.PushOptions{
			Header: http.Header{"Content-Type": []string{"text/css"}},
		})
	}
	retrievedReceipt := model.GetReceipt(r.URL.Path[10:])
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/cashflow", http.StatusMovedPermanently)
		model.UpdateReceiptPaymentStatus(retrievedReceipt)
	}
	vm := viewmodel.RetrievedReceipt(retrievedReceipt)
	w.Header().Add("Content-type", "text/html")
	c.retrievedReceiptTemplate.Execute(w, vm)

}
