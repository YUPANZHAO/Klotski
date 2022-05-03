package middleware

import (
	"KlotskiWeb/common"
	"KlotskiWeb/model"
	"fmt"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Next http.Handler
}

func (am *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if am.Next == nil {
		am.Next = http.DefaultServeMux
	}

	urlPath := r.URL.Path
	if urlPath == "/" ||
		urlPath == "/login" ||
		urlPath == "/register" ||
		urlPath == "/sendCode" ||
		urlPath == "/callback" ||
		urlPath == "/notify" {
		am.Next.ServeHTTP(w, r)
		return
	}

	uri := r.RequestURI
	if strings.LastIndex(uri, ".css") > -1 ||
		strings.LastIndex(uri, ".png") > -1 ||
		strings.LastIndex(uri, ".jpg") > -1 ||
		strings.LastIndex(uri, ".svg") > -1 ||
		strings.LastIndex(uri, ".html") > -1 ||
		strings.LastIndex(uri, ".js") > -1 ||
		strings.LastIndex(uri, ".css") > -1 {
		am.Next.ServeHTTP(w, r)
		return
	}

	token := r.Header.Get("authorization")
	if token == "" {
		model.WriteMessage(w, 401, "未获取到用户token", nil)
		return
	}

	res := common.RedisDB.Get(token[strings.Index(token, " ")+1:])
	userId, err := res.Int()
	if err != nil {
		model.WriteMessage(w, 401, "token已过期, 请重新登录", nil)
		return
	}

	query := r.URL.Query()
	query.Set("user_id", fmt.Sprint(userId))
	r.URL.RawQuery = query.Encode()

	am.Next.ServeHTTP(w, r)
}
