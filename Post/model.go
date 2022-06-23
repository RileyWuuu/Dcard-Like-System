package post

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Member struct {
	MemberID      int    `json:"memberid"`
	MemberName    string `json:"membername"`
	NickName      string `json:"nickname"`
	NationalID    string `json:"nationalid"`
	DateofBirth   string `json:dateofbirth`
	Region        string `json:region`
	City          string `json:city`
	Gender        string `json:"gender"`
	ContactNumber string `json:"contactnumber"`
	UniCode       string `json:unicode`
	MajorCode     string `json:majorcode`
	Email         string `json:"email"`
	Password      string `json:"password"`
	CreateDate    string `json:createdate`
	Dele          string `json:dele`
	Male          string
	Female        string
	Paired        string `json:paired`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Post struct {
	Id       string    `json:"_id" bson:"_id,omitempty"`
	MemberID int       `json:"memberid"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	FileLink string    `json:"filelink"`
	Likes    int       `json:"likes"`
	PostDate time.Time `json:"postdate"`
}
type PostSummary struct {
	Content string `json:"content"`
	Id      string `json:"_id" bson:"_id,omitempty"`
	Likes   int    `json:"likes"`
	Title   string `json:"title"`
}

type Comment struct {
	Id          string    `json:"id,omitempty" bson:"id,omitempty"`
	PostID      string    `json:"postid"`
	MemberID    int       `json:"memberid"`
	Comment     string    `json:"comment"`
	FileLink    string    `json:"filelink"`
	Likes       int       `json:"like"`
	CommentDate time.Time `json:"commentdate"`
}

type MemberPairing struct {
	Value string `json:"value"`
}
type MemID struct {
	Male   int
	Female int
	Pair   []string
}
type Posts struct {
	Page    int `json:"page"`
	PerPage int `json:"perpage"`
}
