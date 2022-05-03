package model

import (
	"KlotskiWeb/common"
)

type User struct {
	ID         int    //用户编号
	Email      string //用户邮箱
	Password   string //用户密码
	GameCounts int    //用户剩余解密次数
}

// 根据邮箱查找用户信息
func FindUserByEmail(email string) (user User, err error) {
	sql := "SELECT id, email, password, gameCounts FROM user WHERE email = ?"
	rows, err := common.MysqlDB.Query(sql, email)
	if err != nil {
		return
	}
	if rows.Next() {
		user = User{}
		err = rows.Scan(&user.ID, &user.Email, &user.Password, &user.GameCounts)
	}
	return
}

// 插入数据
func (user *User) AddUser() (err error) {
	sql := "INSERT INTO user (email, password, gameCounts) VALUES (?,?,?)"
	_, err = common.MysqlDB.Exec(sql, user.Email, user.Password, user.GameCounts)
	return
}

func GetGameCountsByUserId(user_id any) (int, error) {
	cnt := 0
	sql := "SELECT gameCounts FROM user WHERE id = ?"
	err := common.MysqlDB.QueryRow(sql, user_id).Scan(&cnt)
	return cnt, err
}

func SubGameCountsByUserId(user_id any) (err error) {
	sql := "UPDATE user SET gameCounts = gameCounts - 1 WHERE id = ?"
	_, err = common.MysqlDB.Exec(sql, user_id)
	return
}
