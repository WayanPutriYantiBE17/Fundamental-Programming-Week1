package main

import "fmt"

func RemoveDuplicates(array []int) int {
	// your code here
	unik := make(map[int]bool)
	for _,item := range array {
		unik[item]=true
	}

	hasil :=make([]int,0,len(unik))
	for item :=range unik{
		hasil = append(hasil,item)
	}
	
	
	return len(hasil)
}

func main() {
	fmt.Println(RemoveDuplicates([]int{2, 3, 3, 3, 6, 9, 9})) //4
	fmt.Println(RemoveDuplicates([]int{2, 3, 4, 5, 6, 9, 9})) //6
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))         //2
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))         //2
	fmt.Println(RemoveDuplicates([]int{1, 2, 3, 11, 11}))     //4
}
