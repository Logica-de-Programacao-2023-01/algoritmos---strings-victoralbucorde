package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Solicite ao usuário uma string e informe se ela contém pelo menos um número.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase")
	scanner.Scan()
	frase := scanner.Text()
	if strings.ContainsAny(frase, "1234567890") {
		fmt.Println("Contém pelo menos um número")
	} else {
		fmt.Println("Contém nenhum número")
	}
}
