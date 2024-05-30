package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// No context
func Get(url string, header *http.Header) *http.Response {
	fr := FitbitRequest{
		Type: "GET",
		URL: url,
		Header: header,
	}
	return fr.doRequest()
}

func GetWithContext(c context.Context, url string, header *http.Header) *http.Response {
	fr := FitbitRequestWithContext{
		Context: c,
		Type: "GET",
		URL: url,
		Header: header,
	}
	return fr.doRequest()
}


// I Don't think I need these anymore if ^ works
func GetWithQueryParams(url string, header *http.Header, params map[string]string) *http.Response {
	fr := FitbitRequest{
		Type: "GET",
		URL: url,
		Header: header,
		QueryParams: params,
	}
	return fr.doRequest()
}

func GetWithContextWithQueryParams(c context.Context, url string, header *http.Header, params map[string]string) *http.Response {
	fr := FitbitRequestWithContext{
		Context: c,
		Type: "GET",
		URL: url,
		Header: header,
		QueryParams: params,
	}
	return fr.doRequest()
}

func BuildRequest(url string, header *http.Header) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header = *header
	return req
}

func BuildRequestWithContext(c context.Context, url string, header *http.Header) *http.Request {
	req, err := http.NewRequestWithContext(c, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header = *header
	return req
}

func ProcessResponseBody[GRS GetReturnStruct](res *http.Response, grs *GRS) {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, grs)
	if err != nil {
		panic(err)
	}
	var lifetimeStats LifetimeStats
	err = json.Unmarshal(b, &lifetimeStats)
	if err != nil {
		panic(err)
	}
}

// We probably need a utils file or something at this rate
// func Func( ...any )
func MakeUrlAndHeader(urlTemplate string, accessToken string, params ...any) (string, *http.Header) {
	url := fmt.Sprintf(urlTemplate, params...)
	bearerToken := fmt.Sprintf("Bearer %s", accessToken)

	header := http.Header{
		"accept":          {"application/json"},
		"accept-langauge": {"en_US"},
		"accept-locale":   {"en_US"},
		"Authorization":   {bearerToken},
	}

	return url, &header
}

func AddQueryParamsToRequest(req *http.Request, params map[string]string) {
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
}


// TODO Write a Fitbit object that has like Fitbit.auth
// Fitbit.GetSpecificFitbitThing etc
// Does this need to have an AccessToken as a field? I think it should???
type FitbitRequest struct {
    Type, URL string
    QueryParams map[string]string
    Header *http.Header
}

type FitbitRequestWithContext struct {
	Context context.Context
	Type, URL string
    QueryParams map[string]string
    Header *http.Header
}

//something like this?
func (FR FitbitRequest) doRequest() *http.Response { 
	req := BuildRequest(FR.URL, FR.Header)
	if len(FR.QueryParams) != 0 { 
		AddQueryParamsToRequest(req, FR.QueryParams)
	}

	// this should probably be some sort of fitbit client?
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return res
}

func (FR FitbitRequestWithContext) doRequest() *http.Response {
	req := BuildRequestWithContext(FR.Context, FR.Type, FR.Header)
	if len(FR.QueryParams) != 0 { 
		AddQueryParamsToRequest(req, FR.QueryParams)
	}

	// this should probably be some sort of fitbit client?
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return res
}