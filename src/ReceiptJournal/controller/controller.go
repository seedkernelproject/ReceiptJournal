package controller

import (
	"html/template"
	"net/http"
)

var (
	newReceiptController     newreceipt
	createdReceiptController createdreceipt
	cashflowController       cashflow
	clientReceiptController  clientReceipt
)

func Startup(templates map[string]*template.Template) {
	newReceiptController.newreceiptTemplate = templates["newreceipt.html"]
	newReceiptController.registerRoutes()
	createdReceiptController.createdreceiptTemplate = templates["createdreceipt.html"]
	createdReceiptController.registerRoutes()
	cashflowController.cashflowTemplate = templates["cashflow.html"]
	cashflowController.retrievedReceiptTemplate = templates["retrievedreceipt.html"]
	cashflowController.registerRoutes()
	clientReceiptController.clientReceiptTemplate = templates["clientreceipt.html"]
	clientReceiptController.retrievedClientReceiptTemplate = templates["retrievedclientreceipt.html"]
	clientReceiptController.registerRoutes()
	registerAPIRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
