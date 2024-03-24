package main

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	fmt.Println("程序开始执行")
	file, err := excelize.OpenFile("./data.xlsx")
	if err != nil {
		fmt.Println("读取文件出错 ----->>>>", err)
		return
	}

	rows, err := file.GetRows("Sheet1")
	if err != nil {
		fmt.Println("读取列表出错 ----->>>>", err)
		return
	}

	failCount := 0
	emptyCount := 0
	successCount := 0

	for _, row := range rows {
		for _, colCell := range row {
			if len(colCell) == 0 {
				fmt.Println("发现一个空数据，跳过创建")
				emptyCount++
				continue
			}

			err := os.MkdirAll("./created/"+colCell, os.ModePerm)
			if err != nil {
				fmt.Println("文件夹创建出错 ----->>>>", err)
				failCount++
				continue
			}
			successCount++
		}
	}

	fmt.Printf("表格中含空数据：%d 个\n", emptyCount)
	fmt.Printf("文件夹创建出错：%d 个\n", failCount)
	fmt.Printf("文件夹创建完成：%d 个\n", successCount)
}
