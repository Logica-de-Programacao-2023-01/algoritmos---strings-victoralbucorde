package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Escreva um programa que receba uma string e verifique se ela é um palíndromo. Informe ao usuário se é ou não.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase")
	scanner.Scan()
	frase := scanner.Text()
	var palindromo string

	for i := len(frase) - 1; i >= 0; i-- {
		letra := string(frase[i])
		palindromo += letra
	}
	palindromo = strings.ReplaceAll(palindromo, " ", "")
	frase = strings.ReplaceAll(frase, " ", "")
	if strings.ToUpper(palindromo) == strings.ToUpper(frase) {
		fmt.Println("É um palíndromo")
	} else {
		fmt.Println("Não é um palíndromo")
	}
}
