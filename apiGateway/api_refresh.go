package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//token refresh
func refresh(w http.ResponseWriter, r *http.Request) {
	Header := r.Header.Get("Token")
	if Header == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tknStr := Header
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			panic(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		panic(err)
	}

	//Token期限小於三十秒才給新的
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	tknJson := map[string]string{
		"token": tokenString,
	}
	jsonResp, err := json.Marshal(tknJson)
	if err != nil {
		log.Fatalf("Error happened in Json marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
