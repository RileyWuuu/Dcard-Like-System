package member

import (
	"dcard/storage/mysql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkUsrEmail(Email string) bool {
	var isAuthenticated bool
	err := mysql.GetMySQL().QueryRow("SELECT IF(COUNT(*),'true','false') FROM Member WHERE Email = ?", Email).Scan(&isAuthenticated)

	if err != nil {
		log.Fatal(err)
	}
	return isAuthenticated
}

func insert(c *gin.Context) {
	creds := &Member{}
	err := json.NewDecoder(c.Request.Body).Decode(creds)
	ErrorCheck(err)
	repeat := checkUsrEmail(creds.Email)
	if repeat == false {
		insForm, err := mysql.GetMySQL().Prepare("INSERT INTO Member (MemberName,NickName,NationalID,DateofBirth,Region,City,Gender,ContactNumber,UniCode,MajorCode,Email,Password,CreateDate,Dele) VALUES(?,?,?,str_to_date(?,'%Y-%m-%d') ,?,?,?,?,?,?,?,?,NOW(),'0')")
		ErrorCheck(err)
		res, err := insForm.Exec(creds.MemberName, creds.NickName, creds.NationalID, creds.DateofBirth, creds.Region, creds.City, creds.Gender, creds.ContactNumber, creds.UniCode, creds.MajorCode, creds.Email, creds.Password)
		ErrorCheck(err)
		id, err := res.LastInsertId()
		ErrorCheck(err)
		fmt.Println("Inserted New Member ID:", id)
		c.Writer.WriteHeader(http.StatusOK)
	} else {
		c.Writer.WriteHeader(http.StatusBadRequest)
		fmt.Println("This email has already been used")
	}

	// http.Redirect(w, r, "/", 301)
}
