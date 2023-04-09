package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Solicite ao usuário uma string e substitua todas as vogais por "*" (asterisco).
func main() {
	vogais := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U", "ã", "õ", "â", "ô", "û", "ê", "î", "á", "ó", "ú", "é", "í", "Ã", "Õ", "Á", "É", "Í", "Ó", "Ú", "Â", "Ê", "Î", "Ô", "Û"}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma frase")
	scanner.Scan()
	frase := scanner.Text()
	for i := 0; i < len(vogais); i++ {
		frase = strings.ReplaceAll(frase, string(vogais[i]), "*")
	}
	fmt.Println("Nova frase sem vogais", frase)
}
