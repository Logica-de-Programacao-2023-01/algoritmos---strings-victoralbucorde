package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Solicite ao usuário uma string e informe se ela é uma sequência numérica decrescente (exemplo: "987" ou "54321").
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase")
	scanner.Scan()
	frase := scanner.Text()
	decrescente := true
	conversor, err := strconv.ParseFloat(frase, 64)
	if err != nil {
		fmt.Println("Não existe números")
	} else {
		for i := 0; i < len(frase)-1; i++ {
			if frase[i] < frase[i+1] {
				decrescente = false
			}
		}
		if decrescente {
			fmt.Println("A string é uma sequência decrescente.")
			fmt.Println(conversor)
		} else {
			fmt.Println("A string não é uma sequência decrescente.")
			fmt.Println(conversor)
		}
	}
}
