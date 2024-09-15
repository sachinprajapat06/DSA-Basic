package matrix

func Add(a [3][3]int, b [3][3]int) [3][3]int {
	c := [3][3]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

func Sub(a [3][3]int, b [3][3]int) [3][3]int {
	c := [3][3]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = a[i][j] - b[i][j]
		}
	}
	return c
}

func Mul(a [3][3]int, b [3][3]int) [3][3]int {
	c := [3][3]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = a[i][1]*b[1][j] + a[i][2]*b[2][j] + a[i][0]*b[0][j]
		}
	}
	return c
}
