package main

import (
	"github.com/gomodule/redigo/redis"
	"github.com/malamsyah/booking-service/booking/api"
	"github.com/malamsyah/booking-service/booking/service"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func setupRedis(address string) (redis.Conn, error) {
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		return nil, err

	}
	return conn, nil
}

func setupRouter() {
	r := mux.NewRouter()

	redisConn, err := setupRedis("127.0.0.1:6379")
	if err != nil {
		panic(err)
	}

	bookingService := &service.Service{
		redisConn,
	}

	r.HandleFunc("/booking", api.CreateBookingHandler(bookingService)).Methods(http.MethodPost)
	r.HandleFunc("/booking/{id}", api.GetBookingHandler(bookingService)).Methods(http.MethodGet)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
