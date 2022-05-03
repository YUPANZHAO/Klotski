package model

import (
	"bufio"
	"fmt"
	"os"
)

type KlotskiData struct {
	Width       int     //宽度
	Height      int     //高度
	Data        [][]int //滑块编号二维数组
	BlockNumber int     //滑块个数
	Target      int     //目标块
}

type KlotskiResult struct {
	Width   int       //宽度
	Height  int       //高度
	Target  int       //目标块
	DataLen int       //数据长度
	Data    [][][]int //滑块编号二维数组
}

func (data *KlotskiData) WriterToFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, data.Width, data.Height)
	for i := 0; i < data.Height; i++ {
		for j := 0; j < data.Width; j++ {
			fmt.Fprint(writer, data.Data[i][j], " ")
		}
		fmt.Fprintln(writer)
	}
	fmt.Fprintln(writer, data.BlockNumber, data.Target)
	writer.Flush()
	return nil
}

func (result *KlotskiResult) ReadFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	fmt.Fscanf(reader, "%d %d %d %d", &result.Width, &result.Height, &result.Target, &result.DataLen)
	fmt.Fscanln(reader)
	result.Data = make([][][]int, result.DataLen)
	for k := 0; k < result.DataLen; k++ {
		result.Data[k] = make([][]int, result.Height)
		for i := 0; i < result.Height; i++ {
			result.Data[k][i] = make([]int, result.Width)
			for j := 0; j < result.Width; j++ {
				fmt.Fscanf(reader, "%d", &result.Data[k][i][j])
			}
			fmt.Fscanln(reader)
		}
	}
	return nil
}
