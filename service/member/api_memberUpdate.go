package member

import (
	"dcard/storage/mysql"
	"encoding/json"
	"log"
	"net/http"
)

func update(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	creds := &Member{}
	err := json.NewDecoder(r.Body).Decode(creds)
	ErrorCheck(err)
	if r.Method == "POST" {
		insForm, err := mysql.GetMySQL().Prepare("UPDATE Member SET MemberName=?,NickName=?,NationalID=?,Region=?,City=?,Gender=?,ContactNumber=?,UniCode=?,MajorCode=?,Email=?,Password=? WHERE MemberID=?")
		if err != nil {
			panic(err.Error())
		}
		res, err := insForm.Exec(creds.MemberName, creds.NickName, creds.NationalID, creds.Region, creds.City, creds.Gender, creds.ContactNumber, creds.UniCode, creds.MajorCode, creds.Email, creds.Password, creds.MemberID)
		ErrorCheck(err)
		res.RowsAffected()
		log.Println("Member info update succeed")
	}
	w.WriteHeader(http.StatusOK)

}
