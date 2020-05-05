package gochan

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
func TestChan()  {
	wg.Add(20)
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go in(c, i)
		go out(c)
	}
	wg.Wait()
}

func in(c chan int , a int)  {
	defer wg.Done()
	fmt.Println("往通道写数据中....:",a)
	time.Sleep(1*time.Second)
	c <- a // 阻塞的...
	fmt.Println("写完了....")
}
func out(c chan int)  {
	defer wg.Done()
	data := <- c // 阻塞的
	fmt.Println("从通道获取数据:",data)
}