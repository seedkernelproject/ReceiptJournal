package viewmodel

import (
	"ReceiptJournal/model"
)

type Cashflow struct {
	Title             string          `json:"title,omitempty"`
	Active            string          `json:"active,omitempty"`
	Message           string          `json:"message,omitempty"`
	Receipts          []model.Receipt `json:"receipt,omitempty"`
	TotalUnpaidProfit float64         `json:"totalunpaidprofit,omitempty"`
	TotalPaidProfit   float64         `json:"totalpaidprofit,omitempty"`
	TotalProfit       float64         `json:"totalprofit,omitempty"`
}

func NewCashflow() Cashflow {
	receipts, totalUnpaidProfit, totalPaidProfit, totalProfit := TotalProfit()
	return Cashflow{
		Title:             "UPS Logistics Service",
		Active:            "cashflow",
		Receipts:          receipts,
		TotalUnpaidProfit: totalUnpaidProfit,
		TotalPaidProfit:   totalPaidProfit,
		TotalProfit:       totalProfit,
	}
}

func TotalProfit() ([]model.Receipt, float64, float64, float64) {
	receipts := model.GetReceipts()
	totalUnpaidProfit, totalPaidProfit, totalProfit := 0., 0., 0.
	for i := range receipts {
		if receipts[i].Status == true {
			totalPaidProfit += receipts[i].Profit
		} else {
			totalUnpaidProfit += receipts[i].Profit
		}
		totalProfit += receipts[i].Profit
	}
	return receipts, totalUnpaidProfit, totalPaidProfit, totalProfit
}

func RetrievedReceipt(r model.Receipt) (Receipt CreatedReceipt) {
	return CreatedReceipt{
		Title:   "UPS Logistics Service",
		Active:  "cashflow",
		Receipt: r,
	}
}
