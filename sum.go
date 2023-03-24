package main

import "fmt"

func main() {

	var array = [3]int{1, 2, 3}

	sum := 0

	for _, e := range array {
		sum += e
	}

	fmt.Println(sum)

}
