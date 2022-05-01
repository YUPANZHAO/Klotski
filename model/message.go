package model

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Code int
	Msg  string
	Data any
}

func NewMessage() Message {
	return Message{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}

func WriteMessage(w http.ResponseWriter, code int, msg string, data any) {
	w.WriteHeader(code)
	temp := NewMessage()
	temp.Code = code
	temp.Msg = msg
	temp.Data = data
	res, _ := json.Marshal(temp)
	w.Write(res)
}
