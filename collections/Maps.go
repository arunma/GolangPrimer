package main

import "fmt"

func main() {
	//x:=map[string]int -- Wont work

	x:=make(map[string]int)
	x["arun"]=10
	x["jason"]=100

	fmt.Println(x)
	fmt.Println(x["arun"])

	delete(x, "arun")

	fmt.Println(x)

	//lookup result - super nice
	name, ok := x["arun"]

	fmt.Println("Lookup for arun", name, ok)
}
