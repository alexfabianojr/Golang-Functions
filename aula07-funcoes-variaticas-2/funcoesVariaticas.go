package main

import "fmt"

func imprimirAprovados(aprovados ...string) {

	for i, aprovado := range aprovados {
		fmt.Println(i, aprovado)
	}

}

func main() {

	aprovados := []string{"Maria", "Pedro", "Jose", "Rafael"}

	imprimirAprovados(aprovados...) //funciona com slice mas n funciona com array

}
