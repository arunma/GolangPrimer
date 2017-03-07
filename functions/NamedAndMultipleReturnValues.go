package main

import "fmt"

func main() {
	fmt.Println("f2 returning a value named r ", f2())
	x, y :=f3()
	fmt.Println("f3 returning a tuple ", x, y)
}
func f2()(r int) {
	r=100
	return
}

func f3()(int, int){
	return 10, 20
}