package matching

import (
	"dcard/storage/mysql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func matching(w http.ResponseWriter, r *http.Request) {
	var mem = MemID{}
	var mems = []MemID{}
	// var PairingList = []MemID{}
	// var pairedList = []MemID{}
	FindMale, err := mysql.GetMySQL().Query("SELECT MemberID,Paired FROM Member WHERE Gender='0' AND Dele='0' ORDER BY MemberID")
	if err != nil {
		panic(err.Error())
	}
	for FindMale.Next() {
		var MemberID int
		var Pair json.RawMessage
		err = FindMale.Scan(&MemberID, &Pair)
		if err != nil {
			panic(err.Error())
		}
		mem.Male = MemberID
		mems = append(mems, mem)
		// Malemems[rand.Intn(len(Malemems))]
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(mems), func(i, j int) {
			mems[i], mems[j] = mems[j], mems[i]
		})

	}

	FindFemale, err := mysql.GetMySQL().Query("SELECT MemberID,Paired FROM Member WHERE Gender='1' AND Dele='0' ORDER BY MemberID")
	if err != nil {
		panic(err.Error())
	}
	i := 0
	for FindFemale.Next() {
		var MemberID int
		var Pair []uint8
		var PairList []string

		err = FindFemale.Scan(&MemberID, &Pair)
		if err != nil {
			panic(err.Error())
		}
		err = json.Unmarshal(Pair, &PairList)
		if err != nil {
			fmt.Println(err)
		}
		mem.Female = MemberID
		mems[i].Female = mem.Female
		mems[i].Pair = PairList

		i = i + 1
	}
	fmt.Println("Result", mems)
	// copy(PairingList, mems

	InvalidList := []MemID{}
	ValidList := []MemID{}
	for _, ID := range mems {
		result := pairingCheck(strconv.Itoa(ID.Male), ID.Pair)
		fmt.Println("CheckPair", result)
		item := MemID{Male: ID.Male, Female: ID.Female, Pair: ID.Pair}
		if result == true {
			InvalidList = append(InvalidList, item)
		} else {
			ValidList = append(ValidList, item)
		}
	}
	fmt.Println("InvalidList", InvalidList)
	fmt.Println("ValidList", ValidList)
	InvalidL := []MemID{}
	for _, i := range InvalidList {
		vl := MemID{}
		vl.Male = i.Male
		InvalidL = append(InvalidL, vl)
		// Malemems[rand.Intn(len(Malemems))]
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(InvalidL), func(i, j int) {
			InvalidL[i], InvalidL[j] = InvalidL[j], InvalidL[i]
		})
	}

	var s = 0
	for _, i := range InvalidList {
		InvalidL[s].Female = i.Female
		InvalidL[s].Pair = i.Pair
		s++
	}
	fmt.Println("InvalidL", InvalidL)

	InvalidList = nil
	for _, ID := range InvalidL {
		result := pairingCheck(strconv.Itoa(ID.Male), ID.Pair)
		fmt.Println("CheckPair", result)
		item := MemID{Male: ID.Male, Female: ID.Female, Pair: ID.Pair}
		if result == true {
			InvalidList = append(InvalidList, item)
		} else {
			ValidList = append(ValidList, item)
		}
	}
	fmt.Println("InvalidList", InvalidList)

	for _, ID := range ValidList {
		InsertRecordM, err := mysql.GetMySQL().Prepare("INSERT INTO MatchingRecord (MemberID,MatchedWith,Request,MatchedDate) Values(?,?,0,NOW())")
		errorCheck(err)
		InsertRecordF, err := mysql.GetMySQL().Prepare("INSERT INTO MatchingRecord (MemberID,MatchedWith,Request,MatchedDate) Values (?,?,0,NOW())")
		errorCheck(err)

		InsertRecordM.Exec(ID.Female, ID.Male)
		InsertRecordF.Exec(ID.Male, ID.Female)
	}
	fmt.Println("ValidList", ValidList)

	return
}

//檢查是否已配對過
func pairingCheck(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func errorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
