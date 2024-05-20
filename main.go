package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func GetActivities(c *gin.Context, userId string, accessToken string) *http.Response {
	userUrl := fmt.Sprintf("https://api.fitbit.com/1/user/%s/activities/list.json", userId)
	bearerToken := fmt.Sprintf("Bearer %s", accessToken)
	client := http.Client{}
	req, err := http.NewRequestWithContext(c, "GET", userUrl, nil)
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"accept":        {"application/json"},
		"Authorization": {bearerToken},
	}

	q := req.URL.Query()
	t := time.Now()
	f := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fmt.Println(f)
	q.Add("beforeDate", f)
	q.Add("sort", "desc")
	q.Add("limit", "100")
	q.Add("offset", "0")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return res
}

func main() {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes: []string{
			"activity",
			"heartrate",
			"nutrition",
			"oxygen_saturation",
			"respiratory_rate",
			"settings",
			"sleep",
			"temperature weight",
			"profile",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  os.Getenv("AUTH_URL"),
			TokenURL: os.Getenv("ACCREF_TOKEN_URL"),
		},
	}

	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/auth", func(c *gin.Context) {
		session := sessions.Default(c)
		// use PKCE to protect against CSRF attacks
		// https://www.ietf.org/archive/id/draft-ietf-oauth-security-topics-22.html#name-countermeasures-6
		verifier := oauth2.GenerateVerifier()
		session.Set("verifier", verifier)
		session.Save()

		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
		c.IndentedJSON(http.StatusOK, gin.H{
			"url": url,
		})
	})
	r.GET("/success", func(c *gin.Context) {
		session := sessions.Default(c)
		code := c.Query("code")

		v := session.Get("verifier")

		verifier, ok := v.(string)
		if ok {
			tok, err := conf.Exchange(c, code, oauth2.VerifierOption(verifier))
			if err != nil {
				panic(err)
			}

			profile := GetActivityLogList(c, tok.Extra("user_id").(string), tok.AccessToken)

			c.JSON(http.StatusOK, gin.H{
				"message": profile,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "error",
			})
		}
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
