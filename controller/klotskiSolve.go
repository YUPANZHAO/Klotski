package controller

import (
	"KlotskiWeb/model"
	"encoding/json"
	"net/http"
)

func registerKlotskiSolveRoutes() {
	http.HandleFunc("/klotskiSolve", handleKlotskiSolve)
}

func handleKlotskiSolve(w http.ResponseWriter, r *http.Request) {
	//用户id
	user_id := r.URL.Query().Get("user_id")
	//查看用户解密次数是否为0
	cnt, err := model.GetGameCountsByUserId(user_id)
	if err != nil || cnt <= 0 {
		model.WriteMessage(w, 400, "用户次数为零", nil)
		return
	}
	//获取请求Body
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	//将数据解析成KlotskiData是实体
	data := model.KlotskiData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		model.WriteMessage(w, 400, "数据异常: "+err.Error(), nil)
		return
	}
	//执行算法
	result, err := data.Solve()
	if err != nil {
		model.WriteMessage(w, 500, "数据处理失败", nil)
		return
	}
	if result.Data == nil {
		model.WriteMessage(w, 400, "未找到解决方案", nil)
		return
	}
	//处理成功
	model.SubGameCountsByUserId(user_id)
	model.WriteMessage(w, 200, "解密完成", result)
}
