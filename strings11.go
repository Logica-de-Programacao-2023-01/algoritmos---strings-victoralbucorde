package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Escreva um programa que receba uma string e remova todas as vogais. Informe ao usuário o resultado.
func main() {
	vogais := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U", "ã", "õ", "â", "ô", "û", "ê", "î", "á", "ó", "ú", "é", "í", "Ã", "Õ", "Á", "É", "Í", "Ó", "Ú", "Â", "Ê", "Î", "Ô", "Û"}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase")
	scanner.Scan()
	frase := scanner.Text()
	for i := 0; i < len(vogais); i++ {
		frase = strings.ReplaceAll(frase, string(vogais[i]), "")
	}
	fmt.Println("Nova frase sem vogais", frase)
}
