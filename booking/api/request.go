package api

type CreateRequest struct {
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
}
