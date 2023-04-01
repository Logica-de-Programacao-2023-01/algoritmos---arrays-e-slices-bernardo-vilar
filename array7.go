package main

import "fmt"

func main() {
	var verify bool
	var num int
	slice := []int{1, 2, 3, 4, 5}
	fmt.Print("Qual o numero?")
	fmt.Scan(&num)
	for _, x := range slice {
		if num == x {
			verify = true
		}
	}
	if verify {
		fmt.Println("Numero ja esta na lista")
	} else {
		slice = append(slice, num)
		fmt.Println("Sua nova lista e:", slice)
	}
}
