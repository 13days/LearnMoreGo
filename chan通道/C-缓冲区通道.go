package gochan

import (
	"fmt"
	"strconv"
)

/**
可以理解为一个阻塞队列,不指定容量的,容量为1,故读写都阻塞....
非缓冲通道：make(chan T)
		一次发送，一次接收，都是阻塞的
缓冲通道：make(chan T , capacity)
	发送：缓冲区的数据满了，才会阻塞
	接收：缓冲区的数据空了，才会阻塞
 */

func TestBuf()  {
	ch1 := make(chan int, 5)
	fmt.Println(len(ch1),cap(ch1)) //0 5
	for i := 0; i < 5; i++ {
		ch1 <- i
		fmt.Println("大小:",len(ch1), "容量:",cap(ch1))
	}

	ch3 := make(chan string, 4)
	go sendDataStr(ch3)

	for{
		v, ok := <-ch3
		if !ok{
			fmt.Println("读完了。。",ok)
			break
		}
		fmt.Println("\t读取的数据是：",v)
	}
	fmt.Println("main。。over。。。。")
}
func sendDataStr(ch chan  string){
	for i:= 0;i<10;i++{
		ch <- "数据" + strconv.Itoa(i)
		fmt.Printf("子goroutine中写出第 %d 个数据\n",i)
	}
	close(ch)
}
