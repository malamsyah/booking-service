package service

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/malamsyah/booking-service/booking/model"
)

type BookingService interface {
	CreateBooking(booking model.Booking) (*model.Booking, error)
	GetBooking(id string) (*model.Booking, error)
}

type Service struct {
	Repo redis.Conn
}

func (s *Service) CreateBooking(booking model.Booking) (*model.Booking, error) {
	marshalledBooking, err := json.Marshal(booking)
	if err != nil {
		return nil, err
	}
	_, err = s.Repo.Do("SET", fmt.Sprintf("booking/%s", booking.ID), marshalledBooking)
	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (s *Service) GetBooking(id string) (*model.Booking, error) {

	res, err := redis.String(s.Repo.Do("GET", fmt.Sprintf("booking/%s", id)))
	if err != nil {
		return nil, err
	}

	var booking model.Booking

	err = json.Unmarshal([]byte(res), &booking)
	if err != nil {
		return nil, err
	}

	return &booking, nil
}
