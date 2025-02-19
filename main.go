package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

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

			// t := time.Now()
			// f := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
			// 	t.Year(), t.Month(), t.Day(),
			// 	t.Hour(), t.Minute(), t.Second())
			// params := map[string]string{
			// 	"beforeDate": f,
			// 	"sort":       "desc",
			// 	"limit":      "100",
			// 	"offset":     "0",
			// }

			client := NewFitbitClient(tok.AccessToken)
			// activityLogList, err := client.GetActivities(c, params)
			profile, err := client.GetProfile(c)
			if err != nil {
				panic(err)
			}

			c.JSON(http.StatusOK, gin.H{
				"message": profile,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "error getting verifier",
			})
		}
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
