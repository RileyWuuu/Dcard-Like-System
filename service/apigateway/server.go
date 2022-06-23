package apigateway

import (
	"log"
	"net/http"
)

func EnableApiGateway() {
	log.Println("SERVER STARTED ON: HTTP://LOCALHOST:8090")
	http.HandleFunc("/Login", login)
	http.HandleFunc("/Refresh", refresh)
	http.ListenAndServe(":8090", nil)
}
