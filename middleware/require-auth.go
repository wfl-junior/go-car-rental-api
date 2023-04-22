package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	UserRepository "github.com/wfl-junior/go-car-rental-api/repositories/users"
)

func RequireAuth(context *gin.Context) {
	// get token from headers
	authorization := context.Request.Header.Get("Authorization")
	tokenString := strings.Replace(authorization, "Bearer ", "", 1)

	// decode and validate the token
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			// validate token is hashed with same algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			secret := os.Getenv("JWT_SECRET")
			return []byte(secret), nil
		},
	)

	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// validate expiration date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// get user by id from token
		user, err := UserRepository.GetById(claims["sub"].(string))

		if err != nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// attach user in request object
		context.Set("user", user)

		// let the request go through if the user is authenticated
		context.Next()
		return
	}

	// abort with Unauthorized error if anything goes wrong
	context.AbortWithStatus(http.StatusUnauthorized)
}
