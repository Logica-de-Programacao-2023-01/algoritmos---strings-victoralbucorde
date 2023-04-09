package main

import (
	"bufio"
	"fmt"
	"os"
)

// Escreva um programa que receba uma string e remova todos os espaços em branco. Informe ao usuário o resultado.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase")
	scanner.Scan()
	frase := scanner.Text()
	fmt.Println("Digite outra frase")
	scanner.Scan()
	frase2 := scanner.Text()

	if frase == frase2 {
		fmt.Println("As frases são iguais")
	} else {
		fmt.Println("As frases não são iguais")
	}
}
