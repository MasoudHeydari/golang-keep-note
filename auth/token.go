package auth

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var apiSecret = []byte(os.Getenv("API_SECRET"))

func CreateToken(userId uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	// token will expire after 1 hour
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(apiSecret)
}

func TokenValidation(r *http.Request) error {
	tokenStr := ExtractToken(r)
	token, err := jwtParseToken(tokenStr)
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		pretty(claims)
	}

	// token is valid so return nil!
	return nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}

	bearerToken := r.Header.Get("Authorization")
	bearerTokenSlice := strings.Split(bearerToken, " ")
	if len(bearerTokenSlice) == 2 {
		return bearerTokenSlice[1]
	}

	return ""
}

func ExtractTokenID(r *http.Request) (uint32, error) {
	tokenStr := ExtractToken(r)
	token, err := jwtParseToken(tokenStr)

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseInt(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}

	return 0, nil
}

func jwtParseToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing error, method: %v", token.Header["alg"])
		}

		return apiSecret, nil
	})
}

func pretty(data interface{}) {
	marshaledData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(marshaledData))
}
