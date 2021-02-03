package main

import (
	"fmt"
)

//声明变量 直接声明方式
var name string
var age int
var isOk bool

//Go语言推荐使用驼峰命名
//批量声明
var (
	name1 string //""
	age1  string // 0
	isOk1 bool   // false
)

//声明＋初始化

var name2 string = "cyh"

//类型推到声明
var age2 = 18

//短声明，不可以全局声明
//函数中声明age3 := 18

//匿名变量 _ 用于接收函数返回值但是又不使用

//常量：定义后不可修改
const cyh = "chengyuehao"
const (
	name4 = "cyh"
	name5
	name6
	//name5 和name6 默认等于name4
)

//iota 使用
const (
	a1 = iota
	a2
	_ //占位，加一
	a3
)

//iota插队
const (
	b1 = iota
	b2 = 100
	b3 = iota
	b4
)

//iota, 多个声明zai一行
const (
	c1, d1 = iota, iota
	c2, d2 // 1 , 1 ,根据上面的c1，d1递增
)

//进制转换
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
)

func main1() {
	name = "程跃豪"
	age = 18
	isOk = true
	fmt.Println(name, age, isOk)
	//占位符相关可以查看：https://studygolang.com/articles/2644
	fmt.Printf("name :%s , age: %d , isOk: %t\n", name, age, isOk)
	//	短声明

	age2 := 18_00_00
	fmt.Println(age2)

	fmt.Printf("%v \t %v \t %v\n", name4, name5, name6)
	fmt.Printf("%v \t %v \t %v\n", a1, a2, a3)
	fmt.Printf("%v \t %v \t %v\n", b1, b2, b3)
	fmt.Printf("%v \t %v \t %v\t %v\n", c1, c2, d1, d2)
	fmt.Println(KB, MB)
	fmt.Println(name2)
	byteS1 := []byte(name)

	fmt.Println(name2, byteS1)
	byteS1[1] = 'i'
	//name2 = byteS1
	fmt.Println(name2, byteS1)
	a := 3
	fmt.Println("*******************")
	switch {
	case a <= 1:
		fmt.Println(1)
	case a <= 3:
		fmt.Println(3)
		//fallthrough
	case a <= 5:
		fmt.Println(5)
	}

	fmt.Println("************")
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(j)
			if j > 5 {
				break
			}
		}
		fmt.Println(i)
	}

}
