package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo
}

func main() {

	valor1, valor2 := trocar(79, 45)

	fmt.Println(valor1)
	fmt.Println(valor2)

}
