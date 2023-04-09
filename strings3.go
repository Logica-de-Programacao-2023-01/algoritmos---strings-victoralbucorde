package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Escreva um programa que receba uma string e substitua todas as ocorrências desse caractere
// na string por outro caractere especificado pelo usuário. Informe ao usuário o resultado.
func main() {
	var x string
	var y string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Escreva uma frase: ")
	scanner.Scan()
	frase := scanner.Text()
	fmt.Println("Escreva um caractere para ser substituido por outro respectivamente")
	fmt.Scan(&x)
	fmt.Scan(&y)
	frase = strings.Replace(frase, x, y, -1)
	fmt.Println(frase)
}
