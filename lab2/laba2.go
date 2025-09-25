package main

import (
	"fmt"
)

func pnz(x int) {
	//проверка введенного числа на положительность
	if x > 0 {
		fmt.Println("Число", x, "положительное")
	} else if x < 0 {
		fmt.Println("Число", x, "отрицательное")
	} else {
		fmt.Println("Число", x, "равно 0")
	}
}

func lenstr(str string) {
	length := len(str)
	fmt.Println(" ")
	fmt.Printf("Длина строки: %d\n", length)
}

func srznach(a, b int) {
	sred := (a + b) / 2
	fmt.Println("Среднее значение = ", sred)
}

func main() {
	//проверка введенного числа на четность
	var x int
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("Число", x, "четное")
	} else {
		fmt.Println("Число", x, "нечетное")
	}

	//вызов функции pnz
	pnz(x)

	//вывод чисел 1-10 с помощью for
	for i := 1; i < 11; i++ {
		fmt.Print(i, " ")
	}

	//длина строки
	str := "I wanna sleep!"
	lenstr(str)

	//структура Rectangle
	type Rectangle struct {
		Length int
		width  int
	}

	//вычисления площади прямоугольника
	rect := Rectangle{Length: 5, width: 3}
	s := rect.Length * rect.width
	fmt.Println("Площадь прямоугольника = ", s)

	//среднее значение двух чисел
	ch1 := 7
	ch2 := 4
	srznach(ch1, ch2)
}
