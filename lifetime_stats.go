package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
			// These two are always going to be -1, just here for backwards compatibility
			ActiveScore int64 `json:"activeScore"`
			CaloriesOut int64 `json:"caloriesOut"`

			Distance float64 `json:"distance"`
			Floors float64 `json:"floors"`
			Steps int64 `json:"steps"`
		}
	} `json:"lifetime"`
}

func GetLifetimeStats(c *gin.Context, userId string, accessToken string) *LifetimeStats {
	userUrl := fmt.Sprintf("https://api.fitbit.com/1/user/%s/activities.json", userId)
	bearerToken := fmt.Sprintf("Bearer %s", accessToken)
	client := http.Client{}
	req, err := http.NewRequestWithContext(c, "GET", userUrl, nil)
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"accept":        {"application/json"},
		"accept-langauge": {"en_US"},
		"accept-locale": {"en_US"},
		"Authorization": {bearerToken},
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var lifetimeStats LifetimeStats
	err = json.Unmarshal(b, &lifetimeStats)
	if err != nil {
		panic(err)
	}

	return &lifetimeStats
}