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

// This is mutation, not sure I love it
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

// Not sure how worth it this func is
func (c *FitbitClient) buildGET(ctx context.Context, path string, urlargs ...any) (*http.Request, error) {
	p := fmt.Sprintf(path, urlargs...)
	u := c.BaseURL.JoinPath(p)

	fmt.Println(u.String())

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *FitbitClient) GetActivities(ctx context.Context, userId string, params map[string]string) (*ActivityLogList, error) {
	req, err := c.buildGET(ctx, GET_ACTIVITY_LOG_LIST_PATH, userId)
	if err != nil {
		panic(err)
	}

	addQueryParamsToRequest(req, params)
	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed get request at path " + GET_ACTIVITY_LOG_LIST_PATH + ": " + resp.Status)
	}

	var result *ActivityLogList
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *FitbitClient) GetProfile(ctx context.Context, userId string) (*Profile, error) {
	req, err := c.buildGET(ctx, GET_PROFILE_PATH, userId)
	if err != nil {
		panic(err)
	}

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed get request at path " + GET_PROFILE_PATH + ": " + resp.Status)
	}

	var result *Profile
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *FitbitClient) GetBadges(ctx context.Context, userId string) (*BadgeList, error) {
	req, err := c.buildGET(ctx, GET_BADGES_PATH, userId)
	if err != nil {
		panic(err)
	}

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed get request at path " + GET_BADGES_PATH + ": " + resp.Status)
	}

	var result *BadgeList
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *FitbitClient) GetLifetimeStats(ctx context.Context, userId string) (*LifetimeStats, error) {
	req, err := c.buildGET(ctx, GET_LIFETIME_STATS_PATH, userId)
	if err != nil {
		panic(err)
	}

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("fialed get request at path " + GET_LIFETIME_STATS_PATH + ": " + resp.Status)
	}

	var result *LifetimeStats
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
