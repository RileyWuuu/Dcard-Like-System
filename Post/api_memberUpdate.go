package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	db := MysqlConn()
	creds := &Member{}
	err := json.NewDecoder(r.Body).Decode(creds)
	ErrorCheck(err)
	if r.Method == "POST" {
		insForm, err := db.Prepare("UPDATE Member SET MemberName=?,NickName=?,NationalID=?,Region=?,City=?,Gender=?,ContactNumber=?,UniCode=?,MajorCode=?,Email=?,Password=? WHERE MemberID=?")
		if err != nil {
			panic(err.Error())
		}
		res, err := insForm.Exec(creds.MemberName, creds.NickName, creds.NationalID, creds.Region, creds.City, creds.Gender, creds.ContactNumber, creds.UniCode, creds.MajorCode, creds.Email, creds.Password, creds.MemberID)
		ErrorCheck(err)
		id, err := res.RowsAffected()
		ErrorCheck(err)
		log.Println("Member info update succeed, ID:", id)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

