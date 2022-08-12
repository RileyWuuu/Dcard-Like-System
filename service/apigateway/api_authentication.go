package apigateway

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//驗證token正確性
func authentication(c *gin.Context) string {
	var Header string
	c.Header("Token", Header)
	if Header == "" {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return "401StatusUnauthorized"
	}
	tknStr := Header
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("token unauthorized"))
			panic(http.StatusUnauthorized)
		}
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.Writer.Write([]byte("Err"))
		panic(http.StatusBadRequest)
	}
	if !tkn.Valid {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Writer.Write([]byte("token invalid"))
		panic(http.StatusUnauthorized)
	}
	return tknStr
}
