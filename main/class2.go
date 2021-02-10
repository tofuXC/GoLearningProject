package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func arreyTest() {
	fmt.Println("第二节：数组、切片、map、函数内容")
	fmt.Println("数组初始化")
	fmt.Println("数组初始化正常定义")
	var testArray1 []int
	var testArray2 = [3]int{1, 2}
	var testAyyar3 = [...]int{1, 2, 3}
	testAyyar4 := [3]int{4, 5}
	testAyyar5 := [...]int{6, 7, 8, 9}
	fmt.Printf("数组初始化方法打印：A1：%v \t A2: %v \t A3: %v  \t A4: %v \t A5:%v\n",
		testArray1, testArray2, testAyyar3, testAyyar4, testAyyar5)
	fmt.Println("Array 可以通过指定索引的方法进行初始化，这个其他语言没有看到")
	testAyyar6 := [...]int{1: 1, 5: 5, 10: 10}
	fmt.Printf("索引格式化：A6%v\n", testAyyar6)
	fmt.Println("数组遍历，使用for循环，两种方法")
	for i := 0; i < len(testAyyar6); i++ {
		fmt.Printf("%v \t", testAyyar6[i])
	}
	fmt.Println()
	for k, v := range testAyyar6 {
		fmt.Printf("key:%v \t v:%v\t\t", k, v)
	}
	fmt.Println("多为数组的定义")
	var ttArray1 [3][4]int
	ttArray2 := [2][3]int{}
	fmt.Printf("多为数组：\nttA1:%v\nttA2:%v\n", ttArray1, ttArray2)
	for k, v := range ttArray2 {
		fmt.Println(k, v)
	}
	ttArray2[1][1] = 1

	var p2 *int
	p2 = &ttArray1[0][1]
	*p2 = 2
	//*(p2 + 1   ) = 3 指针操作数组后面找一下如何进行。
	fmt.Println(ttArray1)
}

func slideTest() {
	fmt.Println("切片学习")
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Printf("s1 == nil:%v \t s2 == nil:%v\t\n", s1 == nil, s2 == nil)
	s1 = []int{1, 2, 3}
	s2 = []string{"beijing", "hebei", "shandong"}
	fmt.Println(s1, s2)
	fmt.Printf("len(s1):%v \t len(s2):%v\t\n", len(s1), len(s2))
	fmt.Printf("cap(s1):%v \t cap(s2):%v\t\n", cap(s1), cap(s2))

	fmt.Println("make定义动态切片")
	m1 := make([]int, 2, 10)
	m1[1] = 5
	fmt.Printf("len(m1):%v\tcap(m1):%v\nm1:%v", len(m1), cap(m1), m1)

	m1 = append(m1, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("len(m1):%v\tcap(m1):%v\nm1:%v\n", len(m1), cap(m1), m1)

	a1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m2 := a1[:5]
	fmt.Printf("a1:%v\nm2:%v\n", a1, m2)
	a1[0] = 11
	fmt.Printf("a1:%v\nm2:%v\n", a1, m2)
	m2 = append(m2, 12, 13, 14)
	fmt.Printf("a1:%v\nm2:%v\n", a1, m2)

	fmt.Println("几个存储位置的测试")
	a2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m3 := make([]int, 2, 5)
	m3 = a2[:5]
	m4 := m3
	fmt.Println("源数据")
	fmt.Printf("a2:%v\nm3:%v\nm4:%v\n", a2, m3, m4)
	m3 = append(m3, 11, 12)
	fmt.Println("切片增加范围在数组范围内")
	fmt.Printf("a2:%v\nm3:%v\nm4:%v\n", a2, m3, m4)
	a2[0] = 10
	fmt.Println("修改数组内容后的影响")
	fmt.Printf("a2:%v\nm3:%v\nm4:%v\n", a2, m3, m4)
	m3 = append(m3, 13, 14, 15, 16)
	fmt.Println("切片新增跳出了数组的范围")
	fmt.Printf("a2:%v\nm3:%v\nm4:%v\n", a2, m3, m4)
	a2[0] = 17
	fmt.Println("切片跳出后修改数组的内容")
	fmt.Printf("a2:%v\nm3:%v\nm4:%v\n", a2, m3, m4)

	fmt.Println("切片的改变")
	a3 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m5 := a3[:5]
	m6 := m5
	fmt.Println("初始化数据")
	fmt.Printf("a2:%v\nm5:%v\nm6:%v\n", a3, m5, m6)
	m5 = append(m5, 20, 21, 22, 23, 24, 25)
	fmt.Println("一个切片爆掉")
	fmt.Printf("a2:%v\nm5:%v\nm6:%v\n", a3, m5, m6)
	m5[0] = 11
	fmt.Println("越界数组修改")
	fmt.Printf("a2:%v\nm5:%v\nm6:%v\n", a3, m5, m6)
	m7 := m5
	fmt.Println("原始切片m5:", m5)
	fmt.Println("赋值切片m7:", m7)
	m5 = append(m5, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	m5[0] = 100
	m7[1] = 101
	fmt.Println("m5重新分配m5:", m5)
	fmt.Println("赋值切片m7:", m7)

	fmt.Println("切片拷贝")
	m8 := make([]int, 8)
	copy(m8, m7)
	fmt.Println(m8)
	m8 = append(m8[:3], m8[4:]...)
	m8 = append(m8, 3, 4, 5, 6, 7, 8)

}
func mapTest() {
	fmt.Println("Map学习代码")
	fmt.Println("1、map定义。map[k]v")
	m1 := make(map[string]int, 2)
	m1["cyh"] = 18
	m1["sry"] = 18
	fmt.Println(m1)
	fmt.Println("2、判断某个值是否存在：")
	k := "cyh1"
	v, ok := m1[k]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println(k, "不存在")
	}

	fmt.Println("3、map的遍历, for遍历，for k , v = range m1")
	for k, v = range m1 {
		fmt.Println(k, v)
	}

	fmt.Println("4、删除键值对：delete(m1 , key),key可以存在也可以不存在")
	delete(m1, "cyh1")
	fmt.Println(m1)
	fmt.Println("5、按照顺序遍历map：将key放入切片，切片排序后根据key获取")

	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreM = make(map[string]int, 200)

	var kList = make([]string, 0, 200)
	for i := 0; i < 100; i++ {
		k := fmt.Sprintf("stu%02d", i) //stu开头的字符串
		v := rand.Intn(100)            // 身材0 -- 99 的随机整数
		scoreM[k] = v
		kList = append(kList, k)
	}
	for kT, vT := range scoreM {
		fmt.Println(kT, vT)

	}
	sort.Strings(kList)

	fmt.Println("按照排序后的顺序输出")
	for _, key := range kList { //切片遍历中，index,key 方式获取切片中元素.
		fmt.Println(key, scoreM[key])
	}

}
func sumTest(x int, y int) (int, int) {
	return x + y, x - y
}
func sumList(x int, y ...int) int {
	sum := 0
	for _, v := range y {
		sum += v
	}
	return sum + x

}

func sumRe() (sum int) {
	sum = 100
	return
}
func makeSuffixFunc(suffix string) func(string) string {
	return func(s string) string {
		if strings.HasSuffix(s, suffix) {
			return suffix
		}
		return s + suffix
	}
}
func funcLearning() {
	fmt.Println("函数定义")
	m, d := sumTest(10, 15)
	fmt.Printf("函数测试，和：%d \t 差：%d\t", m, d)
	fmt.Println("可以接收可变参数：用...在最后表示，注意类型。参数会变成一个list")
	fmt.Println(sumList(5, 5, 5, 6, 67))
	fmt.Println("返回值可以定义名字后直接使用，精简代码，return哪里也不需要写东西了")
	fmt.Println(sumRe())
	fmt.Println("函数也可以是变量和参数以及返回值，这个我理解很难受，后期深入和需要的时候再使用")
	fmt.Println("下面是匿名函数作为变量的示例，我的理解中，函数理论上参数和返回值确定后，占用空间就确定了，所以可以被赋值")
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))
	fmt.Println("匿名函数多用于实现闭包：这个概念我不是太理解")
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("testTxt"))
	fmt.Println(txtFunc("testTxt"))

}
func main3() {
	/*	fmt.Println("学习结束，下一课开始")
		a := make([]int , 2 ,4 )
		b := make(map[string]int , 3 )
		fmt.Println(a)
		fmt.Println(b)*/

}
