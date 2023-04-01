package main

import "fmt"

func main() {
	matriz := [3][2]int{}
	for linha := 0; linha < 3; linha++ {
		for coluna := 0; coluna < 2; coluna++ {
			fmt.Println("Que numero inteiro deseja inserir na matriz?")
			fmt.Scan(&matriz[linha][coluna])
		}
	}
}
