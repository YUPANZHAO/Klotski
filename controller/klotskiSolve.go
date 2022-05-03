package controller

import (
	"KlotskiWeb/model"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func registerKlotskiSolveRoutes() {
	http.HandleFunc("/klotskiSolve", handleKlotskiSolve)
}

func handleKlotskiSolve(w http.ResponseWriter, r *http.Request) {
	//用户id及随机数
	user_id := r.URL.Query().Get("user_id")
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % 0x100000
	//查看用户解密次数是否为0
	cnt, err := model.GetGameCountsByUserId(user_id)
	if err != nil || cnt <= 0 {
		model.WriteMessage(w, 401, "用户次数为零", nil)
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
	//将数据写入in.txt文本
	infilePath := fmt.Sprintf("__dirname/.././algorithm/in_%s_%05x.txt", user_id, randNum)
	err = data.WriterToFile(infilePath)
	if err != nil {
		model.WriteMessage(w, 500, "服务器异常: "+err.Error(), nil)
		return
	}
	defer os.Remove(infilePath)
	//执行算法程序
	cmd := exec.Command("__dirname/.././algorithm/klotski", user_id, fmt.Sprintf("%05x", randNum))
	err = cmd.Run()
	if err != nil {
		model.WriteMessage(w, 500, "服务端异常: "+err.Error(), nil)
		return
	}
	//获取文件数据
	outfilePath := fmt.Sprintf("__dirname/.././algorithm/out_%s_%05x.txt", user_id, randNum)
	result := model.KlotskiResult{}
	err = result.ReadFromFile(outfilePath)
	if err != nil {
		model.WriteMessage(w, 500, "服务器异常: "+err.Error(), nil)
		return
	}
	defer os.Remove(outfilePath)
	//处理成功
	model.SubGameCountsByUserId(user_id)
	model.WriteMessage(w, 200, "解密完成", result)
}
