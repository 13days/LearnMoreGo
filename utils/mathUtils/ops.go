package mathUtils

import "fmt"

func Add(a, b int)int{
	return a+b
}

func Sub(a, b int)int{
	return a-b
}

func Mul(a, b int)int{
	return a*b
}

func init()  {
	fmt.Println("mathUtils的init1...")
}