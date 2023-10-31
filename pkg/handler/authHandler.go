package handler

import (
	"GoUber/pkg/user"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

const secretKey = "supersecretkey"

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			var jwtToken = request.Header["Token"][0]
			var userClaim user.UserClaim
			token, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})
			if err != nil {
				ErrorJson(writer, err, http.StatusBadRequest)
				return
			}
			if !token.Valid {
				ErrorJson(writer, fmt.Errorf("invalid token"), http.StatusBadRequest)
				return
			}
			endpointHandler(writer, request)
		} else {

			ErrorJson(writer, fmt.Errorf("missing token"), http.StatusBadRequest)
			return
		}
	}
}
