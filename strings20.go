package main

import (
	"fmt"
	"unicode"
)

// Solicite ao usuário uma string e informe se ela é está em camelCase e quantas palavras possuí.
// CamelCase é quando a primeira letra de cada palavra é maiúscula, e as demais são minúsculas.
// Exemplo: "estaStringEstaEmCamelCase".
func main() {
	var frase string
	fmt.Println("Digite uma frase sem espaços")
	fmt.Scan(&frase)
	camel := false

	for _, c := range frase {
		if unicode.IsUpper(c) {
			camel = true
			break
		}
	}
	if camel {
		fmt.Println("Está em camelCase")
		maiusculas := 1
		for _, c := range frase {
			if unicode.IsUpper(c) {
				maiusculas++
			}
		}
		fmt.Printf("A frase contém %d palavras.\n", maiusculas)
	} else {
		fmt.Println("Não está em camelCase")
		maiusculas := 1
		fmt.Printf("A frase contém %d palavras.\n", maiusculas)
	}

}
