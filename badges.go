package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
)

type BadgeList struct {
	Badges []Badge `json:"badges,omitempty"`
} 

type Badge struct {
	BadgeGradientEndColor string `json:"badgeGradientEndColor,omitempty"`
	BadgeGradientStartColor string `json:"badgeGradientStartColor,omitempty"`
	BadgeType string `json:"badgeType,omitempty"`
	Category string `json:"category,omitempty"`
	Cheers []string `json:"cheers,omitempty"`
	DateTime ProfileTime `json:"dateTime,omitempty"`
	Description string `json:"description,omitempty"`
	EarnedMessage string `json:"earnedMessage,omitempty"`
	EncodedId string `json:"encodedId,omitempty"`
	Image100px string `json:"image100px,omitempty"`
	Image125px string `json:"image125px,omitempty"`
	Image300px string `json:"image300px,omitempty"`
	Image50px string `json:"image50px,omitempty"`
	Image75px string `json:"image75px,omitempty"`
	MarketingDescription string `json:"marketingDescription,omitempty"`
	MobileDescription string `json:"mobileDescription,omitempty"`
	Name string `json:"name,omitempty"`
	ShareImage640px string `json:"shareImage640px,omitempty"`
	ShareText string `json:"shareText,omitempty"`
	ShortDescription string `json:"shortDescription,omitempty"`
	ShortName string `json:"shortName,omitempty"`
	TimesAchieved int64 `json:"timesAchieved,omitempty"`
	Value int64 `json:"value,omitempty"`
}

func GetBadges(c *gin.Context, userId string, accessToken string) *BadgeList {
	userUrl := fmt.Sprintf("https://api.fitbit.com/1/user/%s/badges.json", userId)
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

	var badgeList BadgeList
	err = json.Unmarshal(b, &badgeList)
	if err != nil {
		panic(err)
	}

	return &badgeList
}