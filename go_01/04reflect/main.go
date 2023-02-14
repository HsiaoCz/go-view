package main

import (
	"fmt"
	"reflect"
)

// 反射 所谓的反射，是指在程序运行期间可以访问、检测或者修改它本身状态或行为的能力
// 在go语言中，go提供了一种机制，可以在程序运行期间更新变量和检查他们的的值、调用它们的方法，而在编译期间不知道这些变量的具体类型
// 这种机制就叫反射，在go语言中，每个变量都
// 这里有个面试题，reflect 包如何获取字段tag?为什么json包不能导出私有变量的tag
// tag信息可以通过反射(reflect包)内的方法获取

type J struct {
	a string // 小写无tag
	B string `json:"B"` // 小写+tag
	C string //大写无tag
	D string `json:"DD" otherTag:"good"` //大写tag
}

func printTag(stu any) {
	t := reflect.TypeOf(stu).Elem()
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("结构体内第%v个字段 %v 对应的json tag 是%v,还有othertag =%v\n", i+1, t.Field(i).Name, t.Field(i).Tag.Get("json"), t.Field(i).Tag.Get("otherTag"))
	}
}

func main() {
	j := J{
		a: "1",
		B: "2",
		C: "3",
		D: "4",
	}
	printTag(&j)
	var s float32 = 3.14
	reflectValue(s)
	var m int = 12
	reflectSetValue2(&m)
}

// 关于反射的一些方法
// reflect.TypeOf().Elem()获取指针指向的值对应的结构体的内容
// NumField()可以获得该结构体含有几个字段
// 遍历结构体的字段，通过t.Field(i).Tag.Get("json")可以获取到tag为json的字段
// 如果结构体的字段有多个tag，比如otherTag，同样可以通过t.Field(i).Tag.Get("otherTag")获得
// json包不能导出私有变量的tag是因为json包里认为私有变量为不可导出的Unexported，json包会跳过取值
// 在go语言的反射里，任何接口值都是由一个具体的类型和具体类型的值两部分组成
// 也就是reflect.TypeOf()和reflect.ValueOf()
// 使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。
// typeOf()获取类型信息，不过type分类name和kind
// kind针对底层类型
// relfect.ValueOf()获取值信息

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 在反射修改值Elem()获取指针对应的值
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
