package apigateway

import (
	"dcard/storage/mysql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//登入給token
func login(w http.ResponseWriter, r *http.Request) {
	creds := &Member{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Email := creds.Email
	Password := creds.Password

	selDB, err := mysql.GetMySQL().Query("SELECT * FROM Member WHERE Email=? AND Password=? AND Dele=? LIMIT 1", Email, Password, "0")
	if err != nil {
		fmt.Println(selDB)
		panic(err.Error())
	}

	for selDB.Next() {
		expirationTime := time.Now().Add(5 * time.Minute)
		// create the jwt claims, including email and expiry time
		claims := Claims{
			Email: Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		//Declare token with algorithm used for signing
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		tkn := map[string]string{
			"token": tokenString,
		}
		jsonResp, err := json.Marshal(tkn)
		if err != nil {
			log.Fatalf("Error happened in Json marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
}
