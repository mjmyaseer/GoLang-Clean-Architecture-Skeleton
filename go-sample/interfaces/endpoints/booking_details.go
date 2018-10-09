package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-sample/domain"
	"go-sample/encoders"
	"go-sample/services"
)

func GetBookingCountEndpoint(service services.BookingDetails) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		return service.GetTotalCounts()
	}
}

func GetBookingDetailsListEndpoint(service services.BookingDetails) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(domain.AllParam)

		res, pag, err := service.GetBookingsUrl(req)

		resFor := encoders.Response{}
		resFor.Data = res
		resFor.Pagination = pag

		return resFor, nil
	}
}
