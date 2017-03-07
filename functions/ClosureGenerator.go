package main

import "fmt"

func main() {
	nextInt:=nextIntFn()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextEvenFn:=evenNumberGeneratorFn()
	fmt.Println(nextEvenFn())
	fmt.Println(nextEvenFn())
	fmt.Println(nextEvenFn())

}
func nextIntFn() func() int {
	current:=0
	return func() (ret int){
		ret=current
		current+=1
		return
	}

}

func evenNumberGeneratorFn() func() uint{
	current:=uint(0)
	return func() (ret uint) {
		ret=current
		current+=2
		return
	}
}