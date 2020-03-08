package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"strings"
)
var lines = [5]int{2,4,6,8,10}
var titles = []string{"编号","部门","教研室","姓名"}
type  ora struct {
	//Title string
	Content string
}

func main(){
	excelFileName := "1-教务处-教务-莹莹.xlsx"
	fmt.Println("正在读取文件" + excelFileName)
	oraList := readXlsx(excelFileName)
	fmt.Println("文件读取完成，开始写入文件")
	writingXlsx(oraList)
	fmt.Println("文件写入成功")
}

func getPm(filename string) []ora{
	name := strings.Split(filename,".")
	pms := strings.Split(name[0],"-")
	var listOra []ora
	for _,pm := range pms{
		listOra = append(listOra,ora{pm})
	}
	return listOra
}

func readXlsx(filename string) []ora {
	//var listOra []ora
	var listOra = getPm(filename)
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	sheets := xlFile.Sheets
	//for _, sheet := range xlFile.Sheets {
	sheet := sheets[0]
	rows := sheet.Rows
	weeksRow := rows[1]
	weeks := weeksRow.Cells
		for i := 0;i < len(lines);i++{
			title := rows[lines[i]].Cells
			content := rows[lines[i] + 1].Cells
			for j := 0;j < len(title);j++{
				tmpOra := ora{}
				if len(title[j].String()) != 0 {
					var build strings.Builder
					build.WriteString(title[j].String())
					build.WriteString(" ")
					build.WriteString(weeks[j].String())
					title := build.String()
					titles = append(titles, title)
					//tmpOra.Title = title
					tmpOra.Content = content[j].String()
					listOra = append(listOra, tmpOra)
				}
			}
		}
	return listOra
}


func writingXlsx(oraList []ora) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	sheet.SetColWidth(0,3,10)
	sheet.SetColWidth(4,40,30)
	titleRow := sheet.AddRow();
	titleCell := titleRow.AddCell()
	titleCell.Value = "校区工作进度表"
	row = sheet.AddRow()
	rowContent := sheet.AddRow()
	//row.SetHeightCM(0.5)
	rowContent.SetHeightCM(10)
	style := xlsx.NewStyle()
	ali := *xlsx.DefaultAlignment()
	ali.WrapText = true
	ali.Vertical = "top"
	style.Alignment = ali
	for i := 0;i < len(titles);i++{
		cell = row.AddCell()
		cell.Value = titles[i]
		cellContent := rowContent.AddCell()
		cellContent.Value = oraList[i].Content
		cellContent.SetStyle(style)
	}
	err = file.Save("2019-_-_-2019-_-_Lag延时数据.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}