package service

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/malamsyah/booking-service/booking/model"
)

type BookingService interface {
	CreateBooking(booking model.Booking) error
}

type Service struct {
	Repo redis.Conn
}

func (s *Service) CreateBooking(booking model.Booking) error {
	marshalledBooking, err := json.Marshal(booking)
	if err != nil {
		return err
	}
	_, err = s.Repo.Do("SET", fmt.Sprintf("booking/%s", booking.ID), marshalledBooking)
	if err != nil {
		return err
	}

	return nil
}
