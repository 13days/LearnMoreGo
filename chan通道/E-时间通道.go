package gochan

import (
	"fmt"
	"time"
)

/*
	1. func NewTimer(d Duration) *Timer
			创建一个计时器，d时间以后触发


*/

func TestTimer()  {
	newTimer := time.NewTimer(2*time.Second)
	// 阻塞
	fmt.Println("等待中...")
	<- newTimer.C
	fmt.Println("等待完毕")
	//新建一个计时器
	timer := time.NewTimer(5*time.Second)
	//开始goroutine，来处理触发后的事件
	go func() {
		<- timer.C
		fmt.Println("Timer 结束了。。。开始。。。。")
	}()

	time.Sleep(3*time.Second)

	// 中断timer返回是否中断成功
	flag := timer.Stop()
	if flag{
		fmt.Println("Timer 停止了。。。")
	}
}

func TestAfter()  {
	/*
		 2. func After(d Duration) <-chan Time
				返回一个通道：chan，存储的是d时间间隔之后的当前时间

			相当于：return NewTimer(d).C
	*/
	ch :=  time.After(3 *time.Second)
	fmt.Printf("%T\n",ch) //<-chan time.Time
	fmt.Println(time.Now()) //2019-08-15 11:43:33.941039 +0800 CST m=+0.000537462


	time2 := <-ch
	fmt.Println(time2) //2019-08-15 11:43:36.945775 +0800 CST m=+3.005338577
}
