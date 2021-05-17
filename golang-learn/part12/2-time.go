package main

import (
	"fmt"
	"time"
)

func timeDemo()  {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func timestampDemp()  {
	now := time.Now()
	timestamp1 := now.Unix()      //10为时间戳
	timestamp2 := now.UnixNano()   // 13位时间戳
	fmt.Printf("current timestamp: %v\n", timestamp1)
	fmt.Printf("current timestamp: %v\n", timestamp2)
}



// timestamp -> time
func timestampToTime(timestamp int64)  {
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)

	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func timeCompute()  {
	// time add
	now := time.Now()
	later := now.Add(time.Hour)
	fmt.Println(later)

	// sub 求两个时间之间的差值：

	// Equal  判断两个时间是否相同

	// Before  如果t代表的时间点在u之前，返回真；否则返回假。

	// After  如果t代表的时间点在u之后，返回真；否则返回假。

}

func tickDemo()  {
	// 定时器
	ticker := time.Tick(time.Second)
	for i := range ticker{
		fmt.Println(i)//每秒都会执行的任务
	}
}

func timeFormatDemo()  {
	// 时间格式化
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))

	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

	fmt.Println(now.Format("2006/01/02 15:04:05"))

	fmt.Println(now.Format("15:04:05 2006/01/02"))

	fmt.Println(now.Format("2006/01/02"))
}

func stringTimeToTime()  {
	now := time.Now()
	fmt.Println(now)

	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil{
		fmt.Println(err )
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))


}


func main() {
	//timeDemo()
	//timestampDemp()
	//timestamp := time.Now().Unix()
	//timestampToTime(timestamp)
	//timeCompute()
	//tickDemo()
	//timeFormatDemo()
	stringTimeToTime()
}
