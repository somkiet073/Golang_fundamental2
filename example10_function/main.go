package main

import (
	"fmt"
)

func main() {
	// var number = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	helloFunc()
	summation(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fac := factorial(5)
	fmt.Println("factorial =", fac)
	fmt.Println(func() string {
		return "Hello Go"
	}())
}

func helloFunc() {
	fmt.Println("Hello frist func")
}

func summation(args ...int) {
	var total int

	for _, n := range args {
		total += n
	}
	fmt.Println(total)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func factory() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
