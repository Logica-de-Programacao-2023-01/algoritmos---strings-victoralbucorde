package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Solicite ao usuário duas strings e informe se a segunda é uma substring da primeira.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase")
	scanner.Scan()
	frase := scanner.Text()
	fmt.Println("Digite uma segunda frase")
	scanner.Scan()
	frase2 := scanner.Text()

	if strings.Contains(frase, frase2) {
		fmt.Println("A segunda é uma substring da primeira")
	} else {
		fmt.Println("A segunda não é uma substring da primeira")
	}
}
