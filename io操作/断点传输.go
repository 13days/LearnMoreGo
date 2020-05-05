package ioOps

import (
	"fmt"
	"io"
	"os"
)

/*
Seek(offset int64, whence int) (int64, error)，设置指针光标的位置
第一个参数：偏移量
第二个参数：如何设置
0：seekstart，表示相对于文件开始
1：seekcurrent，表示相对于当前位置的偏移量
2：seekend，表示相对于末尾


// Seek whence values.
const (
	SeekStart   = 0 // seek relative to the origin of the file
	SeekCurrent = 1 // seek relative to the current offset
	SeekEnd     = 2 // seek relative to the end
)
*/

// 应用场景:如断电了进行后续传输
func SeekPrint(fileName string, offset int64)  {
	// 通过设置传输偏移量进行传输
	file, err := os.Open(fileName)
	if err != nil{
		fmt.Println("打开文件出错")
		return
	}
	defer file.Close()

	b := []byte{0}

	// 设置断点
	file.Seek(offset,io.SeekStart)
	for{
		n, err := file.Read(b)
		if n==0 || err==io.EOF{
			break
		}else if err != nil{
			fmt.Println("出错啦...")
			return
		}
		fmt.Print(string(b))
	}
}
