package subscribe

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	db "src/db"
	rep "src/model/reponse"
	sub "src/model/subscribe"
)

func Sub(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var subscribe sub.Subscribe
	var response rep.Response
	response.Success = false

	json.NewDecoder(req.Body).Decode(&subscribe)
	db := db.DBConn()
	defer db.Close()
	insForm, err := db.Prepare("INSERT INTO subscribe(requestor,target,blocker) VALUES(?,?,?)")
	check(err)
	insForm.Exec(subscribe.Requestor, subscribe.Target, 1)
	response.Success = true

	json.NewEncoder(w).Encode(response)
}

func Block(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var subscribeRes sub.Subscribe
	var response rep.Response
	response.Success = false
	json.NewDecoder(req.Body).Decode(&subscribeRes)
	db := dbConn()
	update, err := db.Prepare("UPDATE subscribe SET blocker = 0 WHERE requestor = ? AND target = ? ")
	check(err)
	update.Exec(subscribeRes.Requestor, subscribeRes.Target)

	update, err = db.Prepare("UPDATE relationship SET blocker = 0 WHERE user1= ? AND user2 = ?")
	check(err)
	update.Exec(subscribeRes.Requestor, subscribeRes.Target)

	defer db.Close()
	response.Success = true
	json.NewEncoder(w).Encode(response)
}
func Recipients(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var recipientReq sub.RecipientRequest
	var recipientRes sub.RecipientResponse
	recipientRes.Success = false
	json.NewDecoder(req.Body).Decode(&recipientReq)
	emailList := getEmailList(recipientReq.Text)
	if emailList != nil && len(emailList) > 0 {
		db := dbConn()
		sqlStatment := `SELECT requestor FROM subscribe WHERE target =? `
		result, err := db.Query(sqlStatment, recipientReq.Sender)
		check(err)
		for result.Next() {
			var email string
			err := result.Scan(&email)
			check(err)
			emailList = append(emailList, email)
		}
		sqlStatment = `SELECT user2 FROM relationship WHERE user1 = ? and blocked=0 `
		result, err = db.Query(sqlStatment, recipientReq.Sender)
		check(err)
		for result.Next() {
			var email string
			err := result.Scan(&email)
			check(err)
			emailList = append(emailList, email)
		}
		if len(emailList) > 0 {
			recipientRes.Success = true
			recipientRes.Recipients = emailList
		}
	}
	json.NewEncoder(w).Encode(recipientRes)
}

func getEmailList(str string) []string {
	var emails []string
	words := strings.Fields(str)
	for _, word := range words {
		if checkEmailValidate(word) == true {
			emails = append(emails, word)
		}
	}
	return emails
}

func checkEmailValidate(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}
