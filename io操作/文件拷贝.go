package ioOps

import (
	"byteDancer.com/wcp-step-2/LearnMoreGo/utils/errorUtils"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CopyFile(srcFile, destFile string)(int, error){
	// 打开源文件
	file1, err := os.Open(srcFile)
	errorUtils.HandleErr(err)
	defer file1.Close()

	// 打开目标文件
	file2, err := os.OpenFile(destFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	errorUtils.HandleErr(err)
	defer file2.Close()

	// 写过程
	n := -1
	count := 0
	bs := make([]byte, 1024, 1024)
	for{
		n, err = file1.Read(bs)
		if n==0 || err==io.EOF{
			fmt.Println("拷贝完毕...")
			break
		}else if err!=nil{
			fmt.Println("出错啦...")
			return count,err
		}
		count += n
		file2.Write(bs)
	}
	return count,nil
}

// 使用io包下的copy实现,还有copyN,copyBuffer
func CopyFile2(srcFile, destFile string) (int64,error){
	// 打开源文件
	file1, err := os.Open(srcFile)
	errorUtils.HandleErr(err)
	defer file1.Close()

	// 打开目标文件
	file2, err := os.OpenFile(destFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	errorUtils.HandleErr(err)
	defer file2.Close()
	return io.Copy(file2,file1)
}

// 使用ioutil工具包实现,不适合大文件
func CopyFile3(srcFile, destFile string) (int,error){
	bs,err := ioutil.ReadFile(srcFile)
	if err != nil{
		return 0,err
	}
	err = ioutil.WriteFile(destFile,bs,os.ModePerm)
	if err != nil{
		return 0,err
	}
	return len(bs),nil
}
