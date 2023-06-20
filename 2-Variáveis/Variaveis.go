package main

import "fmt"

func main() {
	var variavel1 string = "variaveis 1"
	variavel2 := "variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	variavel3 := "lalala"
	variavel4 := "sisisi"

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "varivael5", "varialvel6"
	fmt.Println(variavel5, variavel6)

	const contante string = "contante 1"
	fmt.Println(variavel5, variavel6)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
