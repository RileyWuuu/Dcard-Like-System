package member

import (
	"dcard/storage/mysql"
	"encoding/json"
	"fmt"
	"net/http"
)

func delete(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	creds := &Member{}
	err := json.NewDecoder(r.Body).Decode(creds)
	ErrorCheck(err)
	delForm, err := mysql.GetMySQL().Prepare("UPDATE Member SET Dele='1' WHERE MemberID=?")
	ErrorCheck(err)
	res, err := delForm.Exec(creds.MemberID)
	ErrorCheck(err)
	id, err := res.RowsAffected()
	ErrorCheck(err)
	fmt.Println("Successfully deleted Member,ID:", id)
	http.Redirect(w, r, "/", 301)
}
