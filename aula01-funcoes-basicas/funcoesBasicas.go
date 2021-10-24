package main

import "fmt"

func f1() {
	fmt.Println("f1")
}

func f2(p1 string, p2 string) {
	fmt.Println(p1, p2)
}

func f3() string {
	return "Deu certo"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "retorno1", "retorno2"
}

func main() {

	f1()

	f2("algo", "outro algo")

	fmt.Println(f3())

	fmt.Println(f4("Aqui", "Ali"))

	retorno1, retorno2 := f5()

	fmt.Println(retorno1, retorno2)

	_, retorno2Test := f5()

	fmt.Println(retorno2Test)
}
