package domain

import (
	"database/sql"
	"encoding/json"
	"github.com/pkg/errors"
	"reflect"
)

type BookingDetails struct {
	TripID         int    `json:"trip_id"`
	Type           int    `json:"type"`
	Status         int    `json:"status"`
	Region         int    `json:"region"`
	PaymentType    int    `json:"payment_type"`
	DriverID       int    `json:"driver_id"`
	PassengerPhone int    `json:"passenger_phone"`
	PassengerName  string `json:"passenger_name"`
	VehicleType    int    `json:"vehicle_type"`
	CreatedDate    int    `json:"created_date"`
}

type Trip struct {
	ID          NullInt64  `json:"id"`
	Status      NullString `json:"status"`
	BookingTime NullInt64  `json:"booking_time"`
	StartTime   NullInt64  `json:"accepted_time"`
	EndTime     NullInt64  `json:"completed_time"`
	Pickup      struct {
		Name    NullString `json:"name"`
		Phone   NullString `json:"phone"`
		Address NullString `json:"address"`
		Lon     NullString `json:"lon"`
		Lat     NullString `json:"lat"`
	} `json:"pickup"`
	Drop struct {
		Address NullString `json:"address"`
		Lon     NullString `json:"lon"`
		Lat     NullString `json:"lat"`
	} `json:"drop"`
	Passenger struct {
		PassengerId NullInt64 `json:"passenger_id"`
	} `json:"passenger"`
	Driver struct {
		DriverId NullInt64 `json:"driver_id"`
	} `json:"driver"`
	Payment struct {
		PaymentMethod NullString `json:"payment_method"`
	} `json:"payment"`
	Order struct {
		OrderDetails NullString `json:"order_details"`
	} `json:"order"`
	Rejecteds []Rejected
}

type Rejected struct {
	JobID    NullInt64 `json:"job_id"`
	DriverID NullInt64 `json:"driver_id"`
	//DriverPhone NullInt64 `json:"driver_phone"`
	//DriverName NullString `json:"driver_name"`
	RejectType NullString `json:"rejection_type"`
	Location   struct {
		Address   NullString `json:"address"`
		Latitude  NullInt64  `json:"rejected_lat"`
		Longitude NullInt64  `json:"rejected_lon"`
	}
}

type UrlParamStr struct {
	Field    string `json:"field"`
	Operator int    `json:"operator"`
	Value    string `json:"value"`
}

type Paging struct {
	PageNo       int `json:"page"`
	PerPage      int `json:"size"`
	TotalRecords int `json:"total_records"`
}

type AllParam struct {
	Param []UrlParamStr
	Pages Paging
}

type TotalCount struct {
	Completed  NullInt64 `json:"completed"`
	Cancelled  NullInt64 `json:"cancelled"`
	App        NullInt64 `json:"app"`
	Dispatcher NullInt64 `json:"dispatcher"`
	RoadPickup NullInt64 `json:"road_pickup"`
}

// CUSTOM NULL Handling structures

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 sql.NullInt64

// Scan implements the Scanner interface for NullInt64
func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{i.Int64, false}
	} else {
		*ni = NullInt64{i.Int64, true}
	}
	return nil
}

// NullString is an alias for sql.NullString data type
type NullString sql.NullString

// Scan implements the Scanner interface for NullString
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return err
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}

func (BookingDetails) Encode(data interface{}) ([]byte, error) {
	book, ok := data.(BookingDetails)
	if !ok {
		return nil, errors.New(`invalid type, expected Rejection`)
	}
	j, _ := json.Marshal(book)
	return j, nil
}

func (BookingDetails) Decode(data []byte) (interface{}, error) {
	o := BookingDetails{}
	json.Unmarshal(data, &o)
	return o, nil
}
