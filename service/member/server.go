package member

import (
	"log"
	"net/http"
)

func EnablePostServer() {
	log.Println("SERVER STARTED ON: HTTP://LOCALHOST:8093")
	// http.HandleFunc("/", index)
	http.HandleFunc("/member_insert", insert)
	http.HandleFunc("/Update", update)
	http.HandleFunc("/member_delete", delete)
	http.HandleFunc("/members_get", membersGet)
	http.HandleFunc("/member_get", singleMemberGet)
	http.ListenAndServe(":8091", nil)
}