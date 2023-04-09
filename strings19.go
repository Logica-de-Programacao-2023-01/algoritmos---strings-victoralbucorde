package main

import (
	"bufio"
	"fmt"
	"os"
)

// Solicite ao usuário uma string e imprima a mesma string invertida.
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
	fmt.Println("Frase invertida ", novafrase)
}
