package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Escreva um programa que receba uma string
// e verifique se ela é um número válido em ponto flutuante.
// Informe ao usuário se é ou não.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase ou numeros")
	scanner.Scan()
	frase := scanner.Text()
	i, err := strconv.ParseFloat(frase, 64)

	if err != nil {
		fmt.Printf("A frase %s não possui nenhum valor associado\n", frase)
	} else {
		fmt.Printf("O valor é %f\n", i)
	}
}
