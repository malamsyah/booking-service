package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/malamsyah/booking-service/booking/model"
	"github.com/malamsyah/booking-service/booking/service"
	"net/http"
)

func GetBookingHandler(service service.BookingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		param := mux.Vars(r)
		id := param["id"]
		if id == "" {
			http.Error(w, "empty id", http.StatusBadRequest)
			return
		}

		booking, err := service.GetBooking(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response, _ := json.Marshal(CreateResponse{
			Status:  "success",
			Booking: *booking,
		})

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
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

		booking, err := service.CreateBooking(model.Booking{
			ID:     uuidWithHyphen.String(),
			Name:   request.Name,
			Amount: request.Amount,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, _ := json.Marshal(CreateResponse{
			Status:  "success",
			Booking: *booking,
		})

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}
}
