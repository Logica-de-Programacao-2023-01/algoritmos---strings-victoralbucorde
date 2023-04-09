package main

import (
	"bufio"
	"fmt"
	"os"
)

// Escreva um programa que receba uma string e inverta a ordem dos caracteres. Informe ao usuário o resultado.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase")
	scanner.Scan()
	frase := scanner.Text()
	var novafrase string

	for i := len(frase) - 1; i >= 0; i-- {
		armazenamento := string(frase[i])
		novafrase += armazenamento
	}
	fmt.Println(novafrase)
}
