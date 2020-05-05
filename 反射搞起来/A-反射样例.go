package goreflect

import (
	"fmt"
	"reflect"
)

// 接口就是维护了Type和Value
func TestTypeValue()  {
	var x float64 = 3.14

	// 根据反射获取变量的类型和值
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))

	fmt.Println(`--------------`)
	// 根据反射值获取:类型,种类,值
	t := reflect.ValueOf(x)
	fmt.Println("Kind:",t.Kind())
	fmt.Println("Type:",t.Type())
	fmt.Println("Value:",t.Float()) // 知道是Float类型
}

func TestReflectInterface()  {
	var num float64 = 1.23
	// "接口类型变量"-->"反射类型对象"
	value := reflect.ValueOf(num)

	// "反射类型对象" --> "接口类型变量"
	convertValue:=value.Interface().(float64)
	fmt.Println(convertValue)

	/*
		反射类型对象-- >接口类型变量，理解为"强制转换"
		Golang对类型要求非常严格，类型一定要完全符合
		一个是*float64，一个float64，如果弄混，直接panic
	*/
	pointer := reflect.ValueOf(&num)
	convertPointer := pointer.Interface().(*float64)
	fmt.Println(convertPointer)
}
