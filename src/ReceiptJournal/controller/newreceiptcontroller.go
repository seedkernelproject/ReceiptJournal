package controller

import (
	"ReceiptJournal/model"
	"ReceiptJournal/viewmodel"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	CreatedReceipt model.Receipt
)

type newreceipt struct {
	newreceiptTemplate *template.Template
}

func (n newreceipt) registerRoutes() {
	http.HandleFunc("/newreceipt", n.handleNewReceipt)
}

func (n newreceipt) handleNewReceipt(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		pusher.Push("/css/app.css", &http.PushOptions{
			Header: http.Header{"Content-Type": []string{"text/css"}},
		})
	}
	vm := viewmodel.NewNewReceipt()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("Error loggin in: %v", err))
		}
		CreatedReceipt.InputDate = vm.Receipt.InputDate
		CreatedReceipt.ReceiptDate = r.Form.Get("receiptDate")
		CreatedReceipt.ClientName = r.Form.Get("clientName")
		CreatedReceipt.Origin = r.Form.Get("origin")
		CreatedReceipt.Destination = r.Form.Get("destination")
		CreatedReceipt.TotalPrice, _ = strconv.ParseFloat(r.Form.Get("totalPrice"), 64)
		CreatedReceipt.DriverName = r.Form.Get("driverName")
		CreatedReceipt.PlateNumber = r.Form.Get("plateNumber")
		CreatedReceipt.DriverCost, _ = strconv.ParseFloat(r.Form.Get("driverCost"), 64)
		CreatedReceipt.OtherCost, _ = strconv.ParseFloat(r.Form.Get("otherCost"), 64)
		CreatedReceipt.TotalCost = CreatedReceipt.DriverCost + CreatedReceipt.OtherCost
		CreatedReceipt.Profit = CreatedReceipt.TotalPrice - CreatedReceipt.TotalCost
		CreatedReceipt.AdditionalRemarks = r.Form.Get("additionalRemarks")
		CreatedReceipt.Status = false
		CreatedReceipt.PaidDate = "9999-12-31"
		if CreatedReceipt.ClientName == "" {
			log.Println(fmt.Errorf("Field Incomplete"))
			vm.Message = "Field Incomplete"
		} else {
			http.Redirect(w, r, "/createdreceipt", http.StatusMovedPermanently)
			log.Println(CreatedReceipt)
		}
	}
	w.Header().Add("Content-type", "text/html")
	n.newreceiptTemplate.Execute(w, vm)
}
