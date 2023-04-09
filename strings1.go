package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Escreva um programa que receba uma string e converta todas as letras minúsculas para maiúsculas.
// Informe ao usuário o resultado.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()
	frase = strings.ToUpper(frase)
	fmt.Println(frase)
}
