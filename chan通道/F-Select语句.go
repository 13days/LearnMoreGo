package gochan

import (
	"fmt"
	"time"
)

/*
	分支语句：if，switch，select
	select语句类型于switch语句
		但是select语句会随机执行一个可运行的case
		如果没有case可以运行，要看是否有default，如果有就执行default，否则就进入阻塞，直到有case可以运行
*/
func TestSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()

	// 从所有就绪的通道随机挑选一个,若有default则执行
	select {
	case num := <-ch1:
		fmt.Println("从ch1中获取数据", num)
	case num, ok := <-ch2:
		if ok {
			fmt.Println("从ch2中获取数据", num)
		} else {
			fmt.Println("ch2通道已经关闭")
		}
	case num := <-time.After(2*time.Second):
		fmt.Println("从timer里获取数据",num)
	//default:
	//	fmt.Println("default被执行了...")
	}
}
