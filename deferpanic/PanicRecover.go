package main

import "fmt"

func main() {
	defer func(){
		str:=recover()
		fmt.Println("error string", str)
	}()
	panic("Panicking")
}
