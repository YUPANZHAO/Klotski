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
	//参数判空
	if userEmail == "" || password == "" {
		model.WriteMessage(w, 400, "参数不全", nil)
		return
	}
	//登录逻辑
	user, err := model.FindUserByEmail(userEmail)
	if err != nil {
		log.Print(err.Error())
		model.WriteMessage(w, 500, "mysql错误: "+err.Error(), nil)
		return
	}
	if user.Email != "" && user.Password == password {
		token, err := model.NewToken(user.ID)
		if err != nil {
			log.Fatal(err.Error())
			model.WriteMessage(w, 500, "token生成失败: "+err.Error(), nil)
			return
		}
		model.WriteMessage(w, 200, "登录成功", token)
	} else {
		model.WriteMessage(w, 401, "登录失败", nil)
	}
}
