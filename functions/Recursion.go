package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(fibonacci(7))
}
func factorial(num int) int {
	if (num==0){
		return 1
	}
	return num * factorial(num - 1 )

}

func fibonacci (num int) int {
	switch num{

	case 0:
		return 1
	case 1:
		return 1
	default:
		return fibonacci(num-1) + fibonacci(num-2)
	}

}