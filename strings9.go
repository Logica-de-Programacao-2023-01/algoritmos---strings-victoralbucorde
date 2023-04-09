package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Solicite ao usuário uma string e substitua todas as ocorrências de uma letra por outra informada pelo usuário.
func main() {
	var x string
	var y string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase")
	scanner.Scan()
	frase := scanner.Text()
	fmt.Println("Escreva duas letras para uma ser substituída por outra respectivamente")
	fmt.Scan(&x)
	fmt.Scan(&y)
	novafrase := strings.ReplaceAll(frase, x, y)
	fmt.Println(novafrase)
}
