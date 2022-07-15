package member

import (
	"dcard/storage/mysql"
	"encoding/json"
	"log"
	"net/http"
)

func membersGet(w http.ResponseWriter, r *http.Request) {
	selDB, err := mysql.GetMySQL().Query("SELECT MemberID,MemberName, NickName, NationalID, Region, City, Gender, ContactNumber, UniCode, MajorCode, Email, Password, Dele, DateofBirth, CreateDate FROM Member WHERE Dele='0' ORDER BY MemberID")
	if err != nil {

		panic(err.Error())
	}
	mem := Member{}
	res := []Member{}
	for selDB.Next() {
		var MemberID int
		var MemberName, NickName, NationalID, Region, City, Gender, ContactNumber, UniCode, MajorCode, Email, Password, Dele, DateofBirth, CreateDate string
		err = selDB.Scan(&MemberID, &MemberName, &NickName, &NationalID, &Region, &City, &Gender, &ContactNumber, &UniCode, &MajorCode, &Email, &Password, &DateofBirth, &CreateDate, &Dele)
		if err != nil {
			panic(err.Error())
		}
		mem.MemberID = MemberID
		mem.MemberName = MemberName
		if Gender == "0" {
			mem.Gender = "男"
		} else {
			mem.Gender = "女"
		}
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
		res = append(res, mem)
	}

	jsonResp, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error happened in Json marshal. Err: %s", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

	return
}
