package controller

import (
	"KlotskiWeb/model"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func registerPayRoutes() {
	http.HandleFunc("/pay", handlePay)
	http.HandleFunc("/callback", handleCallBack)
	http.HandleFunc("/notify", handleNotify)
}

//发起支付
func handlePay(w http.ResponseWriter, r *http.Request) {
	//获取token及用户id
	userId, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	//配置订单
	alipay := model.NewAlipayService()
	//在订单表中添加订单
	indent := model.Indent{}
	indent.IndentId = alipay.OutTradeNo
	indent.Time = time.Now().String()[0:18]
	indent.UserId = userId
	indent.AddGameCounts = 60
	indent.Status = "待支付"
	err := indent.AddIndent()
	if err != nil {
		log.Println(err.Error())
		model.WriteMessage(w, 500, "mysql错误: "+err.Error(), nil)
		return
	}
	//发起订单
	sHtml, err := alipay.DoPay()
	if err != nil {
		log.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(sHtml))
}

//同步回调
func handleCallBack(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//接收get参数
	params := r.URL.Query()

	aliPay := model.NewAlipayService()
	_, err := aliPay.VerifySign(params)
	if err != nil {
		log.Println("error: " + err.Error())
		model.WriteMessage(w, 500, "error: "+err.Error(), nil)
		return
	}
	indent_id := params.Get("out_trade_no")
	model.WriteMessage(w, 200, "支付成功！订单号: "+indent_id, nil)
	// 添加60次解密次数
	err = model.FinishIndentByIndentId(indent_id)
	if err != nil {
		log.Println("订单异常: indent_id = " + indent_id)
	}
}

//异步回调通知处理
func handleNotify(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//接收post参数
	r.ParseForm()
	formdata := r.PostForm

	aliPay := model.NewAlipayService()
	_, err := aliPay.VerifySign(formdata)
	if err != nil {
		log.Print("error: " + err.Error())
		model.WriteMessage(w, 500, "error: "+err.Error(), nil)
		return
	}
	// 添加60次解密次数
	params := r.URL.Query()
	indent_id := params.Get("out_trade_no")
	err = model.FinishIndentByIndentId(indent_id)
	if err != nil {
		log.Println("订单异常: indent_id = " + indent_id)
	}
	//程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）；
	fmt.Fprintf(w, "success")
}
