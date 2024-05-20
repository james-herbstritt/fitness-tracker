package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type ActivityTime struct {
	time.Time
}

func (t *ActivityTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02T15:04:05Z07:00"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

type ActivityLevelSection struct {
	Minutes int64 `json:"minutes,omitempty"`
	Name string `json:"name,omitempty"`
}

type Activity struct {
	ActivityDuration int64 `json:"activityDuration,omitempty"`
	ActivityLevel []ActivityLevelSection `json:"activityLevel,omitempty"`
	ActivityName string `json:"activityName,omitempty"`
	ActivityTypeId int64 `json:"activityTypeId,omitempty"`
	Calories int64 `json:"calories,omitempty"`
	CaloriesLink string `json:"caloriesLink,omitempty"`
	Duration int64 `json:"duration,omitempty"`
	ElevationGain float64 `json:"elevationGain,omitempty"`
	LastModified ActivityTime `json:"lastModified,omitempty"`
	LogId int64 `json:"logId,omitempty"`
	LogType string `json:"logType,omitempty"`
	ManualValuesSpecified struct {
		Calories bool `json:"calories,omitempty"`
		Distance bool `json:"distance,omitempty"`
		Steps bool `json:"steps,omitempty"`
	}
	OriginalDuration int64 `json:"originalDuration,omitempty"`
	OriginalStartTime ActivityTime `json:"originalStartTime,omitempty"`
	StartTime ActivityTime `json:"startTime,omitempty"`
	Steps int64 `json:"steps,omitempty"`
	TcxLink string `json:"tcxLink,omitempty"`
}

type ActivityLogList struct {
	Activities []Activity `json:"activities,omitempty"`
	Pagination struct {
		AfterDate ProfileTime `json:"afterDate,omitempty"`
		Limit int64 `json:"limit,omitempty"`
		Next string `json:"next,omitempty"`
		Offset int64 `json:"offset,omitempty"`
		Previous string `json:"previous,omitempty"`
		Sort string `json:"sort,omitempty"`
	} `json:"pagination,omitempty"`
}

func GetActivityLogList(c *gin.Context, userId string, accessToken string) *ActivityLogList {
	userUrl := fmt.Sprintf("https://api.fitbit.com/1/user/%s/activities/list.json", userId)
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

	q := req.URL.Query()
	t := time.Now()
	f := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	q.Add("beforeDate", f)
	q.Add("sort", "desc")
	q.Add("limit", "100")
	q.Add("offset", "0")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var activityLogList ActivityLogList
	err = json.Unmarshal(b, &activityLogList)
	if err != nil {
		panic(err)
	}

	return &activityLogList
}