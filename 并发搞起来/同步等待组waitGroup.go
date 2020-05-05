package gogo

import (
	"fmt"
	"sync"
)

/*
	WaitGroup：同步等待组
		Add()，设置等待组中要执行的子 goroutine的数量

		Wait()，让主goroutine出于等待


		Done()，让等待组中的counter计数器的值减1，同Add(-1)


*/
var wg sync.WaitGroup

func TestWG()  {
	wg.Add(2)
	go f1()
	go f2()
	wg.Wait()
}

func f1()  {
	defer wg.Done()
	for i:=1;i<10;i++{
		fmt.Println("fun1()函数中打印。。A ",i)
	}
}

func f2()  {
	defer wg.Done()
	for j:=1;j<10;j++{
		fmt.Println("\tfun2()函数打印。。",j)
	}
}
