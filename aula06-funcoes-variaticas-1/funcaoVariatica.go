package main

import "fmt"

func media(numeros ...float64) float64 {

	total := 0.0

	for _, valor := range numeros {

		total += valor

	}

	return total / float64(len(numeros))
}

func main() {

	resultado := media(5.5, 7.0, 6, 54.3, 33.1, 79.79, 80, 55.2331)

	fmt.Println(resultado)

}
