package main

import (
	"fmt"
	"os"
)

func main() {
	defer second()
	first()

	openAndCloseFile("Pointers.go")

}
func first() {
	fmt.Println("first")
}
func second() {
	fmt.Println("second")
}

func openAndCloseFile(fileName string){
	f, e:=os.Open(fileName)
	fmt.Println(e) //Error
	defer f.Close()

}