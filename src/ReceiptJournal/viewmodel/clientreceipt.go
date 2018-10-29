package viewmodel

import (
	"ReceiptJournal/model"
)

type ClientReceipt struct {
	Title      string
	Active     string
	Message    string
	Receipts   []model.Receipt
	TotalPrice float64
}

func NewClientReceipt() ClientReceipt {
	receipts, totalPrice := TotalPrice()
	return ClientReceipt{
		Title:      "UPS Logistics Service",
		Active:     "clientreceipt",
		Receipts:   receipts,
		TotalPrice: totalPrice,
	}
}

func RetrievedClientReceipt(r model.Receipt) (Receipt CreatedReceipt) {
	return CreatedReceipt{
		Title:   "UPS Logistics Service",
		Active:  "clientreceipt",
		Receipt: r,
	}
}

func TotalPrice() ([]model.Receipt, float64) {
	receipts := model.GetReceipts()
	totalPrice := 0.
	for i := range receipts {
		if receipts[i].Status != true {
			totalPrice += receipts[i].TotalPrice
		}
	}
	return receipts, totalPrice
}
