package middleware

import (
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func IsTokenExpired(token string) bool {
	parsedToken, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return true
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return true
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return true
	}

	expTime := time.Unix(int64(exp), 0)
	return expTime.Before(time.Now())
}

func RefreshToken(token string, refreshToken string) (string, error) {
	clientId := os.Getenv("CLIENT_ID")
	// refreshToken := grab the refresh token from where you stored it

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	data.Set("client_id", clientId)

	req, _ := http.NewRequest("POST", "https://api.fibit.com/oauth2/token", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
}
