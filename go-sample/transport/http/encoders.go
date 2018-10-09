package http

import (
	"context"
	"encoding/json"
	"net/http"
)

type Response struct {
	Pagination interface{} `json:"pagination"`
	Data       interface{} `json:"data"`
}

func EncodeBookingCountResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func EncodeBookingDetailsListResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
