package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

//驗證token正確性
func authentication(w http.ResponseWriter, r *http.Request) string {
	Header := r.Header.Get("Token")
	if Header == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return "401StatusUnauthorized"
	}
	tknStr := Header
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token unauthorized"))
			panic(http.StatusUnauthorized)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Err"))
		panic(http.StatusBadRequest)
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("token invalid"))
		panic(http.StatusUnauthorized)
	}
	return tknStr
}
