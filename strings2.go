package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Escreva um programa que receba uma string e remova todos os espaços em branco. Informe ao usuário o resultado.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()
	frase = strings.Replace(frase, " ", "", -1)
	fmt.Println(frase)
}
