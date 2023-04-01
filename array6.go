package main

import "fmt"

func main() {
	var verify bool
	var num int
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Print("Informe um numero para verificar se ta na lista:")
	fmt.Scan(&num)
	for _, x := range array {
		if num == x {
			verify = true
		}
	}
	if verify {
		fmt.Print("O numero esta na lista")
	} else {
		fmt.Println("O numero nao esta na lista")
	}
}
