package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Solicite ao usuário uma string e imprima somente as suas letras únicas (que aparecem apenas uma vez).
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase")
	scanner.Scan()
	frase := scanner.Text()
	frase = strings.ReplaceAll(frase, " ", "")
	frase = strings.ToUpper(frase)

	frequencia := make(map[rune]int)
	for _, c := range frase {
		frequencia[c]++
	}

	fmt.Println("Caracteres que aparecem somente uma vez:")
	for c, quantidade := range frequencia {
		if quantidade == 1 {
			fmt.Printf("%c", c)
		}
	}
}
