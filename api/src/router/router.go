package router

import (
	friend "/src/controller/friend"
	sub "/src/controller/subscribe"
	get "/src/controller/user"

	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/api/emails/showall", get.GetEmails()).Methods("GET")
	r.HandleFunc("/api/emails/addfriend", friend.AddFriendEmail()).Methods("POST")
	r.HandleFunc("/api/emails/showlistfriend", friend.ShowListFriendEmail()).Methods("GET")
	r.HandleFunc("/api/emails/showmutualfriend", friend.ShowMutualFriendEmail()).Methods("POST")
	r.HandleFunc("/api/emails/subscribe", sub.Sub()).Methods("POST")
	r.HandleFunc("/api/emails/block", sub.Block()).Methods("PUT")
	r.HandleFunc("/api/emails/recipients", Recipients).Methods("POST")
}
