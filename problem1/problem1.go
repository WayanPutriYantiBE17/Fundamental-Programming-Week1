package main

import "fmt"

func DrawXYZ(n int) string{
	// your code here
	var Out string
	
	for i := 0; i < n; i++ {
		for j := 1; j <= n; j++ {
			if ((i*n)+j)%3 == 0 {
				Out += "X "
			}else if ((i*n)+j)%2 == 0 {
				Out +="Z "
			}else if ((i*n)+j)%2 == 1 {
				Out += "Y "
			}
			
		}
		Out +="\n"	
	}
	return Out
}

func main() {
	hasil1 :=DrawXYZ(3)
	hasil2 :=DrawXYZ(5)
	fmt.Println(hasil1)
	fmt.Println(hasil2)

}
