package utils

import (
	"fmt"
	"time"
)

func PrintTime()  {
	fmt.Println(time.Now())
}

func init()  {
	fmt.Println("utils初始化1...")
}

func init()  {
	fmt.Println("utils初始化2...")
}

func init()  {
	fmt.Println("utils初始化3...")
}