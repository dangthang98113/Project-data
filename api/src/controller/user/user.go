package user

import (
	db "/src/db"
	userReq "/src/model/user"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetEmails(w http.ResponseWriter, req *http.Request) {

	db := db.DBConn()
	defer db.Close()

	result, err := db.Query("SELECT email, firstname, lastname FROM users")
	check(err)
	var res []userReq.Emails
	for result.Next() {
		var emails userReq.Emails
		var firstName, lastName sql.NullString
		err := result.Scan(&emails.email, &firstName, &lastName)
		check(err)
		if firstName.Valid {
			emails.firstname = firstName.String
		}
		if lastName.Valid {
			emails.lastname = lastName.String
		}
		res = append(res, emails)

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(res)
	json.NewEncoder(w).Encode(res)

}
