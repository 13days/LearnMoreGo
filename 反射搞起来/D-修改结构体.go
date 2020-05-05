package goreflect

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
	School string
}

func TestModifyStruct()  {
	s1 := Student{"Lhh",23,"GDUF"}
	//通过反射，更改对象的数值：前提也是数据可以被更改
	fmt.Printf("%T\n",s1) //main.Student
	p1 := &s1
	fmt.Printf("%T\n",p1)  //*main.Student
	fmt.Println(s1.Name)
	fmt.Println((*p1).Name,p1.Name)

	// 改变数值
	value := reflect.ValueOf(&s1)

	if value.Kind() == reflect.Ptr{
		newValue := value.Elem()
		fmt.Println("是否可以修改:",newValue.CanSet())

		f1 :=newValue.FieldByName("Name")
		f1.SetString("刘慧慧")
		f3:=newValue.FieldByName("School")
		f3.SetString("幼儿园")
		fmt.Println(s1)
	}

}
