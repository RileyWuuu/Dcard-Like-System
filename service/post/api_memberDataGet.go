package post

import (
	"dcard/storage/mysql"
	"encoding/json"
	"fmt"
	"net/http"
)

func edit(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	creds := &Member{}
	if err := json.NewDecoder(r.Body).Decode(creds); err != nil {
		fmt.Println(err)
	}
	selDB, err := mysql.GetMySQL().Query("SELECT * FROM Member WHERE MemberID=? LIMIT 1", creds.MemberID)
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
		if Gender == "0" {
			mem.Male = "Checked"
			mem.Female = ""
		} else {
			mem.Male = ""
			mem.Female = "Checked"
		}
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
	}
	// tmpl.ExecuteTemplate(w, "Edit", mem)
	a, err := json.Marshal(mem)
	w.Write(a)
	return
}
