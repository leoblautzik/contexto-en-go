package main

import (
	"fmt"

	"untref.edu/leoblau.recursivas/palindromo"
)

func main() {
	fmt.Println(palindromo.EsPalindromo("oso"))
	fmt.Println(palindromo.EsPalindromo("12345677654321"))
	fmt.Println(palindromo.EsPalindromo("o"))
	fmt.Println(palindromo.EsPalindromo("Ala"))
	fmt.Println(palindromo.EsPalindromo("Neuquen"))
}

/**
func productoItera(a int, b int) int {
	producto := 0
	for i := 0; i < b; i++ {
		producto += a
	}
	return producto
}

func productoRecu(a int, b int) int {
	if b == 0 || a == 0 {
		return 0
	}
	return a + productoRecu(a, b-1)

}
func cocienteIte(a, b int) int {
	cociente := 0
	a -= b
	for a >= 0 {
		cociente++
		a -= b
	}
	return cociente
}
func cocienteRecu(a, b int) int {
	cociente := 1
	if a-b < 0 {
		return 0
	}
	return cociente + cocienteRecu(a-b, b)
}

func Resto(a, b int) int {
	resto := 0
	if a-b < 0 {
		return a
	}
	return resto + Resto(a-b, b)
}

func ImprimirDecimal(n int) {
	if n >= 10 {
		ImprimirDecimal(n / 10)
	}
	fmt.Println(n % 10)
}
func ImprimirBinario(n int) {
	if n >= 2 {
		ImprimirDecimal(n / 2)
	}
	fmt.Println(n % 2)
}

func ImprimirEnBase(n int, base int) {
	if n >= base {
		ImprimirEnBase(n/base, base)
	}

	fmt.Println(string(digits[n%base]))
}
func SumaElementosDelArray(array []int) int {
	if len(array) <= 1 {
		return array[0]
	} else {
		a := array[:len(array)/2]
		b := array[len(array)/2:]
		return SumaElementosDelArray(a) + SumaElementosDelArray(b)
	}

}
	func producto(a int, b int) int {
		if b == 0 || a == 0 {
			return 0
		}
		return a + producto(a, b-1)
	}
*/
