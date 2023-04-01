package main

import "fmt"

func main() {
	var tamanho int
	var elemento int
	var contador int
	slice := []int{}
	fmt.Print("Qual o tamanho da slice?")
	fmt.Scan(&tamanho)
	for x := 0; x < tamanho; x++ {
		fmt.Print("Qual numero quer botar na lista?")
		fmt.Scan(&elemento)
		slice = append(slice, elemento)
		contador += slice[x]
	}
	fmt.Println(slice)
	fmt.Println(contador)
}
