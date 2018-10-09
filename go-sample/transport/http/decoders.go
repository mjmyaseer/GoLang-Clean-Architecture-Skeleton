package http

import (
	"context"
	"encoding/json"
	"go-sample/utils/go-util/log"
	"go-sample/domain"
	"net/http"
	"net/url"
)

func DecodeBookingCountRequest(_ context.Context, r *http.Request) (interface{}, error) {

	return r, nil
}

func DecodeBookingDetailsListRequest(_ context.Context, r *http.Request) (interface{}, error) {

	str, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Error(err)
	}

	v := string(str["filters"][0])
	p := string(str["paging"][0])

	fil := []byte(v)
	val2 := []domain.UrlParamStr{}
	err = json.Unmarshal(fil, &val2)
	if err != nil {
		log.Error(err)
	}

	page := []byte(p)
	pa := domain.Paging{}
	err = json.Unmarshal(page, &pa)

	if err != nil {
		log.Error(err)
	}

	allPar := domain.AllParam{}

	allPar.Param = val2
	allPar.Pages = pa

	return allPar, nil
}
