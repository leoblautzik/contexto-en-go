package division

func Cociente(a, b int) int {
	cociente := 1
	if a-b < 0 {
		return 0
	}
	return cociente + Cociente(a-b, b)
}

func Resto(a, b int) int {
	resto := 0
	if a-b < 0 {
		return a
	}
	return resto + Resto(a-b, b)
}

func Division(a,b int)(int,int){
    return Cociente(a, b ),Resto(a , b)
}
