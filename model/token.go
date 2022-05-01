package model

import (
	"KlotskiWeb/common"
	"math/rand"
	"time"
)

type Token struct {
	Token string
}

// 创建token并把token和用户id写入redis，有效期24小时
func NewToken(id int) (token Token, err error) {
	// 生成不重复的token
	for {
		token.Token = randStr(32)
		if !IsTokenExist(token.Token) {
			break
		}
	}
	// 写入redis
	err = setToken(token.Token, id)
	return
}

// 生成长度为length的随机字符串
func randStr(length int) string {
	var str string
	str = ""
	rand.Seed(time.Now().UnixNano())
	for len := 0; len < length; len++ {
		t := rand.Intn(3)
		if t == 0 {
			str += string(rand.Intn(10) + '0')
		} else if t == 1 {
			str += string(rand.Intn(26) + 'a')
		} else {
			str += string(rand.Intn(26) + 'A')
		}
	}
	return str
}

// 判断token是否存在
func IsTokenExist(token string) (exist bool) {
	res := common.RedisDB.Get(token)
	_, err := res.Int()
	return err == nil
}

// 把token-id写入redis，过期时间24小时
func setToken(token string, id int) (err error) {
	err = common.RedisDB.Set(token, id, 24*time.Hour).Err()
	return err
}
