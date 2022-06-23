package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func checkUsrEmail(Email string) bool {
	db := mysqlConn()
	var isAuthenticated bool
	err := db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM Member WHERE Email = ?", Email).Scan(&isAuthenticated)
	if err != nil {
		log.Fatal(err)
	}
	return isAuthenticated
}
func insert(w http.ResponseWriter, r *http.Request) {
	db := mysqlConn()
	creds := &Member{}
	err := json.NewDecoder(r.Body).Decode(creds)
	ErrorCheck(err)
	checkUsrEmail(creds.Email)
	if r.Method == "POST" {
		insForm, err := db.Prepare("INSERT INTO Member (MemberName,NickName,NationalID,DateofBirth,Region,City,Gender,ContactNumber,UniCode,MajorCode,Email,Password,CreateDate,Dele) VALUES(?,?,?,str_to_date(?,'%Y-%m-%d') ,?,?,?,?,?,?,?,?,NOW(),'0')")
		ErrorCheck(err)
		res, err := insForm.Exec(creds.MemberName, creds.NickName, creds.NationalID, creds.DateofBirth, creds.Region, creds.City, creds.Gender, creds.ContactNumber, creds.UniCode, creds.MajorCode, creds.Email, creds.Password)
		ErrorCheck(err)
		id, err := res.LastInsertId()
		ErrorCheck(err)
		fmt.Println("Inserted New Member ID:", id)
	}

	defer db.Close()
	// http.Redirect(w, r, "/", 301)
}
