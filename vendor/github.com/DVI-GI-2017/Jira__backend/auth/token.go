package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

type Token struct {
	Token string `json:"token" bson:"token"`
}

func NewToken() (Token, error) {
	token := jwt.New(jwt.SigningMethodRS256)

	claims := make(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 730).Unix()
	claims["iat"] = time.Now().Unix()

	token.Claims = claims

	result, err := token.SignedString(signKey)

	if err != nil {
		return Token{Token: result}, fmt.Errorf("can not create signed string: %v", err)
	}

	return Token{result}, nil
}

func ValidateToken(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return verifyKey, nil
			})

		if err == nil {
			if token.Valid {
				h(w, r)

				return
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Bad token")
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Unauthorized")
	}
}
