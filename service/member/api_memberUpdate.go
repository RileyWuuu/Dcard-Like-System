package member

import (
	"dcard/storage/mysql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func update(c *gin.Context) {
	// Authentication(w, r)
	creds := &Member{}
	err := json.NewDecoder(c.Request.Body).Decode(creds)
	ErrorCheck(err)

	insForm, err := mysql.GetMySQL().Prepare("UPDATE Member SET MemberName=?,NickName=?,NationalID=?,Region=?,City=?,Gender=?,ContactNumber=?,UniCode=?,MajorCode=?,Email=?,Password=? WHERE MemberID=?")
	ErrorCheck(err)
	res, err := insForm.Exec(creds.MemberName, creds.NickName, creds.NationalID, creds.Region, creds.City, creds.Gender, creds.ContactNumber, creds.UniCode, creds.MajorCode, creds.Email, creds.Password, creds.MemberID)
	ErrorCheck(err)
	res.RowsAffected()
	log.Println("Member info update succeed")
	c.Writer.WriteHeader(http.StatusOK)

}
