package apigateway

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//token refresh
func refresh(c *gin.Context) {
	var Header string
	c.Header("Token", Header)
	if Header == "" {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	tknStr := Header
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			panic(err)
		}
		c.Writer.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	if !tkn.Valid {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		panic(err)
	}

	//Token期限小於三十秒才給新的
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		c.Writer.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/json")
	tknJson := map[string]string{
		"token": tokenString,
	}
	jsonResp, err := json.Marshal(tknJson)
	if err != nil {
		log.Fatalf("Error happened in Json marshal. Err: %s", err)
	}
	c.Writer.Write(jsonResp)
}
