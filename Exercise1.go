package main

import "fmt"

func main() {
	input:=[]int {1, 2}
	for _,v := range input{
		half, even:=halfAndEven(v)
		fmt.Printf("Input %v Half %v Even %v \n", v, half, even)
	}


	//Find greatest
	fmt.Println("Greatest", greatest(1,3,5,2,3,9,6))

	makeOddFn:=makeOddGenerator()
	fmt.Println("Odd", makeOddFn())
	fmt.Println("Odd", makeOddFn())
	fmt.Println("Odd", makeOddFn())
}
func halfAndEven(num int) (int, bool) {
	return num/2, num%2==0

}

func greatest(nums ...int) int {
	var max int
	for _, each:= range nums{
		if max<each{
			max=each
		}
	}
	return max
}

func makeOddGenerator() func() uint {
	start:=uint(1)
	return func() (ret uint) {
		ret=start
		start+=2
		return
	}
}