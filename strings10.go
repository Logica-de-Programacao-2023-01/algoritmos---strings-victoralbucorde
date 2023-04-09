package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Escreva um programa que receba duas strings e verifique se elas são anagramas. Informe ao usuário se são ou não.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase")
	scanner.Scan()
	frase := scanner.Text()
	fmt.Println("Digite outra frase")
	scanner.Scan()
	frase2 := scanner.Text()
	anagrama := false
	if strings.ContainsAny(frase, frase2) {

		anagrama = true

	}
	if anagrama {
		fmt.Println("São anagramas")
	} else {
		fmt.Println("Não são anagramas")
	}
}
