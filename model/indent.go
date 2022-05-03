package model

import "KlotskiWeb/common"

type Indent struct {
	IndentId      string //订单编号
	Time          string //订单创建时间
	UserId        int    //用户编号
	AddGameCounts int    //增加解密次数
	Status        string //订单状态
}

func (indent *Indent) AddIndent() (err error) {
	sql := "INSERT INTO indent (indent_id, time, user_id, addGameCounts, status) VALUES (?,?,?,?,?)"
	_, err = common.MysqlDB.Exec(sql, indent.IndentId, indent.Time, indent.UserId, indent.AddGameCounts, indent.Status)
	return
}

func FinishIndentByIndentId(indent_id string) (err error) {
	sql := "UPDATE user SET gameCounts = gameCounts + (SELECT addGameCounts FROM indent WHERE indent_id = ? AND status = '待支付')"
	_, err = common.MysqlDB.Exec(sql, indent_id)
	if err != nil {
		return
	}
	sql = "UPDATE indent SET status = '已完成' WHERE indent_id = ?"
	_, err = common.MysqlDB.Exec(sql, indent_id)
	return
}
