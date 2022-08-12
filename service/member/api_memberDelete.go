package member

import (
	"dcard/storage/mysql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func delete(c *gin.Context) {
	// Authentication(w, r)
	creds := &Member{}
	err := json.NewDecoder(c.Request.Body).Decode(creds)
	ErrorCheck(err)
	delForm, err := mysql.GetMySQL().Prepare("UPDATE Member SET Dele='1' WHERE MemberID=?")
	ErrorCheck(err)
	res, err := delForm.Exec(creds.MemberID)
	ErrorCheck(err)
	id, err := res.RowsAffected()
	ErrorCheck(err)
	fmt.Println("Successfully deleted Member,ID:", id)
	c.Writer.WriteHeader(http.StatusOK)

}
