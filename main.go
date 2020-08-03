package main

import (
	"github.com/go-programming-tour-book/tour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}

//func main()  {
//	location, _ := time.LoadLocation("Asia/Shanghai")
//	inputTime := "2022-02-02 11:11:11"
//	layout := "2006-01-02 15:04:05"
//
//	//t, _ := time.Parse(layout, inputTime)
//	//dataTime := time.Unix(t.Unix(), 0).In(location).Format(layout)
//	//log.Printf("input time: %s, output time: %s", inputTime, dataTime)
//
//	/*
//	  2020/08/03 12:21:26 input time: 2022-02-02 11:11:11, output time: 2022-02-02 19:11:11
//	  与Parse有关系，因为Parse方法会尝试在入参的参数中分析并读取时区信息
//	  如果入参的参数没有指定时区信息，那么会默认使用UTC时间
//	  在这种情况下，我们采用ParseInLocation方法指定时区
//	*/
//
//	t, _ := time.ParseInLocation(layout, inputTime, location)
//	dataTime := time.Unix(t.Unix(), 0).In(location).Format(layout)
//	log.Printf("input time: %s, output time: %s", inputTime, dataTime)
//	// 2020/08/03 12:28:09 input time: 2022-02-02 11:11:11, output time: 2022-02-02 11:11:11
//}