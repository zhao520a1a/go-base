package main

import (
	"encoding/base64"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// 源码:https://gitee.com/xurime/excelize
// 帮助文档：https://xuri.me/excelize/zh-hans/stream.html#Flush
// 注：
// 1.图片缩放建议使用一个开源库(github.com/nfnt/resize)，不要使用360库自带的缩放，会有一点问题
// 2.先对单元格做大小控制，再插入图片，否则有可能会导致一些图片被拉伸
func main() {
	file1Name := "utils/excel/test1.xlsx"
	newNormalExcelFile(file1Name)
	readNormalExcel(file1Name)

	file2Name := "utils/excel/test2.xlsx"
	newChartExcel(file2Name)

	file3Name := "utils/excel/test3.xlsx"
	newPicExcel(file3Name)
	readPicExcel(file3Name)

	file4Name := "utils/excel/test4.xlsx"
	WriteExcelFile(file4Name)
}

func newNormalExcelFile(fileName string) {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet2")
	// 设置单元格的值
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
}

func readNormalExcel(fileName string) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	//获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func newChartExcel(fileName string) {
	//创建柱状图
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	if err := f.AddChart("Sheet1", "E1", `{"type":"col3DClustered","series":[{"name":"Sheet1!$A$2","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$2:$D$2"},{"name":"Sheet1!$A$3","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$3:$D$3"},{"name":"Sheet1!$A$4","categories":"Sheet1!$B$1:$D$1","values":"Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Clustered Column Chart"}}`); err != nil {
		fmt.Println(err)
		return
	}
	// 根据指定路径保存文件
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
}

func newPicExcel(fileName string) {
	// 创建一个工作表
	f := excelize.NewFile()
	// 插入图片
	if err := f.AddPicture("Sheet1", "A2", "image.png", ""); err != nil {
		fmt.Println(err)
	}
	// 在工作表中插入图片，并设置图片的缩放比例
	if err := f.AddPicture("Sheet1", "D2", "image.jpg", `{"x_scale": 0.5, "y_scale": 0.5}`); err != nil {
		fmt.Println(err)
	}
	// 在工作表中插入图片，并设置图片的打印属性
	if err := f.AddPicture("Sheet1", "H2", "image.gif", `{"x_offset": 15, "y_offset": 10, "print_obj": true, "lock_aspect_ratio": false, "locked": false}`); err != nil {
		fmt.Println(err)
	}
	// 根据指定路径保存文件
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
}

func readPicExcel(fileName string) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取工作表中指定单元格的值
	//name, ff, err := f.GetPicture("Sheet1", "A2")
	//name, ff, err := f.GetPicture("Sheet1", "D2")
	_, ff, err := f.GetPicture("Sheet1", "H2")
	if err != nil {
		fmt.Println(err)
		return
	}

	bufstore := make([]byte, 5000000)                                              //数据缓存
	base64.StdEncoding.Encode(bufstore, ff)                                        // 文件转base64
	_ = ioutil.WriteFile(fmt.Sprintf("./%d.png", time.Now().UnixNano()), ff, 0666) //直接写入到文件就ok完活了。
}

// 流式写数据
func WriteExcelFile(fileName string) {
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	styleID, err := file.NewStyle(`{"font":{"color":"#777777"}}`)
	if err != nil {
		fmt.Println(err)
	}
	if err := streamWriter.SetRow("A1", []interface{}{excelize.Cell{StyleID: styleID, Value: "Data"}}); err != nil {
		fmt.Println(err)
	}
	for rowID := 2; rowID <= 10000; rowID++ {
		row := make([]interface{}, 5)
		for colID := 0; colID < 5; colID++ {
			row[colID] = rand.Intn(640000)
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		if err := streamWriter.SetRow(cell, row); err != nil {
			fmt.Println(err)
		}
	}
	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
	}
	if err := file.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}

}
