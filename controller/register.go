package controller

import (
	"KlotskiWeb/common"
	"KlotskiWeb/config"
	"KlotskiWeb/model"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"gopkg.in/gomail.v2"
)

func registerRegisterRoutes() {
	http.HandleFunc("/register", handleRegitser)
	http.HandleFunc("/sendCode", handleSendCode)
}

// 用户注册接口
func handleRegitser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userEmail := query.Get("userEmail")
	password := query.Get("password")
	code := query.Get("code")
	//参数判空
	if userEmail == "" || password == "" || code == "" {
		model.WriteMessage(w, 400, "参数不全", nil)
		return
	}
	// 检查验证码
	ok := checkCode(userEmail, code)
	// 验证码错误或已过期
	if !ok {
		model.WriteMessage(w, 401, "验证码错误或已过期", nil)
		return
	}
	// 检查该邮箱是否已注册
	user, err := model.FindUserByEmail(userEmail)
	if err != nil {
		log.Print(err.Error())
		model.WriteMessage(w, 500, "mysql发生错误: "+err.Error(), nil)
		return
	}
	// 邮箱已经注册
	if user != (model.User{}) {
		model.WriteMessage(w, 401, "该邮箱已被注册", nil)
		return
	}
	// 注册账号
	user.Email = userEmail
	user.Password = password
	user.GameCounts = 5
	err = user.AddUser()
	if err != nil {
		log.Print(err.Error())
		model.WriteMessage(w, 500, "mysql发生错误: "+err.Error(), nil)
		return
	}
	// 注册成功
	model.WriteMessage(w, 200, "注册成功", nil)
}

// 获取验证码接口，并将验证码设置60秒生存期
func handleSendCode(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userEmail := query.Get("userEmail")
	//参数判空
	if userEmail == "" {
		model.WriteMessage(w, 400, "参数不全", nil)
		return
	}
	// 生成并发送验证邮件
	code := generateCode()
	err := sendVerifyEmail(userEmail, code)
	if err != nil {
		log.Print(err.Error())
		model.WriteMessage(w, 500, "验证邮件发送失败: "+err.Error(), nil)
		return
	}
	// 写入redis并设置60秒生存期
	err = setCode(userEmail, code)
	if err != nil {
		log.Print(err.Error())
		model.WriteMessage(w, 500, "redis错误: "+err.Error(), nil)
		return
	}
	// 验证邮件发送成功
	model.WriteMessage(w, 200, "验证邮件发送成功", nil)
}

// 向toEmail的邮箱发送包含验证码的邮件
func sendVerifyEmail(toEmail string, code string) (err error) {
	// 获取服务端邮箱配置信息
	email := config.NewEmailConfig()

	// 邮件内容配置
	msg := gomail.NewMessage()
	msg.SetHeader("From", email.UserName)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", email.Title)
	msg.SetBody("text/html", "您在 <b>华容道解密网站</b> 的注册验证码为 <b>"+
		code+"</b>。此验证码在 <b>1</b> 分钟内有效，请尽快完成注册。<br>"+
		"如果您没有在 <b>华容道解密网站</b> 上注册，请忽略本邮件。")
	// 发送邮件
	dialer := gomail.NewDialer(email.Host, email.Port, email.UserName, email.Password)
	err = dialer.DialAndSend(msg)
	return
}

// 生成6位数字的验证码
func generateCode() string {
	num, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	str := fmt.Sprintf("%06d", num)
	return str
}

// 写入redis并设置60秒生存期
func setCode(key string, value string) (err error) {
	rdb := common.RedisDB
	err = rdb.Set("verifyCode_"+key, value, 60*time.Second).Err()
	return
}

// 检查验证码是否正确且未过期
func checkCode(key string, value string) (result bool) {
	rdb := common.RedisDB
	code := rdb.Get("verifyCode_" + key).Val()
	result = (code != "" && code == value)
	return
}
