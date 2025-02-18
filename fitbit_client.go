package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type FitbitClient struct {
	BaseURL   *url.URL
	AuthToken string

	HttpClient *http.Client
}

func addQueryParamsToRequest(req *http.Request, params map[string]string) {
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
}

func NewFitbitClient(authToken string) *FitbitClient {
	return &FitbitClient{
		BaseURL: &url.URL{
			Scheme: "https",
			Host:   "api.fitbit.com",
		},
		AuthToken:  authToken,
		HttpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *FitbitClient) doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", "Bearer "+c.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	return c.HttpClient.Do(req)
}

func (c *FitbitClient) GetActivities(ctx context.Context, params map[string]string) (*ActivityLogList, error) {
	path := fmt.Sprintf(GET_ACTIVITY_LOG_LIST_PATH, "-")
	u := c.BaseURL.JoinPath(path)

	fmt.Println(u)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	addQueryParamsToRequest(req, params)
	fmt.Println(req.URL.String())
	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch activities: " + resp.Status)
	}

	var result *ActivityLogList
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// the return http request might be a Fitbit request?
// func (c *FitbitClient) newRequest(method string, path string, body interface{}) (*http.Request, error) {
//     rel := &url.URL{Path: path}
//     u := c.BaseURL.ResolveReference(rel)

//     if (body == nil) {
//        return nil, errors.New("body argument to newRequest was nil")
//     }
//     encodedBody, err := encodeBody(body)
//     if (err != nil) {
//         return nil, err
//     }
//     req, err := http.NewRequest(method, u.String(), encodedBody)
//     if err != nil {
//         return nil, err
//     }

//     // TODO:: Make this cleaner somehow
//     bearerToken := fmt.Sprintf("Bearer %s", c.AccessToken)
// 	header := http.Header{
// 		"accept":          {"application/json"},
// 		"accept-langauge": {"en_US"},
// 		"accept-locale":   {"en_US"},
// 		"Authorization":   {bearerToken},
// 	}

//     req.Header = header
//     return req, err
// }

func encodeBody(body interface{}) (io.ReadWriter, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	return buf, err
}
