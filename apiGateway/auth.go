package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
)

type Member struct {
	MemberID      int    `json:"memberid"`
	MemberName    string `json:"membername"`
	NickName      string `json:"nickname"`
	NationalID    string `json:"nationalid"`
	DateofBirth   string `json:dateofbirth`
	Region        string `json:region`
	City          string `json:city`
	Gender        string `json:"gender"`
	ContactNumber string `json:"contactnumber"`
	UniCode       string `json:unicode`
	MajorCode     string `json:majorcode`
	Email         string `json:"email"`
	Password      string `json:"password"`
	CreateDate    string `json:createdate`
	Dele          string `json:dele`
	Male          string
	Female        string
	Paired        string `json:paired`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func MysqlConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "0000"
	dbName := "testdb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

var jwtKey = []byte("Secret")

var ctxb = context.Background()
var (
	CommentCollection *mongo.Collection
	PostCollection    *mongo.Collection
	ctx               = context.TODO()
)

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

//登入給token
func Login(w http.ResponseWriter, r *http.Request) {
	db := MysqlConn()
	creds := &Member{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Email := creds.Email
	Password := creds.Password

	selDB, err := db.Query("SELECT * FROM Member WHERE Email=? AND Password=? AND Dele=? LIMIT 1", Email, Password, "0")
	if err != nil {
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

		//json.NewEncoder(w).Encode(tkn)
		// w.Write([]byte(tokenString))
		// return
	}
	defer db.Close()
}

//驗證token正確性
func Authentication(w http.ResponseWriter, r *http.Request) string {

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

//token refresh
func Refresh(w http.ResponseWriter, r *http.Request) {
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
func main() {
	log.Println("SERVER STARTED ON: HTTP://LOCALHOST:8090")
	http.HandleFunc("/Login", Login)
	http.HandleFunc("/Refresh", Refresh)
	http.ListenAndServe(":8090", nil)
}
