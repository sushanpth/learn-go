package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sushanpth/learn-go/go-jwt/initializers"
	"github.com/sushanpth/learn-go/go-jwt/models"
)

func RequireAuth(c *gin.Context) {

	// get the cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil || len(tokenString) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// decode and validate
	// Parse takes the token string and a function for looking up the key.
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// find user with token sub
		var user models.User
		results := initializers.DB.First(&user, claims["sub"])

		if results.Error != nil || user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// attach to request
		c.Set("user", user)

		// continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}
