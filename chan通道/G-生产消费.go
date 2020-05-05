package gochan

import (
	"fmt"
	"math/rand"
	"time"
)

func product(ch chan <- int, name string){
	rand.Seed(time.Now().UnixNano())
	for{
		time.Sleep(time.Millisecond*time.Duration(rand.Intn(5000)))
		n := rand.Intn(1000)
		ch <- n
		fmt.Println(name,"生产数据:",n)
	}
}

func consume(ch <- chan int, name string)  {
	rand.Seed(time.Now().UnixNano())
	for{
		time.Sleep(time.Millisecond*time.Duration(rand.Intn(2000)))
		n := <- ch
		fmt.Println(name,"消费数据:",n)
	}
}

func TestPS()  {
	ch := make(chan int, 20)
	go consume(ch,"消费者-A")
	go product(ch,"生产者-A")
	go consume(ch,"消费者-B")
	go product(ch,"生产者-B")
	go consume(ch,"消费者-C")
	go product(ch,"生产者-C")
	time.Sleep(1000*time.Second)
}

