package main

import "fmt"

func main() {
	num:=5
	fmt.Println(num)
	fmt.Println("Num reference ", &num)
	makeZero(&num)
	fmt.Println(num)
}

func makeZero(num *int){
	*num=0
}