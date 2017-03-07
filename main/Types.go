package main

import "fmt"

func main ()  {
	fmt.Println("Hello world")
	var x string
	x = "Interesting variable def"
	fmt.Println(x)

	y:="Variable with starting value"

	fmt.Println("Y is ", y)

	const aconstant string ="Some constant"

	//aconstant=10

	fmt.Println(aconstant)
}