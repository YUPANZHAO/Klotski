package model

import (
	"KlotskiWeb/common"
)

type User struct {
	ID         int
	Email      string
	Password   string
	GameCounts int
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
		if err != nil {
			return
		}
	}
	return
}

// 插入数据
func (user *User) AddUser() (err error) {
	sql := "INSERT INTO user (email, password, gameCounts) VALUES (?,?,?)"
	_, err = common.MysqlDB.Exec(sql, user.Email, user.Password, user.GameCounts)
	return
}
