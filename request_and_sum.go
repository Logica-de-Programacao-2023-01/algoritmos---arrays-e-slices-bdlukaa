package main

import "fmt"

// Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos.
// Em seguida, imprima o Slice e a soma dos valores armazenados.

func main() {

	var amount int
	fmt.Println("Qual a quantidade?")
	fmt.Scan(&amount)

	var slice []int
	for i := 0; i < amount; i++ {
		slice = append(slice, 0)
	}

	for i := 0; i < amount; i++ {
		fmt.Print("Número na posição ", i, " é ")
		fmt.Scan(&slice[i])

	}

	fmt.Println("A lista é ", slice)

	var sum int = 0
	for _, e := range slice {
		sum += e
	}
	fmt.Println("A soma dos valores é", sum)

}
