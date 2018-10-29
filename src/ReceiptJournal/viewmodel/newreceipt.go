package viewmodel

import (
	"ReceiptJournal/model"
	"time"
)

type NewReceipt struct {
	Title   string
	Active  string
	Message string
	Receipt model.Receipt
}

type CreatedReceipt struct {
	Title   string
	Active  string
	Message string
	Receipt model.Receipt
}

func NewNewReceipt() NewReceipt {
	current := time.Now()
	return NewReceipt{
		Title:  "UPS Logistics Service",
		Active: "newreceipt",
		Receipt: model.Receipt{
			InputDate:   current.String(),
			ReceiptDate: current.Format("2006-01-02"),
		},
	}
}

func NewCreatedReceipt(r model.Receipt) (Receipt CreatedReceipt) {
	return CreatedReceipt{
		Title:   "UPS Logistics Service",
		Active:  "newreceipt",
		Receipt: r,
	}
}
