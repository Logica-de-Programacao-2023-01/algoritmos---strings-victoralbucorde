package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Solicite ao usuário uma string e informe se ela é uma sequência numérica crescente (exemplo: "123" ou "4567").
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase")
	scanner.Scan()
	frase := scanner.Text()
	crescente := true
	conversor, err := strconv.ParseFloat(frase, 64)
	if err != nil {
		fmt.Println("Não existe números")
	} else {
		for i := 0; i < len(frase)-1; i++ {
			if frase[i+1] < frase[i] {
				crescente = false
			}
		}
		if crescente {
			fmt.Println("A string é uma sequência crescente.")
			fmt.Println(conversor)
		} else {
			fmt.Println("A string não é uma sequência crescente.")
			fmt.Println(conversor)
		}
	}
}
