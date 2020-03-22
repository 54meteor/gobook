package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"os"
	"strings"
)
var lines = [5]int{2,4,6,8,10}
var titles = []string{"编号","部门","教研室","姓名"}
type  ora struct {
	//Title string
	Content string
}

type oraList struct{
	list []ora
}

func main(){
	path := "xls"
	fileList := readDir(path)
	var newList []oraList
	for key,file := range fileList{
		fmt.Println("正在读取文件" + file)
		list := readXlsx(path,file,key)
		newList = append(newList,list)
		fmt.Println("文件读取完成")
	}
	writingXlsx(newList)
	fmt.Println("文件写入成功")
}

func readDir(path string) []string {
	f,err := os.OpenFile(path,os.O_RDONLY,os.ModeDir)
	if err != nil{
		fmt.Println(err.Error())
	}
	defer f.Close()
	fileInfo , _ := f.Readdir(-1)

	fileList := []string {}
	for _,info := range fileInfo {
		if !info.IsDir(){
			fileList = append(fileList,info.Name())
		}
	}
	return fileList
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

func readXlsx(path,filename string,key int) oraList {

	var listOra = getPm(filename)
	sepa := string(os.PathSeparator)
	xlFile, err := xlsx.OpenFile(path + sepa + filename)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	sheets := xlFile.Sheets
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
				if key == 0 {
					titles = append(titles, title)
				}
				tmpOra.Content = content[j].String()
				listOra = append(listOra, tmpOra)
			}
		}

	}
	oraList := oraList{listOra}
	return oraList
}


func writingXlsx(oraList []oraList) {
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
	titleRow := sheet.AddRow()
	titleCell := titleRow.AddCell()
	titleCell.Value = "校区工作进度表"
	for key,list := range oraList{
		if key == 0{
			row = sheet.AddRow()
		}
		rowContent := sheet.AddRow()
		rowContent.SetHeightCM(10)
		style := xlsx.NewStyle()
		ali := *xlsx.DefaultAlignment()
		ali.WrapText = true
		ali.Vertical = "top"
		style.Alignment = ali
		for i := 0;i < len(titles);i++{
			if key == 0{
				cell = row.AddCell()
				cell.Value = titles[i]
			}
			cellContent := rowContent.AddCell()
			cellContent.Value = list.list[i].Content
			cellContent.SetStyle(style)
		}
	}
	err = file.Save("统计数据.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}