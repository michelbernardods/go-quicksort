package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, n := range entrada {
		numero, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s não é um número válido\n", n)
			os.Exit(1)
		}
		numeros[i] = numero
	}

	fmt.Println(quicksort(numeros))
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	numsCopy := make([]int, len(numbers))
	copy(numsCopy, numbers)

	indicePivo := len(numsCopy) / 2
	pivo := numsCopy[indicePivo]
	numsCopy = append(numsCopy[:indicePivo], numsCopy[indicePivo+1:]...)

	numsMenores, numsMaiores := particionar(numsCopy, pivo)

	return append(append(quicksort(numsMenores), pivo), quicksort(numsMaiores)...)

}

func particionar(numbers []int, pivo int) (numsMenores []int, numsMaiores []int) {
	for _, n := range numbers {
		if n <= pivo {
			numsMenores = append(numsMenores, n)
		} else {
			numsMaiores = append(numsMaiores, n)
		}
	}

	return numsMenores, numsMaiores
}
