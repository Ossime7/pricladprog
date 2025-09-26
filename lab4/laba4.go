package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// функция для расчета среднего возраста
func calculateAverageAge(people map[string]int) float32 {
	if len(people) == 0 {
		return 0
	}

	sum := 0
	for _, age := range people {
		sum += age
	}

	return float32(sum) / float32(len(people))
}

func main() {
	//создание карты
	people := map[string]int{
		"Анна":  25,
		"Иван":  30,
		"Мария": 28,
	}
	fmt.Println("Карта людей до добавления:")
	for name, age := range people {
		fmt.Printf("%s: %d лет\n", name, age)
	}
	fmt.Println()

	//добавление нового человека
	people["Петр"] = 35

	//вывод всех записей
	fmt.Println("Карта людей после добавления:")
	for name, age := range people {
		fmt.Printf("%s: %d лет\n", name, age)
	}
	fmt.Println()

	//расчет среднего возраста
	average := calculateAverageAge(people)
	fmt.Printf("Средний возраст: %.2f лет\n\n", average)

	//удаление записи по имени
	delete(people, "Иван")

	fmt.Println("Карта людей после удаления Ивана:")
	for name, age := range people {
		fmt.Printf("%s: %d лет\n", name, age)
	}
	fmt.Println()

	//преобразование строки в верхний регистр
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	upperCase := strings.ToUpper(input)

	fmt.Printf("Результат: %s\n\n", upperCase)

	//сумма нескольких чисел
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Введите целые числа через пробел: ")
	input2, _ := reader2.ReadString('\n')

	//разбиваем на части
	input2 = strings.TrimSpace(input2)
	numbers := strings.Fields(input2)

	sum := 0
	validNumbers := 0

	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			sum += num
			validNumbers++
		}
	}

	fmt.Printf("Количество введенных чисел: %d\n", validNumbers)
	fmt.Printf("Сумма чисел: %d\n\n", sum)

	//вывод массива в обратном порядке
	reader3 := bufio.NewReader(os.Stdin)
	fmt.Print("Введите целые числа через пробел: ")
	input3, _ := reader3.ReadString('\n')

	//разбиваем на части
	input3 = strings.TrimSpace(input3)
	numberStrings := strings.Fields(input3)

	//создание массива чисел
	numbers2 := make([]int, 0, len(numberStrings))

	for _, numStr := range numberStrings {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			numbers2 = append(numbers2, num)
		}
	}

	fmt.Println("Исходный массив:", numbers2)

	//вывод в обратном порядке
	fmt.Print("Массив в обратном порядке: [")
	for i := len(numbers2) - 1; i >= 0; i-- {
		fmt.Print(numbers2[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}
