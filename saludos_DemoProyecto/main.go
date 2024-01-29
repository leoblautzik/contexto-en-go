package main

import (
	enaleman "untref.edu/leoblau.saludos/enAleman"
	enespaniol "untref.edu/leoblau.saludos/enEspaniol"
	enfrances "untref.edu/leoblau.saludos/enFrances"
	eningles "untref.edu/leoblau.saludos/enIngles"
	enmorse "untref.edu/leoblau.saludos/enMorse"
)

func main() {
	enespaniol.Saludo()
	eningles.Saludo()
	enaleman.Saludo()
	enfrances.Saludo()
	enmorse.Saludo()
}
