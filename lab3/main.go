package main

import (
	"fmt"
	"main/mathutils"
	"main/stringutils"
	"math/rand"
	"time"
)

func findLongestString(strings []string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]
	for _, str := range strings {
		if len(str) > len(longest) {
			longest = str
		}
	}
	return longest
}

func main() {
	//проверка факториала
	var x int
	fmt.Scan(&x)
	fact := mathutils.Factorial(x)
	fmt.Println(fact)

	//переворот строки
	str := "Music is gorgeous!"
	reversed := stringutils.ReverseString(str)
	fmt.Printf("Оригинал: %s\nПеревернуто: %s\n", str, reversed)

	//создание массива и 5 чисел
	var arr [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)

	//создаем срез из массива
	slice := arr[:]
	fmt.Println("Срез из массива:", slice)

	//добавление элементов в срез
	slice = append(slice, 6, 7, 8)
	fmt.Println("После добавления 6,7,8:", slice)

	//удаление элемента по индексу
	indexToRemove := 2
	slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	fmt.Println("После удаления элемента с индексом 2:", slice)

	//срез из строк и нахождение самой длинной строки
	words := []string{"cat", "elephant", "dog", "giraffe", "mouse"}
	longest := findLongestString(words)
	fmt.Printf("Срез строк: %v\n", words)
	fmt.Printf("Самая длинная строка: '%s' (длина: %d)\n", longest, len(longest))

}
