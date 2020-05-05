package main

import (
	goreflect "byteDancer.com/wcp-step-2/LearnMoreGo/反射搞起来"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func GetFileInfo(fileName string) {
	//fileName = "C:\\Users\\WuChuPeng\\data_for_wuchupeng\\git_project_for_go\\src\\byteDancer.com\\wcp-step-2\\byteDancer.com\\wcp-step-2\\testPaacket\\aa.txt"
	// 打开文件
	//fileName = "aa.txt"
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)
	//文件名
	fmt.Println(fileInfo.Name())
	//文件大小
	fmt.Println(fileInfo.Size())
	//是否是目录
	fmt.Println(fileInfo.IsDir()) //IsDirectory
	//修改时间
	fmt.Println(fileInfo.ModTime())
	//权限
	fmt.Println(fileInfo.Mode()) //-rw-r--r--
}
func main() {
	//fmt.Println("main....")
	//utils.PrintTime()
	//a := mathUtils.Add(4,5)
	//fmt.Println(a)

	//sql包.InitDB()

	//time包使用.TimeTest()

	// 我的天...以当前调用的路径作为路径
	//fileOps.GetFileInfo("aa.txt")
	//fileOps.GetFileInfo("./文件操作/bb.txt")
	//GetFileInfo("aa.txt")
	//fileOps.TestFileOps()

	// 读文件
	// ioOps.TestRead("aa.txt")
	// 写文件
	//ioOps.WriteBytes("cc.txt",[]byte{'A','B','C','D'})
	//ioOps.WriteString("cc.txt","sedfsfsd")
	// 文件复制
	//ioOps.CopyFile("cc.txt","hhh.txt")
	// 断点操作
	//ioOps.SeekPrint("cc.txt", 5)

	// 缓存IO
	//fmt.Print(ioOps.ScanLine())
	//ioOps.BufferReadAndPrint("hhh.txt")

	// 答应目录
	//ioOps.ListFiles("C:\\Users\\WuChuPeng\\data_for_wuchupeng\\git_project_for_go\\src\\byteDancer.com\\wcp-step-2",0)

	// Go并发示例
	//gogo.TestGo()
	//gogo.TestRace()

	// runtime
	//gogo.GetOsPath()
	//gogo.QuitCPU()
	//gogo.TestGoExit()
	//gogo.TestSellTickets()
	//gogo.TestWG()
	//gogo.TestMutex()
	//gogo.TestRWLock()

	// 通道
	// 互斥的通道
	//gochan.TestChan()
	// 关闭遍历
	//gochan.TestClose()
	// range遍历
	//gochan.TestRangClose()
	// 缓冲区通道
	//gochan.TestBuf()
	// 单向通道
	//gochan.TestSingleChan()
	// 时间通道
	//gochan.TestTimer()
	// select挑选可执行通道
	//gochan.TestSelect()
	// 小demo生产消费
	//gochan.TestPS()

	// 反射
	// 获取反射的Type和Value
	//goreflect.TestTypeValue()
	// 接口和反射类型的转换
	//goreflect.TestReflectInterface()
	// 获取结构体信息
	//goreflect.TestGetStructMessage()
	// 修改值,修改底层都是不安全指针来修改的,不安全指针可以转换为任何非结构体指针
	//goreflect.TestModifyByPtr()
	// 修改结构体
	//goreflect.TestModifyStruct()
	// 方法调用
	// goreflect.TestReflectMethod()
	// 函数也可以进行反射
	goreflect.ReflectMethod()
}



