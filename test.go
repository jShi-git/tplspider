package main

import (
	"github.com/lealife/leacrawler"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strings"
)

var a string

func getTemplate(url string, customPath string) {
	customPath = strings.Replace(customPath, "\\", "\\\\", -1)
	customPath = strings.Replace(customPath, "/", "\\\\", -1)

	leacrawler.Fetch(url, customPath)
}

func main() {
	//var outTE *walk.TextEdit
	var le, le2 *walk.LineEdit
	var lb1, lb2 *walk.Label
	copytext, _ := walk.Clipboard().Text()
	MainWindow{
		Title:   "模板下载器",
		MinSize: Size{400, 100},
		Layout:  VBox{},
		Children: []Widget{
			Label{
				AssignTo: &lb1,
				Text:     "网址",
			},
			LineEdit{
				AssignTo: &le,
				Text:     copytext,
			},
			Label{
				AssignTo: &lb2,
				Text:     "保存目录",
			},
			LineEdit{
				AssignTo: &le2,
				Text:     `D:\APMServ5.2.6\www\htdocs\down\tpls\`,
			},
			//TextEdit{
			//	AssignTo: &outTE,
			//},
			PushButton{
				Text: "开始抓取",
				OnClicked: func() {
					//outTE.SetText(strings.ToUpper(inTE.Text()))
					getTemplate(le.Text(), le2.Text())
				},
			},
		},
	}.Run()
}
