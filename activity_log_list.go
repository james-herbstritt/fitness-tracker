package main

import (
	"context"
	"time"
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

func GetActivityLogList(userId string, accessToken string, params map[string]string) *ActivityLogList {
	url, header := MakeUrlAndHeader(GetActivityLogListUrl, accessToken, userId)
	res := GetWithQueryParams(url, header, params)
	var activityLogList ActivityLogList
	ProcessResponseBody(res, &activityLogList)
	return &activityLogList
}

func GetActivityLogListWithContext(c context.Context, userId string, accessToken string, params map[string]string) *ActivityLogList {
	url, header := MakeUrlAndHeader(GetActivityLogListUrl, accessToken, userId)
	res := GetWithContextWithQueryParams(c, url, header, params)
	var activityLogList ActivityLogList
	ProcessResponseBody(res, &activityLogList)
	return &activityLogList
}