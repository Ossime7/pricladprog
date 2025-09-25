package main

import (
	"fmt"
	"time"
)

func floatoper() {
	//вычисление суммы и разности двух чисел с плавающей точкой
	var i float32 = 3.46
	var j float32 = 8.12
	var sum2 float32 = i + j
	fmt.Println("3.46 + 8.12 = ", sum2)

	var k float32 = 11.57
	var l float32 = 9.13
	var diff2 float32 = k - l
	fmt.Println("11.57 - 9.13 = ", diff2)
}

func main() {
	//вывод текущей даты
	now := time.Now()
	formattedDateTime := now.Format("2006-01-02 15:04:05")

	//создание и вывод переменных
	var f1 int = 5
	var f2 float64 = 3.54973
	var f3 string = "Hello, World!"
	var f4 bool = true

	fmt.Println(" ")
	fmt.Println("Текущая дата и время:", formattedDateTime)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)

	//создание и вывод переменных в краткой форме
	age := 20
	name := "Sofia"

	fmt.Println(" ")
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(" ")

	//арифметические операции с двумя целыми числами
	a := 11
	b := 9
	sum := a + b
	fmt.Println("11 + 9 = ", sum)

	c := 11
	d := 9
	diff := c - d
	fmt.Println("11 - 9 = ", diff)

	e := 11
	f := 9
	prod := e * f
	fmt.Println("11 * 9 = ", prod)

	g := 21
	h := 3
	del := g / h
	fmt.Println("21 / 3 = ", del)

	//вызов функции
	floatoper()

	//среднее значение трех чисел
	var ch1 int = 34
	var ch2 int = 67
	var ch3 int = 12
	var sred int = (ch1 + ch2 + ch3) / 3
	fmt.Println("среднее значение 34, 67, 12 = ", sred)
}
