package main

import "fmt"

func main() {
	floatArr:=[]float64{10,21,31,41,51,61,71,81}
	fmt.Printf("average of %v is %v", floatArr, calculateAverage(floatArr))
}
func calculateAverage(floatArray []float64) float64{
	total:=0.0

	for _, v := range floatArray{
		fmt.Println(v)
		total+=v
	}

	fmt.Println(total)

	return total/float64(len(floatArray))
}
