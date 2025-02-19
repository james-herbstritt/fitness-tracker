package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	addQueryParamsToRequest(req, params)
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

func (c *FitbitClient) GetProfile(ctx context.Context) (*Profile, error) {
	path := fmt.Sprintf(GET_PROFILE_PATH, "-")
	u := c.BaseURL.JoinPath(path)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch profile: " + resp.Status)
	}

	var result *Profile
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
