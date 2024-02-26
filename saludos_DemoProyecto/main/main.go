package main

import (
	"fmt"

	"untref.edu/leoblau.saludos/constantes/matematicas"
	"untref.edu/leoblau.saludos/constantes/stringas"
	"untref.edu/leoblau.saludos/enAleman"
	enespaniol "untref.edu/leoblau.saludos/enEspaniol"
	enfrances "untref.edu/leoblau.saludos/enFrances"
	eningles "untref.edu/leoblau.saludos/enIngles"
	enmorse "untref.edu/leoblau.saludos/enMorse"
)

// variable global
var cadena = "Saludito"

func main() {
	{
		// variable local
		entero := 5
		fmt.Println(entero)
		fmt.Println(cadena)
	}

	// fmt.Println(entero)

	fmt.Println(cadena)
	saludo()
	enespaniol.Saludo()
	eningles.Saludo()
	enaleman.Saludo()
	enfrances.Saludo()
	enmorse.Saludo()
	fmt.Println(matematicas.Ee)
	fmt.Println(stringas.Materia)
}
