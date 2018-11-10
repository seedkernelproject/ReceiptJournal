package model

import (
	"log"
	"time"
)

type Receipt struct {
	InputDate         string  `json:"inputdate"`
	ReceiptDate       string  `json:"receiptdate"`
	ClientName        string  `json:"clientname"`
	Origin            string  `json:"origin"`
	Destination       string  `json:"destination"`
	TotalPrice        float64 `json:"totalprice"`
	DriverName        string  `json:"drivername"`
	PlateNumber       string  `json:"platenumber"`
	DriverCost        float64 `json:"drivercost"`
	OtherCost         float64 `json:"othercost"`
	TotalCost         float64 `json:"totalcost"`
	Profit            float64 `json:"profit"`
	AdditionalRemarks string  `json:"additionalremarks"`
	Status            bool    `json:"status"`
	PaidDate          string  `json:"paiddate"`
}

func SaveReceipt(r Receipt) {
	_, err := db.Exec(`
	INSERT INTO public."receipt journal"(
		"inputDate", "receiptDate", "clientName", origin, destination, "totalPrice", "driverName", "plateNumber", "driverCost", "otherCost", "totalCost", profit, "additionalRemarks", "Status", "paidDate")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);`, r.InputDate, r.ReceiptDate, r.ClientName, r.Origin, r.Destination, r.TotalPrice, r.DriverName, r.PlateNumber, r.DriverCost, r.OtherCost, r.TotalCost, r.Profit, r.AdditionalRemarks, r.Status, r.PaidDate)
	if err != nil {
		log.Printf("Failed to save: %v", err)
	}
}

func GetReceipts() []Receipt {
	result := []Receipt{}
	r := Receipt{}
	rows, _ := db.Query(`SELECT "inputDate", "receiptDate", "clientName", "origin", "destination", "totalPrice", "driverName", "plateNumber", "driverCost", "otherCost", "totalCost", "profit", "additionalRemarks", "Status", "paidDate"
	FROM public."receipt journal";`)
	for rows.Next() {
		rows.Scan(&r.InputDate, &r.ReceiptDate, &r.ClientName, &r.Origin, &r.Destination, &r.TotalPrice, &r.DriverName, &r.PlateNumber, &r.DriverCost, &r.OtherCost, &r.TotalCost, &r.Profit, &r.AdditionalRemarks, &r.Status, &r.PaidDate)
		r.ReceiptDate = r.ReceiptDate[8:10] + "-" + r.ReceiptDate[5:7] + "-" + r.ReceiptDate[0:4]
		r.PaidDate = r.PaidDate[8:10] + "-" + r.PaidDate[5:7] + "-" + r.PaidDate[0:4]
		result = append(result, r)
	}
	return result
}

func GetReceipt(receiptInputDate string) Receipt {
	r := Receipt{}
	row := db.QueryRow(`SELECT "inputDate", "receiptDate", "clientName", "origin", "destination", "totalPrice", "driverName", "plateNumber", "driverCost", "otherCost", "totalCost", "profit", "additionalRemarks", "Status", "paidDate"
	FROM public."receipt journal" WHERE "inputDate" = $1`, receiptInputDate)
	row.Scan(&r.InputDate, &r.ReceiptDate, &r.ClientName, &r.Origin, &r.Destination, &r.TotalPrice, &r.DriverName, &r.PlateNumber, &r.DriverCost, &r.OtherCost, &r.TotalCost, &r.Profit, &r.AdditionalRemarks, &r.Status, &r.PaidDate)
	r.ReceiptDate = r.ReceiptDate[8:10] + "-" + r.ReceiptDate[5:7] + "-" + r.ReceiptDate[0:4]
	return r
}

func UpdateReceiptPaymentStatus(r Receipt) {
	db.Exec(`
	UPDATE public."receipt journal"
	SET "Status" = $1, "paidDate" = $2
	WHERE "inputDate"=$3`, true, time.Now(), r.InputDate)

}
