package main

import "fmt"

func argsPointerT1(a *int) {
	*a = 20
}

func argsPointerT2(a int) {
	a = 30
}
func getSumAndSub(n1 int, n2 int) (int, int) {
	return n1 + n2, n1 - n2

}

func main() {
	var a1 = 10
	fmt.Println("a1", a1)
	var a2 *int = &a1
	argsPointerT1(a2)
	fmt.Println("a1", a1)
	argsPointerT2(a1)
	fmt.Println("a1", a1)
	fmt.Println(getSumAndSub(3, 5))
}
