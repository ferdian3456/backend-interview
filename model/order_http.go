package model

import "time"

type OrderResponse struct {
	ID              string    `json:"id"`
	ProductID       string    `json:"productID"`
	ProductName     string    `json:"productName"`
	Amount          string    `json:"amount"`
	CustomerName    string    `json:"customerName"`
	Status          int       `json:"status"`
	TransactionDate time.Time `json:"transactionDate"`
	CreateBy        string    `json:"createBy"`
	CreateOn        time.Time `json:"createOn"`
}
