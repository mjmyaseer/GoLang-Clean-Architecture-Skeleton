package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-sample/utils/go-util/log"
	"go-sample/utils/go-util/mysql"
	"go-sample/domain"
	"go-sample/interfaces/repository/mysql"
	"strconv"
	"strings"
	"time"
)

const EQUAL = "="

const LIKE = "like"

const BEGIN_WITH = "beginWith"

const END_WITH = "endWith"

const BETWEEN = "between"

const IN = "in"

const CUSTOM = "cust"

const NOT_EQUAL = "notEqual"

const GREATER_THAN = ">"

const LESS_THAN = "<"

func SelectTripDetails(s domain.BookingDetails) []domain.Trip {
	arg := []interface{}{}
	arg = append(arg, s.TripID)

	q, err := database.Connections.Read.QueryContext(context.Background(), `SELECT * FROM event_sequence where trip_id = ?`, arg...)
	if err != nil {
		fmt.Println(err)
	}

	trp := []domain.Trip{}
	var tr domain.Trip
	for q.Next() {
		a := q.Scan(&tr.ID, &tr.BookingTime, &tr.StartTime, &tr.EndTime,
			&tr.Pickup.Address, &tr.Drop.Address, &tr.Passenger.PassengerId, &tr.Status,
			&tr.Driver.DriverId, &tr.Payment.PaymentMethod)
		if a != nil {
			println(a)
			continue
		}
		trp = append(trp, tr)
	}
	return trp
}

func SelectTripDetailsUrl(s domain.AllParam) (trips []domain.Trip, paging domain.Paging, err error) {
	var where string
	var oper string

	where = "where "
	joins := " "
	having := ""
	f := time.Now().Format("2006-01-02")
	toDate := "where jcre.created_date = '" + f + "' "
	x := 0

	for i, res := range s.Param {
		x = x + 1

		if i >= 0 {
			toDate = " and jcre.created_date = '" + f + "'"
		}

		switch res.Operator {
		case 0:
			oper = CUSTOM
		case 1:
			oper = EQUAL
		case 2:
			oper = LIKE
		case 3:
			oper = GREATER_THAN
		case 4:
			oper = LESS_THAN
		case 5:
			oper = BETWEEN
		case 6:
			oper = IN
		case 7:
			oper = NOT_EQUAL
		case 8:
			oper = BEGIN_WITH
		case 9:
			oper = END_WITH
		}

		if i > 0 {
			where = where + " and "
		}

		fieldVal := ""
		search := " "

		if res.Field == "job_id" || res.Field == "event_type" {
			fieldVal = "jcre." + res.Field
			search = fieldVal + oper + "'" + res.Value + "'"
		}
		if res.Field == "event_type" {
			joins = "join " + mysql.TableEventSequence + " as es on jcre.job_id = es.job_id and es.event_type = '" + res.Value + "' "
			fieldVal = "es." + res.Field
			having = "having status = '" + res.Value + "' "
			search = fieldVal + oper + "'" + res.Value + "'"
		}
		if res.Field == "passenger_id" || res.Field == "module" {
			fieldVal = "jcre." + res.Field
			search = fieldVal + oper + "'" + res.Value + "'"
		}
		if res.Field == "driver_id" {
			fieldVal = "ja." + res.Field
			search = fieldVal + oper + "'" + res.Value + "'"
		}
		if res.Field == "completed_date" {
			fieldVal = "jcre.created_date"
			search = fieldVal + oper + "'" + res.Value + "'"
			toDate = ""

		}

		if res.Field == "payment_method" {
			joins = "join " + mysql.TableJobPaymentType + " as jpay on jcre.job_id = jpay.job_id "
			fieldVal = "jpay.payment_method"
			search = fieldVal + oper + "'" + res.Value + "'"
		}

		if res.Field == "region_id" {
			joins = "join " + mysql.TableJobRegion + " as jreg on jcre.job_id = jreg.job_id and jreg.region_id=" + res.Value + " "
			fieldVal = "jreg.region_id"
			search = fieldVal + oper + "'" + res.Value + "'"
		}
		where = where + " " + search
	}

	if x == 0 {
		where = ""
	}

	selectQ := "select "
	selectFields := `jcre.job_id,
		jcre.created_at as booking_time,
		ja.created_at as accepted_time,
		jcom.created_at as completed_time,
		jcre.pickup_name,
		jcre.pickup_phone,
		jcre.pickup_location,
		jcre.drop_location,
		jcre.passenger_id,`

	selectSub := "(SELECT event_type FROM `event_sequence` WHERE `job_id` = jcre.job_id ORDER BY `id` DESC LIMIT 1) as status,"
	selectSecField := `ja.driver_id,
		jcre.payment_method,
		jcre.pickup_lat,
		jcre.pickup_lon,
		jcre.drop_lat,
		jcre.drop_lon,
		jord.order_details`

	fromQ := ` FROM ` + mysql.TableJobCreated + ` as jcre `
	joins = joins + `Left Join ` + mysql.TableJobAccepted + ` as ja on jcre.job_id = ja.job_id `
	joins = joins + ` Left Join ` + mysql.TableJobCompleted + ` as jcom on jcre.job_id = jcom.job_id `
	joins = joins + ` Left Join ` + mysql.TableOrderDetails + ` as jord on jcre.job_id = jord.job_id `
	order := " order by jcre.job_id desc"
	pageNum := s.Pages.PageNo - 1
	newPageLim := pageNum * s.Pages.PerPage
	limit := ` limit ` + strconv.Itoa(s.Pages.PerPage)
	offset := ` offset ` + strconv.Itoa(newPageLim)
	countStr := "SELECT COUNT(job_id) as Counts FROM ("
	countEnd := ") AS A"

	q, err := database.Connections.Read.QueryContext(context.Background(), selectQ+selectFields+selectSub+
		selectSecField+fromQ+joins+" "+where+" "+toDate+" "+having+" "+order+limit+offset)
	if err != nil && err != sql.ErrNoRows {
		log.Error(err)
	}

	trp := []domain.Trip{}
	var tr domain.Trip
	for q.Next() {
		a := q.Scan(&tr.ID, &tr.BookingTime, &tr.StartTime, &tr.EndTime, &tr.Pickup.Name, &tr.Pickup.Phone,
			&tr.Pickup.Address, &tr.Drop.Address, &tr.Passenger.PassengerId, &tr.Status,
			&tr.Driver.DriverId, &tr.Payment.PaymentMethod, &tr.Pickup.Lat, &tr.Pickup.Lon, &tr.Drop.Lat, &tr.Drop.Lon, &tr.Order.OrderDetails)

		if a != nil {
			log.Error("error from query no 1 ", a)
			continue
		}
		trp = append(trp, tr)
	}

	row := database.Connections.Read.QueryRow(countStr + selectQ + selectFields + selectSub +
		selectSecField + fromQ + joins + " " + where + " " + toDate + " " + having + " " + order + countEnd)

	var counts int
	err = row.Scan(&counts)
	if err != nil && err != sql.ErrNoRows {
		log.Error(err)
	}

	var rar string
	for _, value := range trp {
		rar = rar + strconv.FormatInt(value.ID.Int64, 10) + ","
	}

	rar = strings.TrimRight(rar, ",")

	if rar != "" {
		SelectRej := "select "
		SelectRejField := `job_id, driver_id, rejection_type,address,rejected_lat,rejected_lon `
		SelectRejFrom := "From " + mysql.TableJobRejected
		SelectRejWhere := ` where job_id IN `
		OrderBy := "order by id desc"

		ids, err := database.Connections.Read.QueryContext(context.Background(), SelectRej+SelectRejField+
			SelectRejFrom+SelectRejWhere+"("+rar+") "+OrderBy)

		if err != nil {
			log.Error("Query only ", err)
		}

		rejs := []domain.Rejected{}
		var rej domain.Rejected
		for ids.Next() {

			b := ids.Scan(&rej.JobID, &rej.DriverID, &rej.RejectType, &rej.Location.Address,
				&rej.Location.Latitude, &rej.Location.Longitude)
			if b != nil {
				log.Error("error from query no 2 ", b)
				continue
			}
			rejs = append(rejs, rej)
		}

		for i, _ := range trp {
			for _, rng := range rejs {
				if trp[i].ID.Int64 == rng.JobID.Int64 {
					trp[i].Rejecteds = append(trp[i].Rejecteds, rng)
				}
			}
		}

	}

	paging.PageNo = s.Pages.PageNo
	paging.PerPage = s.Pages.PerPage
	paging.TotalRecords = counts
	return trp, paging, nil
}

func GetTotalTripCounts() (domain.TotalCount, error) {
	var cnt domain.TotalCount

	//comp, err := database.Connections.Read.QueryContext(context.Background(), "Select count(id) as completed From job_completed")
	//canc, err := database.Connections.Read.QueryContext(context.Background(), "Select count(id) as cancelled From " +
	//	""+ mysql.TableEventSequence +" where event_type = 'job_cancelled'")
	//app, err := database.Connections.Read.QueryContext(context.Background(), "Select count(id) as app From " +
	//	""+ mysql.TableJobCreated +" where module = '2'")

	return cnt, nil
}
