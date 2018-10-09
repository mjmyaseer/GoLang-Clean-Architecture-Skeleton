package http

import (
	"context"
	"fmt"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-sample/utils/go-util/config"
	"go-sample/utils/go-util/log"
	"go-sample/utils/go-util/response"
	"go-sample/infrastructure"
	"go-sample/interfaces/endpoints"
	"go-sample/services"
	"net/http"
	"time"
)

var server *http.Server

var bookingDetails services.BookingDetails

func init() {
	bookingDetails = services.NewBookingDetails()
}

var serverOptions = []httpTransport.ServerOption{
	httpTransport.ServerErrorLogger(new(infrastructure.ErrorLogger)),
	httpTransport.ServerErrorEncoder(response.HandleError),
}

func InitHttpRouter() {
	r := mux.NewRouter()

	r.Handle(`/v2/sequence/deliveryList`, httpTransport.NewServer(
		endpoints.GetBookingDetailsListEndpoint(bookingDetails),
		DecodeBookingDetailsListRequest,
		EncodeBookingDetailsListResponse,
		serverOptions...,
	)).Methods(http.MethodGet)

	r.Handle(`/metrics`, promhttp.Handler())

	startServer(r)
}

func StopHttpRouter() {
	if err := server.Shutdown(context.Background()); err != nil {
		log.Error(err)
	}
}

func startServer(r http.Handler) {
	port := config.AppConf.Port
	running := make(chan interface{}, 1)

	server = &http.Server{
		Addr: fmt.Sprintf(`:%d`, port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.RecoveryHandler()(r), // Pass our instance of gorilla/mux in.
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(log.WithPrefix(`Cannot start web server : `, err))
		}
		running <- `done`
	}()
	log.Info(log.WithPrefix(`transport.http.router`, fmt.Sprintf(`http router started on port [%d]`, port)))
	<-running
}
