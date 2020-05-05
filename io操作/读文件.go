package ioOps

import (
	"fmt"
	"io"
	"os"
)

/*
	读取数据：
		Reader接口：
			Read(p []byte)(n int, error)
*/
func TestRead(fileName string)  {
	//step1：打开文件
	file,err := os.Open(fileName)
	if err != nil{
		fmt.Println("打开文件失败")
		return
	}
	//step3：关闭文件
	defer file.Close()
	//step2：读取数据
	bs := make([]byte, 4, 4)
	n := -1
	for{
		n,err = file.Read(bs)
		if n == 0 || err == io.EOF{
			fmt.Println("读取到了文件的末尾，结束读取操作。。")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
