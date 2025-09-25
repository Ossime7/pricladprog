package mathutils

func Factorial(a int) int {

	if a < 0 {
		return -1 //ошибка: факториал отрицательного числа не определен
	}
	if a == 0 || a == 1 {
		return 1
	}
	result := 1
	for i := 2; i < a+1; i++ {
		result *= i
	}
	return result
}
