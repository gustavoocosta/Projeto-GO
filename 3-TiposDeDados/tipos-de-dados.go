package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 1000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	// BYTE = INT8
	var numero4 byte = 123
	fmt.Println(numero4)


	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// FIM NÚMEROS REAIS

	
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	//FIM STRINGS

	texto := 5
	fmt.Println(texto)
	
	var booleanol bool
	fmt.Println(booleanol)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}