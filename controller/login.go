package controller

import (
	"KlotskiWeb/model"
	"log"
	"net/http"
)

func registerLoginRoutes() {
	http.HandleFunc("/login", handleLogin)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userEmail := query.Get("userEmail")
	password := query.Get("password")
	user, err := model.FindUserByEmail(userEmail)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(500) // Mysql发生错误
		return
	}
	w.WriteHeader(200)
	if user.Password == password {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAIL"))
	}
}
