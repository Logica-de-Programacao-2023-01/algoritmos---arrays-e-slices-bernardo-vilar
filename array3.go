package main

import "fmt"

func main() {
	var count float32
	count = 1
	array := [4]float32{1.1, 2.2, 3.3, 4.4}
	for p := range array {
		count *= array[p]
	}
	fmt.Print(count)
}
