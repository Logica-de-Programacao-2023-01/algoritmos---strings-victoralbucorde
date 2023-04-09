package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Escreva um programa que receba uma string e conte quantas palavras ela contém. Informe ao usuário o resultado.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase ou numeros")
	scanner.Scan()
	frase := scanner.Text()
	contador := strings.Split(frase, " ")
	fmt.Printf("A frase %s contém %d palavras", frase, len(contador))

}
