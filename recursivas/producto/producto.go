package producto

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
