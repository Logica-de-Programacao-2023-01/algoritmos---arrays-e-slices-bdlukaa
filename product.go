package main

import "fmt"

// Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.

func main() {
	var floats = [5]float32{4.5, 6.5, 7.6, 8.7}

	var product float32 = 0

	for _, float := range floats {
		product *= float
	}

	fmt.Println("O produto é", product)
}
