package post

import (
	"dcard/storage/mysql"
	"encoding/json"
	"net/http"
)

func cancelLikeBtn(w http.ResponseWriter, r *http.Request) {
	creds := &Post{}
	err := json.NewDecoder(r.Body).Decode(creds)
	ErrorCheck(err)
	if r.Method == "POST" {
		insForm, err := mysql.GetMySQL().Prepare("UPDATE Member SET Liked = TRIM(TRAILING'?' FROM Liked)")
		ErrorCheck(err)
		res, err := insForm.Exec(creds.Id)
		res.RowsAffected()
		w.WriteHeader(http.StatusOK)
	}
}
