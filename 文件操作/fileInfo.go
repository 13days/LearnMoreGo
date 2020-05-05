package fileOps

import (
	"fmt"
	"os"
)

/*
	FileInfo：文件信息
		interface
			Name()，文件名
			Size()，文件大小，字节为单位
			IsDir()，是否是目录
			ModTime()，修改时间
			Mode()，权限

*/

func GetFileInfo(fileName string) {
	// fileName = "C:\\Users\\WuChuPeng\\data_for_wuchupeng\\git_project_for_go\\src\\byteDancer.com\\wcp-step-2\\byteDancer.com\\wcp-step-2\\testPaacket\\aa.txt"
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
