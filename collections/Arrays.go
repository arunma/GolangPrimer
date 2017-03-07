package main

import "fmt"

func main() {
	var x [5] int
	x[4]=100
	fmt.Print(x)

	//Init on create and loop
	floats :=[]float64{
		12,12,4,5,2,3,5,3535,5,3535,35.0,
	}

	var total float64
	for counter:=0; counter< len(floats); counter++ {
		fmt.Println(counter)
		total += floats[counter]
	}

	fmt.Println("fmt println ", total)
	println("println total , ", total)

	//Slice using make
	madeArray:= make([]float64, 5, 10)
	fmt.Println("Made array", madeArray)

	slicedArray:=floats[0:5]

	fmt.Println(slicedArray)


	//Append
	slice1:=[]int {1,2,3}
	slice2:=append (slice1, 4,5,6)

	fmt.Println(slice1)

	fmt.Println(slice2)

	//Copy - to a smaller array
	src:=[]int{1,2,3,4,5}
	dest:=make([]int,2)

	copy(dest, src)

	fmt.Println("Dest", dest)





}
