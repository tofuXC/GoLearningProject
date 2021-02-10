package main

import "fmt"

func ptr1(p, n int) {
	p = n
}

func ptr2(p *int, n int) {
	*p = n
}

func main() {
	fmt.Println("第三课：指针、结构体")
	n1 := 10
	ptr1(n1, 100)
	fmt.Printf("n1 ptr1 后的值：%v\n", n1)
	ptr2(&n1, 100)
	fmt.Printf("n1 ptr2 后的值：%v\n", n1)

}
