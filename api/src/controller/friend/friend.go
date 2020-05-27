package friend

import (
	"encoding/json"
	"fmt"
	"net/http"
	db "src/db"
	friend "src/model/friend"
	rep "src/model/reponse"
)

func AddFriendEmail(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response rep.Response
	response.Success = false

	var friendReq friend.FriendRequest
	json.NewDecoder(req.Body).Decode(&friendReq)

	if len(friendReq.Friends) == 2 {
		db := db.DBConn()
		for _, user := range friendReq.Friends {
			result, err := db.Prepare("INSERT INTO users(email) values(?)")
			check(err)
			result.Exec(user)
			fmt.Println(user)

		}
		ins, err := db.Prepare("INSERT INTO relationship(user1,user2,blocker) VALUES(?,?,?)")
		check(err)
		ins.Exec(friendReq.Friends[0], friendReq.Friends[1], 1)

		defer db.Close()
	}
	response.Success = true
	json.NewEncoder(w).Encode(response)

}

func ShowListFriendEmail(w http.ResponseWriter, req *http.Request) {

	emails, ok := req.URL.Query()["email"]

	if !ok || len(emails[0]) < 1 {
		return
	}

	email := emails[0]
	var friendResult friend.FriendConnection
	friendResult.Success = false
	fmt.Println("hi", email)
	db := db.DBConn()
	defer db.Close()

	result, err := db.Query(`SELECT user2 from relationship where user1 = ? AND blocker= ?;`, email, 1)
	check(err)
	for result.Next() {
		var emails string
		err = result.Scan(&emails)
		check(err)
		friendResult.Friends = append(friendResult.Friends, emails)
		friendResult.Count++
	}

	friendResult.Success = true
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friendResult)

}

func ShowMutualFriendEmail(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var friendResult friend.FriendConnection
	friendResult.Success = false
	json.NewDecoder(req.Body).Decode(&friendResult)

	if len(friendResult.Friends) > 1 {
		db := dbConn()
		defer db.Close()
		sqlStatement := `select fr1.user2 
		from relationship fr1 
		inner join relationship fr2 
		on fr1.user1=? and fr2.user1=? 
		and fr1.user2=fr2.user2`
		result, err := db.Query(sqlStatement, friendResult.Friends[0], friendResult.Friends[1])
		check(err)
		for result.Next() {
			var email string
			err := result.Scan(&email)
			check(err)
			friendResult.Friends = append(friendResult.Friends, email)
			friendResult.Count++
		}
		friendResult.Success = true
	}
	json.NewEncoder(w).Encode(friendResult)
}
