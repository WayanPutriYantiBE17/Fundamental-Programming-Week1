package main

import (
	"fmt"
)

func MunculSekali(angka string) []int {
	// your code here
	//deklarasi untuk menghitung jumlah kemunculan setiap karakter
	count :=make(map[rune]int)

	//menghitung jumlah kemunculan setiap karakter pada string
	for _, huruf := range angka{
		count[huruf]++
	}

	//membuat slice untuk menyimpan karakter yang muncul sekali
	hasil :=[]int{}
	for _, huruf := range angka{
		if count[huruf] == 1{
			hasil = append(hasil, int(huruf-'0'))
		}
	}
	return hasil

}

func main() {
	fmt.Println(MunculSekali("1234123"))    // [4]
	fmt.Println(MunculSekali("76523752"))   // [6, 3]
	fmt.Println(MunculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(MunculSekali("1122334455")) // []
	fmt.Println(MunculSekali("0872504"))    // [8 7 2 5 4]
}
