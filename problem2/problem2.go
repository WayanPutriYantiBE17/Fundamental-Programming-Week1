package main

import (
	"fmt"
	"sort"
)

func MeanMedian(arrayInput []float64) (float64, float64) {
	// your code here
	sort.Float64s(arrayInput)

	var median float64
	var sum float64
	
	mid := len(arrayInput)/2
	if len(arrayInput)%2 ==0{
		median = (arrayInput[mid-1] + arrayInput[mid])/2
	}else{
		median = arrayInput[mid]
	}

	
	for _,nilai := range arrayInput{
		sum +=nilai
	}
	mean :=sum/float64(len(arrayInput))

	return mean,median
}

func main() {
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4}))          // 2.5 2.5
	fmt.Println(MeanMedian([]float64{1, 2, 3, 4, 5}))       // 3 3
	fmt.Println(MeanMedian([]float64{7, 8, 9, 13, 15}))     // 10.4 9
	fmt.Println(MeanMedian([]float64{10, 20, 30, 40, 50}))  // 30 30
	fmt.Println(MeanMedian([]float64{15, 20, 30, 60, 120})) // 49 30
}
