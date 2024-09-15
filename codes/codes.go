package codes

import "math"

func Palindrome(a string) bool {
	strlen := len(a) - 1
	for index, _ := range a {
		if index == strlen || index > strlen {
			return true
		}
		if a[index] != a[strlen] {
			return false
		}
		strlen--
	}
	return true
}

func RevNum(a int) int {
	num := 0
	for a > 0 {
		temp := a % 10
		a -= temp
		a /= 10
		num = num*10 + temp
	}
	return num
}

func Fibonacci(a int) {
	num := 0
	oldNum := 1
	for a > 0 {
		if num == 0 {
			println(0)
			println(1)
			num, oldNum = oldNum+num, 1
			a -= 2
		}
		println(num)
		num, oldNum = oldNum+num, num
		a--
	}
}

func AmstrongNum(a int) bool {
	num, final, b := 0, 0, a
	for a > 0 {
		temp := a % 10
		a -= temp
		a /= 10
		num++
	}
	a = b
	for b > 0 {
		temp := b % 10
		b -= temp
		b /= 10
		final = final + int(math.Pow(float64(temp), float64(num)))
	}
	return a == final
}

func PrimeNum(a int) bool {
	for i := 2; i <= a/2; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func PerfectNum(a int) bool {
	num := 0
	for i := 1; i <= a/2; i++ {
		if a%i == 0 {
			num += i
		}
	}
	return num == a
}
