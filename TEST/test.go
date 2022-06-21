package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

var tmpl = template.Must(template.ParseGlob("form/*"))

func MysqlConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "0000"
	dbName := "testdb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func MongoConn() (db *mongo.Database) {
	host := "127.0.0.1"
	port := "27017"
	connectionURI := "mongodb://" + host + ":" + port + "/"
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	mdb := client.Database("testdb")
	PostCollection = mdb.Collection("Post")
	CommentCollection = mdb.Collection("Comment")
	return mdb
}
func RedisConn() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	ErrorCheck(err)
	fmt.Println(pong)
	return client
}

var jwtKey = []byte("Secret")

var ctxb = context.Background()
var (
	CommentCollection *mongo.Collection
	PostCollection    *mongo.Collection
	ctx               = context.TODO()
)

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

//登入給token
// func Login(w http.ResponseWriter, r *http.Request) {
// 	db := MysqlConn()
// 	creds := &Member{}
// 	err := json.NewDecoder(r.Body).Decode(creds)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	Email := creds.Email
// 	Password := creds.Password
// 	selDB, err := db.Query("SELECT * FROM Member WHERE Email=? AND Password=? AND Dele=? LIMIT 1", Email, Password, "0")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	for selDB.Next() {
// 		expirationTime := time.Now().Add(5 * time.Minute)
// 		// create the jwt claims, including email and expiry time
// 		claims := Claims{
// 			Email: Email,
// 			StandardClaims: jwt.StandardClaims{
// 				ExpiresAt: expirationTime.Unix(),
// 			},
// 		}
// 		//Declare token with algorithm used for signing
// 		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 		tokenString, err := token.SignedString(jwtKey)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		w.Header().Set("Content-Type", "application/json")
// 		tkn := map[string]string{
// 			"token": tokenString,
// 		}
// 		jsonResp, err := json.Marshal(tkn)
// 		if err != nil {
// 			log.Fatalf("Error happened in Json marshal. Err: %s", err)
// 		}
// 		w.Write(jsonResp)
// 		return

// 		//json.NewEncoder(w).Encode(tkn)
// 		// w.Write([]byte(tokenString))
// 		// return
// 	}
// 	defer db.Close()
// }

//驗證token正確性
// func Authentication(w http.ResponseWriter, r *http.Request) string {

// 	Header := r.Header.Get("Token")
// 	if Header == "" {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return "401StatusUnauthorized"
// 	}
// 	tknStr := Header
// 	claims := &Claims{}
// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("token unauthorized"))
// 			panic(http.StatusUnauthorized)
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Err"))
// 		panic(http.StatusBadRequest)
// 	}
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		w.Write([]byte("token invalid"))
// 		panic(http.StatusUnauthorized)
// 	}
// 	return tknStr
// }

//token refresh
// func Refresh(w http.ResponseWriter, r *http.Request) {
// 	Header := r.Header.Get("Token")
// 	if Header == "" {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	tknStr := Header
// 	claims := &Claims{}
// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			panic(err)
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		panic(err)
// 	}
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		panic(err)
// 	}

// 	//Token期限小於三十秒才給新的
// 	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
// 		w.WriteHeader(http.StatusBadRequest)
// 		panic(err)
// 	}

// 	expirationTime := time.Now().Add(5 * time.Minute)
// 	claims.ExpiresAt = expirationTime.Unix()

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	tknJson := map[string]string{
// 		"token": tokenString,
// 	}
// 	jsonResp, err := json.Marshal(tknJson)
// 	if err != nil {
// 		log.Fatalf("Error happened in Json marshal. Err: %s", err)
// 	}
// 	w.Write(jsonResp)
// }
func CheckUsrEmail(Email string) bool {
	db := MysqlConn()
	var isAuthenticated bool
	err := db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM Member WHERE Email = ?", Email).Scan(&isAuthenticated)
	if err != nil {
		log.Fatal(err)
	}
	return isAuthenticated
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := MysqlConn()
	selDB, err := db.Query("SELECT MemberID,MemberName, NickName, NationalID, Region, City, Gender, ContactNumber, UniCode, MajorCode, Email, Password, Dele, DateofBirth, CreateDate FROM Member WHERE Dele='0' ORDER BY MemberID")
	if err != nil {

		panic(err.Error())
	}
	mem := Member{}
	res := []Member{}
	for selDB.Next() {
		var MemberID int
		var MemberName, NickName, NationalID, Region, City, Gender, ContactNumber, UniCode, MajorCode, Email, Password, Dele, DateofBirth, CreateDate string
		err = selDB.Scan(&MemberID, &MemberName, &NickName, &NationalID, &DateofBirth, &Region, &City, &Gender, &ContactNumber, &UniCode, &MajorCode, &Email, &Password, &CreateDate, &Dele)
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
	defer db.Close()
	jsonResp, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error happened in Json marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
func Show(w http.ResponseWriter, r *http.Request) {
	db := MysqlConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Member WHERE MemberID=?", nId)
	if err != nil {
		panic(err.Error())
	}
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

	}
	tmpl.ExecuteTemplate(w, "Show", mem)
	defer db.Close()
}
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}
func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "LoginPage", nil)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	db := MysqlConn()
	creds := &Member{}
	err := json.NewDecoder(r.Body).Decode(creds)
	ErrorCheck(err)
	delForm, err := db.Prepare("UPDATE Member SET Dele='1' WHERE MemberID=?")
	ErrorCheck(err)
	res, err := delForm.Exec(creds.MemberID)
	ErrorCheck(err)
	id, err := res.RowsAffected()
	ErrorCheck(err)
	fmt.Println("Successfully deleted Member,ID:", id)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
func Insert(w http.ResponseWriter, r *http.Request) {
	db := MysqlConn()
	creds := &Member{}
	err := json.NewDecoder(r.Body).Decode(creds)
	ErrorCheck(err)
	CheckUsrEmail(creds.Email)
	if r.Method == "POST" {
		insForm, err := db.Prepare("INSERT INTO Member (MemberName,NickName,NationalID,DateofBirth,Region,City,Gender,ContactNumber,UniCode,MajorCode,Email,Password,CreateDate,Dele) VALUES(?,?,?,str_to_date(?,'%Y-%m-%d') ,?,?,?,?,?,?,?,?,NOW(),'0')")
		ErrorCheck(err)
		res, err := insForm.Exec(creds.MemberName, creds.NickName, creds.NationalID, creds.DateofBirth, creds.Region, creds.City, creds.Gender, creds.ContactNumber, creds.UniCode, creds.MajorCode, creds.Email, creds.Password)
		ErrorCheck(err)
		id, err := res.LastInsertId()
		ErrorCheck(err)
		fmt.Println("Inserted New Member ID:", id)
	}

	defer db.Close()
	// http.Redirect(w, r, "/", 301)
}
func Edit(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	db := MysqlConn()
	creds := &Member{}
	err := json.NewDecoder(r.Body).Decode(creds)
	selDB, err := db.Query("SELECT * FROM Member WHERE MemberID=? LIMIT 1", creds.MemberID)
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
	defer db.Close()
	a, err := json.Marshal(mem)
	w.Write(a)
	return
}
func Update(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	db := MysqlConn()
	creds := &Member{}
	err := json.NewDecoder(r.Body).Decode(creds)
	ErrorCheck(err)
	if r.Method == "POST" {
		insForm, err := db.Prepare("UPDATE Member SET MemberName=?,NickName=?,NationalID=?,Region=?,City=?,Gender=?,ContactNumber=?,UniCode=?,MajorCode=?,Email=?,Password=? WHERE MemberID=?")
		if err != nil {
			panic(err.Error())
		}
		res, err := insForm.Exec(creds.MemberName, creds.NickName, creds.NationalID, creds.Region, creds.City, creds.Gender, creds.ContactNumber, creds.UniCode, creds.MajorCode, creds.Email, creds.Password, creds.MemberID)
		ErrorCheck(err)
		id, err := res.RowsAffected()
		ErrorCheck(err)
		log.Println("Member info update succeed, ID:", id)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	db := MongoConn()
	client := RedisConn()
	PostCollection = db.Collection("Post")
	pst := &Post{}
	err := json.NewDecoder(r.Body).Decode(pst)
	pst.PostDate = time.Now()
	pst.Id = ""
	now := time.Now()
	timestamp := float64(now.Unix())
	result, err := PostCollection.InsertOne(ctx, pst)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ContentTune := []rune(pst.Content)
	if len(pst.Content) > 30 {
		pst.Content = string(ContentTune[:30])
	}
	resultid := result.InsertedID
	ID := resultid.(primitive.ObjectID).Hex()

	post := map[string]string{
		"ID":      ID,
		"Title":   pst.Title,
		"Content": pst.Content,
		"Likes":   strconv.Itoa(pst.Likes),
	}
	PJson, err := json.Marshal(post)
	postString := PJson
	fmt.Println("postStringpostString", postString)

	_, errr := client.ZAdd("Posts", redis.Z{timestamp, PJson}).Result()
	ErrorCheck(errr)

	w.Write([]byte(fmt.Sprintf("%v", result.InsertedID)))
	return
}
func GetPost(w http.ResponseWriter, r *http.Request) {
	db := MongoConn()
	var p Post
	pst := &Post{}
	err := json.NewDecoder(r.Body).Decode(pst)
	PostCollection = db.Collection("Post")
	objectid, err := primitive.ObjectIDFromHex(pst.Id)
	fmt.Println(pst)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = PostCollection.FindOne(ctx, bson.D{{"_id", objectid}}).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonResp, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Error happened in Json marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
func GetPosts(w http.ResponseWriter, r *http.Request) {
	var post PostSummary
	var posts []PostSummary
	var posts2 []PostSummary
	db := MongoConn()
	rdb := RedisConn()
	page := &Posts{}
	err := json.NewDecoder(r.Body).Decode(page)
	ErrorCheck(err)
	fmt.Println(page)
	total := page.Page * page.PerPage
	total = total - 1
	result := rdb.ZRevRange("Posts", 0, int64(total))
	ErrorCheck(err)
	aaa := result.Val()
	i := 0
	for _, count := range aaa {
		i = i + 1
		arr := strings.Split(count, ",")

		post.Content = arr[0]
		post.Id = arr[1]
		post.Likes, _ = strconv.Atoi(arr[2])
		post.Title = arr[3]
		posts = append(posts, post)
	}

	if i != total {
		fmt.Println(i, total)
		// k := total - i
		PostCollection = db.Collection("Post")
		cursor, err := PostCollection.Find(ctx, bson.D{})
		// cursor := cursor.SetSort(bson.D{{"_id",-1}})
		if err != nil {
			defer cursor.Close(ctx)
			fmt.Println("ERROR")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for cursor.Next(ctx) {
			err := cursor.Decode(&post)
			fmt.Println(err)
			if err != nil {
				fmt.Println("ERROR")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			ContentTune := []rune(post.Content)
			if len(post.Content) > 50 {
				post.Content = string(ContentTune[:50])
			}
			posts2 = append(posts2, post)
			fmt.Println("POST2", post)
		}
	}

	// for cursor.Next(ctx) {
	// 	err := cursor.Decode(&post)
	// 	fmt.Println(err)
	// 	if err != nil {
	// 		fmt.Println("ERROR")
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// 	ContentTune := []rune(post.Content)
	// 	if len(post.Content) > 50 {
	// 		post.Content = string(ContentTune[:50])
	// 	}
	// 	posts = append(posts, post)
	// 	fmt.Println("POST", post)
	// }
	jsonResp, err := json.Marshal(posts2)
	if err != nil {
		log.Fatalf("Error happened in Json marshal. Err: %s", err)
	}
	fmt.Println("posts2:", posts2)

	w.Write(jsonResp)
	return
}
func AddLike(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	db := MongoConn()
	post := &Post{}
	PostCollection = db.Collection("Post")
	err := json.NewDecoder(r.Body).Decode(post)
	objectid, err := primitive.ObjectIDFromHex(post.Id)
	err = PostCollection.FindOne(ctx, bson.D{{"_id", objectid}}).Decode(&post)
	result, err := PostCollection.UpdateOne(
		ctx,
		bson.M{"_id": objectid},
		bson.D{
			{"$set", bson.D{{"likes", (post.Likes + 1)}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Added like : %v \n", result.ModifiedCount)
}
func GetComments(w http.ResponseWriter, r *http.Request) {
	var condition bson.D
	db := MongoConn()
	cmt := &Comment{}
	var comment Comment
	var comments []Comment
	err := json.NewDecoder(r.Body).Decode(cmt)
	condition = append(condition, bson.E{Key: "postid", Value: cmt.PostID})
	fmt.Println(cmt)
	CommentCollection = db.Collection("Comment")
	cursor, err := CommentCollection.Find(ctx, condition)
	if err != nil {
		defer cursor.Close(ctx)
		fmt.Println("ERROR")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(&comment)
		fmt.Println(err)
		if err != nil {
			fmt.Println("ERROR")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		comments = append(comments, comment)
	}
	jsonResp, err := json.Marshal(comments)
	if err != nil {
		log.Fatalf("Error happened in Json marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
func PostComment(w http.ResponseWriter, r *http.Request) {
	// Authentication(w, r)
	db := MongoConn()
	PostCollection = db.Collection("Post")
	CommentCollection = db.Collection("Comment")
	cmt := &Comment{}
	err := json.NewDecoder(r.Body).Decode(cmt)
	cmt.CommentDate = time.Now()
	result, err := CommentCollection.InsertOne(ctx, cmt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(fmt.Sprintf("%v", result.InsertedID)))
	return
}

// func Matching(w http.ResponseWriter, r *http.Request) {
// 	db := MysqlConn()

// 	var mem = MemID{}
// 	var mems = []MemID{}
// 	// var PairingList = []MemID{}
// 	// var pairedList = []MemID{}
// 	FindMale, err := db.Query("SELECT MemberID,Paired FROM Member WHERE Gender='0' AND Dele='0' ORDER BY MemberID")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	for FindMale.Next() {
// 		var MemberID int
// 		var Pair json.RawMessage
// 		err = FindMale.Scan(&MemberID, &Pair)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		mem.Male = MemberID
// 		mems = append(mems, mem)
// 		// Malemems[rand.Intn(len(Malemems))]
// 		rand.Seed(time.Now().UnixNano())
// 		rand.Shuffle(len(mems), func(i, j int) {
// 			mems[i], mems[j] = mems[j], mems[i]
// 		})

// 	}
// 	FindFemale, err := db.Query("SELECT MemberID,Paired FROM Member WHERE Gender='1' AND Dele='0' ORDER BY MemberID")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	i := 0
// 	for FindFemale.Next() {
// 		var MemberID int
// 		var Pair []uint8
// 		var PairList []string

// 		err = FindFemale.Scan(&MemberID, &Pair)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		err = json.Unmarshal(Pair, &PairList)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		mem.Female = MemberID
// 		mems[i].Female = mem.Female
// 		mems[i].Pair = PairList

// 		i = i + 1
// 	}
// 	fmt.Println("Result", mems)
// 	// copy(PairingList, mems
// 	InvalidList := []MemID{}
// 	ValidList := []MemID{}
// 	for _, ID := range mems {
// 		result := PairingCheck(strconv.Itoa(ID.Male), ID.Pair)
// 		fmt.Println("CheckPair", result)
// 		item := MemID{Male: ID.Male, Female: ID.Female, Pair: ID.Pair}
// 		if result == true {
// 			InvalidList = append(InvalidList, item)
// 		} else {
// 			ValidList = append(ValidList, item)
// 		}
// 	}
// 	fmt.Println("InvalidList", InvalidList)
// 	fmt.Println("ValidList", ValidList)
// 	InvalidL := []MemID{}
// 	for _, i := range InvalidList {
// 		vl := MemID{}
// 		vl.Male = i.Male
// 		InvalidL = append(InvalidL, vl)
// 		// Malemems[rand.Intn(len(Malemems))]
// 		rand.Seed(time.Now().UnixNano())
// 		rand.Shuffle(len(InvalidL), func(i, j int) {
// 			InvalidL[i], InvalidL[j] = InvalidL[j], InvalidL[i]
// 		})
// 	}
// 	var s = 0
// 	for _, i := range InvalidList {
// 		InvalidL[s].Female = i.Female
// 		InvalidL[s].Pair = i.Pair
// 		s++
// 	}
// 	fmt.Println("InvalidL", InvalidL)
// 	InvalidList = nil
// 	for _, ID := range InvalidL {
// 		result := PairingCheck(strconv.Itoa(ID.Male), ID.Pair)
// 		fmt.Println("CheckPair", result)
// 		item := MemID{Male: ID.Male, Female: ID.Female, Pair: ID.Pair}
// 		if result == true {
// 			InvalidList = append(InvalidList, item)
// 		} else {
// 			ValidList = append(ValidList, item)
// 		}
// 	}
// 	fmt.Println("InvalidList", InvalidList)
// 	for _, ID := range ValidList {
// 		InsertRecordM, err := db.Prepare("INSERT INTO MatchingRecord (MemberID,MatchedWith,Request,MatchedDate) Values(?,?,0,NOW())")
// 		InsertRecordF, err := db.Prepare("INSERT INTO MatchingRecord (MemberID,MatchedWith,Request,MatchedDate) Values (?,?,0,NOW())")
// 		ErrorCheck(err)
// 		InsertRecordM.Exec(ID.Female, ID.Male)
// 		InsertRecordF.Exec(ID.Male, ID.Female)
// 	}
// 	fmt.Println("ValidList", ValidList)
// 	//batch update
// 	defer db.Close()

// 	return
// }

// //檢查是否已配對過
// func PairingCheck(a string, list []string) bool {
// 	for _, b := range list {
// 		if b == a {
// 			return true
// 		}
// 	}
// 	return false
// }

func main() {
	log.Println("SERVER STARTED ON: HTTP://LOCALHOST:8091")
	http.HandleFunc("/", Index)
	http.HandleFunc("/Insert", Insert)
	http.HandleFunc("/Update", Update)
	http.HandleFunc("/Show", Show)
	http.HandleFunc("/New", New)
	http.HandleFunc("/Edit", Edit)
	http.HandleFunc("/Delete", Delete)
	http.HandleFunc("/LoginPage", LoginPage)
	http.HandleFunc("/GetPost", GetPost)
	http.HandleFunc("/GetPosts", GetPosts)
	http.HandleFunc("/CreatePost", CreatePost)
	http.HandleFunc("/PostComment", PostComment)
	http.HandleFunc("/GetComments", GetComments)
	http.HandleFunc("/AddLike", AddLike)
	// http.HandleFunc("/RedisConn", RedisConn)
	// http.HandleFunc("/SendRequest", SendRequest)
	// http.HandleFunc("/Matching", Matching)
	http.ListenAndServe(":8091", nil)
}
