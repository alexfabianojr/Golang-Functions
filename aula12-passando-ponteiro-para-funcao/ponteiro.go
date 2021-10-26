package main

import "fmt"

func inc1(n int) {
	n++
}

//revisao: um ponteiro eh representado por *
func inc2(n *int) {

	//revisao: * eh usado para desreferenciar, ou seja, ter acesso ao valor no qual o ponteiro aponta

	*n++
}

func main() {

	n := 1

	inc1(n) //n incrementa

	fmt.Println(n)

	inc2(&n) //revisao: & eh usado para passar a referencia em memoria

	fmt.Println(n)
}
