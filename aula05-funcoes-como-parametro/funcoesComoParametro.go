package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int) int {
	return funcao(5, 8)
}

func main() {

	fmt.Println(exec(multiplicacao))

}
