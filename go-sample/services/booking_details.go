package services

import (
	"go-sample/utils/go-util/log"
	"go-sample/domain"
	"go-sample/interfaces/repository"
)

type BookingDetails interface {
	GetBookings(domain.BookingDetails) ([]domain.Trip, error)
	GetBookingsUrl(param domain.AllParam) ([]domain.Trip, domain.Paging, error)
	GetTotalCounts() (domain.TotalCount, error)
}

type bookingDetails struct{}

func NewBookingDetails() BookingDetails {
	return new(bookingDetails)
}

func (s bookingDetails) GetBookings(b domain.BookingDetails) ([]domain.Trip, error) {
	res := repository.SelectTripDetails(b)
	return res, nil
}

func (s bookingDetails) GetBookingsUrl(b domain.AllParam) ([]domain.Trip, domain.Paging, error) {
	res, pag, err := repository.SelectTripDetailsUrl(b)

	if err != nil {
		log.Warn(err)
	}
	return res, pag, err
}

func (s bookingDetails) GetTotalCounts() (domain.TotalCount, error) {
	res, err := repository.GetTotalTripCounts()
	return res, err
}
