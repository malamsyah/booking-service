package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/malamsyah/booking-service/booking/model"
	"github.com/malamsyah/booking-service/booking/service"
	"net/http"
)

func GetBookingHandler(service service.BookingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func CreateBookingHandler(service service.BookingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request CreateRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		uuidWithHyphen := uuid.New()

		err = service.CreateBooking(model.Booking{
			ID:     uuidWithHyphen.String(),
			Name:   request.Name,
			Amount: request.Amount,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}
}
