package star

func LeftTriangle(a int) {
	for i := 0; i < a; i++ {
		for j := 0; j <= i; j++ {
			print("*")
		}
		println("")
	}
}

func RightTriangle(a int) {
	for i := 0; i < a; i++ {
		for k := a - 1; k >= i; k-- {
			print(" ")
		}
		for j := 0; j <= i; j++ {
			print("*")
		}
		println("")
	}
}

func Heart(a int) {
	z := 2*a - 1
	for i := 0; i < 2; i++ {
		for j := 0; j < z; j++ {
			if i == 0 {
				if j == 0 {
					print(" ")
				} else if j == a-2 {
					print("   ")
					j += 3
				} else {
					print("*")
				}
			} else if i == 1 {
				if j == a-1 {
					print(" ")
				} else {
					print("*")
				}
			}
		}
		println("")
	}
	for i := 0; i < a; i++ {
		for j := 0; j < i; j++ {
			print(" ")
		}
		for k := 0; k < z-(2*i); k++ {
			print("*")
		}
		println("")
	}
}
