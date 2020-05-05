package ioOps

import (
	"byteDancer.com/wcp-step-2/LearnMoreGo/utils/errorUtils"
	"fmt"
	"os"
)

func WriteBytes(fileName string, bs []byte)  {
	// 打开文件,支持创建和写操作和写追加,附加权限777
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	errorUtils.HandleErr(err)
	// 关闭文件
	defer file.Close()
	// 写数据
	n,err := file.Write(bs)
	errorUtils.HandleErr(err)
	fmt.Println("成功写入",n,"个字节")
}

func WriteString(fileName string, str string)  {
	// 打开文件,支持创建和写操作和写追加,附加权限777
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	errorUtils.HandleErr(err)
	// 关闭文件
	defer file.Close()
	// 写数据
	n,err := file.WriteString(str)
	errorUtils.HandleErr(err)
	fmt.Println("成功写入",n,"个字节")
}