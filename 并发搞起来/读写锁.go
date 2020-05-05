package gogo

import (
	"fmt"
	"sync"
	"time"
)

var iwg *sync.WaitGroup
var rwLock *sync.RWMutex

var data int
func TestRWLock()  {
	rwLock = new(sync.RWMutex)
	iwg = new(sync.WaitGroup)

	iwg.Add(4)
	go readData()
	go readData()
	go readData()
	go writeData(3)
	iwg.Wait()
}
func readData()  {
	defer iwg.Done()
	rwLock.RLock()
	fmt.Println("读取中...")
	time.Sleep(time.Second)
	fmt.Println("读取数据",data)
	rwLock.RUnlock()
}
func writeData(i int)  {
	defer iwg.Done()
	rwLock.Lock()
	fmt.Println("写入中...")
	time.Sleep(time.Second*3)
	data = i
	fmt.Println("写入完毕...")
	rwLock.Unlock()
}
