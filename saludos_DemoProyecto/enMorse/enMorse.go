package enmorse

import (
	"fmt"
	"strings"

	"github.com/alwindoss/morse"
)

// Saluda en español
func Saludo() {
	fmt.Print("Te saludo en Morse: ")

	h := morse.NewHacker()
	morseCode, _ := h.Encode(strings.NewReader("Hola Mundo"))
	// Morse Code is: -.-. --- ...- . .-. - / - .... .. ... / - --- / -- --- .-. ... .
	fmt.Printf("%s\n", string(morseCode))
}
