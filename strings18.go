package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Solicite ao usuário uma string e informe se ela contém somente números (0 a 9).
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase")
	scanner.Scan()
	frase := scanner.Text()
	i, err := strconv.ParseInt(frase, 10, 64)
	if err != nil {
		fmt.Println("A frase não contém apenas números")
	} else {
		fmt.Println("A frase contém apenas números", i)
	}

}
