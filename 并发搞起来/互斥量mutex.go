package gogo

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	4个goroutine，模拟4个售票口，


	在使用互斥锁的时候，对资源操作完，一定要解锁。否则会出现程序异常，死锁等问题。
	defer语句
*/

// 全局变量,表示票
var tickets = 10

// 互斥锁
var mutex sync.Mutex

// 同步等待组
var _wg sync.WaitGroup

func TestMutex()  {
	_wg.Add(4)
	go sellTicket("销售台-1")
	go sellTicket("销售台-2")
	go sellTicket("销售台-3")
	go sellTicket("销售台-4")
	_wg.Wait()
}

func sellTicket(name string)  {
	rand.Seed(time.Now().UnixNano())
	defer _wg.Done()

	for{
		// 上锁
		mutex.Lock()
		if tickets > 0{
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name,"售出票:",tickets)
			tickets--
		}else{
			mutex.Unlock()
			fmt.Println("售完...")
			break
		}
		mutex.Unlock()
	}
}



