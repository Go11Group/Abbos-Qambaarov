package models

type Transactions struct {
	Id   string `json:"id"`
	Cardid string `json:"card_id"`
	Amount int `json:"amount"`
	TerminalId string `json:"terminal_id"`
	TransactionType string `json:"transaction_type"`
}

type CreateTransaction struct {
	Cardid string `json:"card_id"`
	Amount int `json:"amount"`
	TerminalId string `json:"terminal_id"`
	TransactionType string `json:"transaction_type"`
}
