package dto

type TopUpBalanceRequest struct {
	Amount float64 `json:"amount"`
}

type TopUpBalanceResponse struct {
	Balance float64 `json:"balance"`
}
