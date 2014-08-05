package main

import (
	"flag"
	"fmt"
	"github.com/lealife/leacrawler"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"os"
	"os/exec"
	"strings"
)

var a string

func getTemplate(url string, customPath string) {
	customPath = strings.Replace(customPath, "\\", "\\\\", -1)
	customPath = strings.Replace(customPath, "/", "\\\\", -1)

	leacrawler.Fetch(url, customPath)
}

func main() {
	var le, le2 *walk.LineEdit
	var lb1, lb2 *walk.Label
	var ck1 *walk.CheckBox
	var mw *walk.MainWindow

	copytext, _ := walk.Clipboard().Text()

	updateIcon := func() {
		mainIcon, err := walk.NewIconFromResource("101")
		if err != nil {
			// do some work
		}
		mw.SetIcon(mainIcon)
	}

	if err := (MainWindow{
		AssignTo: &mw,
		Title:    "模板下载器",
		MinSize:  Size{400, 100},
		Layout:   VBox{},
		Children: []Widget{
			CheckBox{
				AssignTo: &ck1,
				Name:     "cksnap",
				Text:     "带截图抓取",
				Checked:  false,
			},
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
			PushButton{
				Text: "开始抓取",
				OnClicked: func() {
					fmt.Println(ck1.Checked())
					if ck1.Checked() {
						command := flag.String("cmd", "phantomjs", "Set the command.")
						patharr := strings.Split(le2.Text(), "\\\\")
						args := flag.String("args", "rasterize.js "+le.Text()+" "+patharr[len(patharr)-1]+".png", "Set the args. (separated by spaces)")
						var argArray []string
						if *args != "" {
							argArray = strings.Split(*args, " ")
						} else {
							argArray = make([]string, 0)
						}

						flag.Parse()
						cmd := exec.Command(*command, argArray...)

						buf, err := cmd.Output()
						if err != nil {
							fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, *command, *args)
							return
						}
						fmt.Fprintf(os.Stdout, "Result: %s", buf)

					}
					getTemplate(le.Text(), le2.Text())
				},
			},
		},
	}.Create()); err != nil {
		fmt.Println(err)
	}
	updateIcon()
	mw.Run()
}
