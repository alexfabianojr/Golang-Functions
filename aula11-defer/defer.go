package main

import "fmt"

func obterValorAprovado(numero int) int {

	defer fmt.Println("Fim!")

	if numero > 5000 {
		fmt.Println("Valor alto")
		return 5000
	} else {
		fmt.Println("Valor baixo")
		return 5000
	}
}

func main() {

	obterValorAprovado(98234932)
}
