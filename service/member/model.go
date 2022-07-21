package member

import (
	"github.com/dgrijalva/jwt-go"
)

type Member struct {
	MemberID      int    `json:"memberid"`
	MemberName    string `json:"membername"`
	NickName      string `json:"nickname"`
	NationalID    string `json:"nationalid"`
	DateofBirth   string `json:"dateofbirth"`
	Region        string `json:"region"`
	City          string `json:"city"`
	Gender        string `json:"gender"`
	ContactNumber string `json:"contactnumber"`
	UniCode       string `json:"unicode"`
	MajorCode     string `json:"majorcode"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	CreateDate    string `json:"createdate"`
	Dele          string `json:"dele"`
	Paired        string `json:"paired"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
