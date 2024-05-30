package main

import (
	"context"
	"time"
)

type TotalTime struct {
	time.Time
}

func (t *TotalTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

type LifetimeStats struct {
	Best struct {
		Total struct {
			Distance struct {
				Date TotalTime `json:"date,omitempty"`
				Value float64 `json:"value,omitempty"`
			} `json:"distance"`
			Floors struct {
				Date TotalTime `json:"date,omitempty"`
				Value float64 `json:"value,omitempty"`
			} `json:"floors"`
			Steps struct {
				Date TotalTime `json:"date,omitempty"`
				Value int64 `json:"value,omitempty"`
			} `json:"steps"`
		} `json:"total"`
		Tracker struct {
			Distance struct {
				Date TotalTime `json:"date"`
				Value float64 `json:"value"`
			} `json:"distance"`
			Floors struct {
				Date TotalTime `json:"date"`
				Value float64 `json:"value"`
			} `json:"floors"`
			Steps struct {
				Date TotalTime `json:"date"`
				Value int64 `json:"value"`
			} `json:"steps"`
		} `json:"tracker"`
	} `json:"best"`
	Lifetime struct {
		Total struct {
			// ActiveScore and CaloriesOut are here for backwards compatibility, will always be -1
			ActiveScore int64 `json:"activeScore"`
			CaloriesOut int64 `json:"caloriesOut"`
			Distance float64 `json:"distance"`
			Floors float64 `json:"floors"`
			Steps int64 `json:"steps"`
		}
	} `json:"lifetime"`
}

func GetLifetimeStats(userId string, accessToken string) *LifetimeStats {
	url, header := MakeUrlAndHeader(GetLifetimeStatsUrl, accessToken, userId)
	res := Get(url, header)
	var lifetimeStats LifetimeStats
	ProcessResponseBody(res, &lifetimeStats)
	return &lifetimeStats
}

func GetLifetimeStatsWithContext(c context.Context, userId string, accessToken string) *LifetimeStats {
	url, header := MakeUrlAndHeader(GetLifetimeStatsUrl, accessToken, userId)
	res := GetWithContext(c, url, header)
	var lifetimeStats LifetimeStats
	ProcessResponseBody(res, &lifetimeStats)
	return &lifetimeStats
}