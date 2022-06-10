package model

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type KlotskiData struct {
	Width       int     //宽度
	Height      int     //高度
	Data        [][]int //滑块编号二维数组
	BlockNumber int     //滑块个数
	Target      int     //目标块
}

type BlockMoveInfo struct {
	BlockId		int		//滑块编号		
	MoveDire	int		//移动方向
} 

type KlotskiResult struct {
	Width   int       			//宽度
	Height  int       			//高度
	Target  int       			//目标块
	DataLen int       			//数据长度
	Data    []BlockMoveInfo     //解密过程
}

func (data *KlotskiData) Obj2Str() string {
	str := ""
	str += fmt.Sprintf("%d %d", data.Width, data.Height)
	for i := 0; i < data.Height; i++ {
		for j := 0; j < data.Width; j++ {
			str += " "
			str += fmt.Sprint(data.Data[i][j])
		}
	}
	str += fmt.Sprintf(" %d %d", data.BlockNumber, data.Target)
	return str
}

func (data *KlotskiData) Solve() (result KlotskiResult, err error) {
	//连接算法解析接口
	conn, err := net.Dial("tcp", "192.168.71.2:4331")
	if err != nil {
		log.Println("连接：" + err.Error())
		return
	}
	defer conn.Close()
	//将对象转换为字符串格式
	str := data.Obj2Str()
	//发送数据
	_, err = conn.Write([]byte(str))
	if err != nil {
		log.Println("发送：" + err.Error())
		return
	}
	// fmt.Println("发送数据完毕")
	//接收数据
	var buf [1024]byte
	str = ""
	reader := bufio.NewReader(conn)
	for {
		len, erro := reader.Read(buf[:])
		if erro != nil {
			log.Println("接收：" + err.Error())
			err = erro
			return
		}
		// fmt.Println("接收数据长度：", len)
		temp := string(buf[:len])
		// fmt.Println("数据: " + temp)
		str += temp
		if strings.HasSuffix(temp, "#") {
			break
		}
	}
	// fmt.Println("接收数据完毕")
	// fmt.Println("result: " + str)
	if str == "-1 #" {
		return
	}
	//将字符串转换为对象
	result.Str2Obj(str)
	return
}

func (result *KlotskiResult) Str2Obj(str string) {
	nums := strings.Split(str, " ")
	flag := 0
	idx := 0
	for i, value := range nums {
		// fmt.Println("i = ", i, " value = ", value)
		if value == "#" {
			break
		}
		num, _ := strconv.Atoi(value)
		if i == 0 {
			result.Width = num
		} else if i == 1 {
			result.Height = num
		} else if i == 2 {
			result.Target = num
		} else if i == 3 {
			result.DataLen = num
			result.Data = make([]BlockMoveInfo, num)
		} else {
			if flag == 0 {
				result.Data[idx].BlockId = value
			}else {
				result.Data[idx].MoveDire = value
				idx++
			}
			flag ^= 1
		}
	}
}
