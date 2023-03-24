package main

import "fmt"

func main() {
	lista := [3]int{1, 2, 3}
	soma := 0
	for x := 0; x < len(lista); x++ {
		soma += lista[x]
	}
	fmt.Println(soma)
}
