package api

import "github.com/malamsyah/booking-service/booking/model"

type CreateRequest struct {
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
}

type CreateResponse struct {
	Status  string        `json:"name"`
	Booking model.Booking `json:"booking"`
}
