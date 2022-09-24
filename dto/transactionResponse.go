package dto

type TransactionResponse struct {
	Id              int     `json:"transaction_id" xml:"transaction_id"`
	AccountId       int     `json:"account_id" xml:"account_id"`
	Amount          float64 `json:"new_balance" xml:"new_balance"`
	TransactionType string  `json:"transaction_type" xml:"transaction_type"`
	TransactionDate string  `json:"transaction_date" xml:"transaction_date"`
}
