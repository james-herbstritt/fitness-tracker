package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/james-herbstritt/fitness-tracker/fitbit"
	"github.com/james-herbstritt/fitness-tracker/internal"
	"golang.org/x/oauth2"
)

func JSONResponse(c *gin.Context, code int, obj interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)

	encoder := json.NewEncoder(c.Writer)
	encoder.SetEscapeHTML(false) // Prevent escaping
	_ = encoder.Encode(obj)      // Encode JSON and write to response
}

func Run() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/auth", func(c *gin.Context) {
		session := sessions.Default(c)
		// use PKCE to protect against CSRF attacks
		// https://www.ietf.org/archive/id/draft-ietf-oauth-security-topics-22.html#name-countermeasures-6
		verifier := oauth2.GenerateVerifier()
		session.Set("verifier", verifier)
		session.Save()

		url := internal.GenerateAuthCodeURL(verifier)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"AuthURL": url,
		})

	})
	r.GET("/success", func(c *gin.Context) {
		session := sessions.Default(c)
		code := c.Query("code")
		v := session.Get("verifier")

		verifier, ok := v.(string)
		if ok {
			tok, err := internal.ExchangeAuthCode(c, code, verifier)
			if err != nil {
				panic(err)
			}

			tokenSource := internal.GetTokenSource(c, tok)
			token, err := internal.GetToken(tokenSource)

			client := fitbit.NewFitbitClient(token.AccessToken)
			lifetimeStats, err := client.GetLifetimeStats(c, "-")
			profile, err := client.GetProfile(c, "-")

			if err != nil {
				panic(err)
			}

			c.HTML(http.StatusOK, "profile.tmpl", gin.H{
				"Name":        profile.User.DisplayName,
				"Avatar":      profile.User.Avatar,
				"MemberSince": profile.User.MemberSince,
				"Steps":       lifetimeStats.Lifetime.Total.Steps,
				"Distance":    lifetimeStats.Lifetime.Total.Distance,
				"Floors":      lifetimeStats.Lifetime.Total.Floors,
			})
		}
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
