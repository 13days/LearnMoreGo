package gochan

import "fmt"

/*
		单向：定向
		chan <- T，只支持写
		<- chan T，只读


	定向通道：
		双向：-->函数：只读，只写
*/
//该函数，只能操作只写的通道
func fun1(ch chan <- int){
	//在函数内部，对于ch1通道，只能写数据，不能读取数据
	ch <- 100
	fmt.Println("fun1函数结束。。。")
}

//该函数，只能操作只读的通道
func fun2(ch <- chan int){
	data := <-ch
	fmt.Println("fun2函数，从ch中读取的数据是：",data)
}

func TestSingleChan(){
	ch1 := make(chan int) //双向，读，写
	go fun1(ch1)
	fun2(ch1)
	fmt.Println("main..over...")
}

