package member

import (
	"dcard/storage/mysql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func singleMemberGet(c *gin.Context) {
	// Authentication(w, r)
	var status int
	creds := &Member{}
	fmt.Println("ERROR")
	
	if err := json.NewDecoder(c.Request.Body).Decode(creds); err != nil {
		fmt.Println(err)
	}
	selDB, err := mysql.GetMySQL().Query("SELECT MemberID,MemberName, NickName, NationalID, DateofBirth, Region, City, Gender, ContactNumber, UniCode, MajorCode, Email, Password, CreateDate, Dele FROM Member WHERE MemberID=? AND Dele = '0' LIMIT 1", creds.MemberID)
	mem := Member{}
	for selDB.Next() {
		var MemberID int
		var MemberName, NickName, NationalID, DateofBirth, Region, City, Gender, ContactNumber, UniCode, MajorCode, Email, Password, CreateDate, Dele string
		err = selDB.Scan(&MemberID, &MemberName, &NickName, &NationalID, &DateofBirth, &Region, &City, &Gender, &ContactNumber, &UniCode, &MajorCode, &Email, &Password, &CreateDate, &Dele)
		if err != nil {
			panic(err.Error())
		}
		mem.MemberID = MemberID
		mem.MemberName = MemberName
		mem.Gender = Gender
		mem.NickName = NickName
		mem.NationalID = NationalID
		mem.DateofBirth = DateofBirth
		mem.Region = Region
		mem.City = City
		mem.ContactNumber = ContactNumber
		mem.UniCode = UniCode
		mem.MajorCode = MajorCode
		mem.Email = Email
		mem.Password = Password
		mem.CreateDate = CreateDate
		mem.Dele = Dele
	}
	// tmpl.ExecuteTemplate(w, "Edit", mem)
	a, err := json.Marshal(mem)
	if err != nil {
		status = http.StatusBadRequest
	} else {
		status = http.StatusOK
	}
	if mem.MemberID == 0 {
		a, err = json.Marshal("查無資料")
		status = http.StatusBadRequest
	}
	c.Writer.Write(a)
	c.Writer.WriteHeader(status)
	// w.WriteHeader(status)
	// w.Write(a)
	// return
}
