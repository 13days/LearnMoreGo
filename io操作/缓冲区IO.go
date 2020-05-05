package ioOps

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读取一行
func ScanLine() string{
	b := bufio.NewReader(os.Stdin)
	s,_ := b.ReadString('\n')
	return s
}

// 缓冲区读
func BufferReadAndPrint(fileName string)  {
	file,err := os.Open(fileName)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	//创建Reader对象
	b1 := bufio.NewReader(file)
	//1.Read()，高效读取
	//p := make([]byte,1024)
	//n1,err := b1.Read(p)
	//fmt.Println(n1)
	//fmt.Println(string(p[:n1]))

	//for{
	//	s1,err := b1.ReadString('\n')
	//	if err == io.EOF{
	//		fmt.Println("读取完毕。。")
	//		break
	//	}
	//	fmt.Println(s1)
	//}

	//4.ReadBytes()
	//data,err :=b1.ReadBytes('\n')
	//fmt.Println(err)
	//fmt.Println(string(data))

	for{
		s1,err := b1.ReadBytes('\n')
		fmt.Println(string(s1)) // 怎么肥事...多了那么多空行
		if err == io.EOF{
			fmt.Println("读取完毕。。")
			break
		}
	}
}

// 缓冲区写
