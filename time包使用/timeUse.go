package time包使用

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	time包:
		day
		hour
		minute
		second
		millisecond
		nanosecond
		picosecond
 */

func TimeTest()  {
	// 1.获取当前时间
	t1 := time.Now()
	fmt.Printf("%T\n",t1)
	fmt.Println(t1)

	// 2.指定时间
	t2 := time.Date(2020,5,4,14,17,30,0,time.Local)
	fmt.Println(t2)

	// 3.time->string
	s1 := t1.Format("2006-1-2 15:4:5")
	fmt.Println(s1)

	// 4.string->time
	t3, _ := time.Parse("2006-1-2 15:4:5", "1999-2-3 12:20:3")
	fmt.Println(t3)

	// 5.获取时间内容
	y,m,d := t1.Date()
	fmt.Println(y,m,d)

	h,mm,s := t1.Clock()
	fmt.Println(h,mm,s)

	yd := t1.YearDay()
	fmt.Println("今年过了:",yd,"天")
	fmt.Println("年:",t1.Year())
	fmt.Println("月:",t1.Month())
	fmt.Println("日:",t1.Day())
	fmt.Println("时:",t1.Hour())
	fmt.Println("分:",t1.Minute())
	fmt.Println("秒:",t1.Second())
	fmt.Println("纳秒:",t1.Nanosecond())
	fmt.Println("星期",t1.Weekday())

	// 6.时间戳
	fmt.Println("毫秒:",t2.Unix())
	fmt.Println("纳秒:",t2.UnixNano())

	// 7.时间间隔:Duration
	t4 := t2.Add(time.Nanosecond)
	fmt.Println(t4)

	// 8.加减
	t5 := t1.AddDate(1,0,0)
	fmt.Println(t5)

	t6 := t5.Sub(t1) // 俩时间的差
	fmt.Println(t6)

	// 9.睡眠
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	fmt.Println("睡眠:",n,"秒")
	time.Sleep(time.Duration(n)*time.Second)
	fmt.Println("睡眠结束")
}