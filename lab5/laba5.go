package main

import (
	"fmt"
	"math"
)

// структура Person
type Person struct {
	name string
	age  int
}

// метод для вывода информации о человеке
func (p Person) PrintInfo() {
	fmt.Printf("Имя: %s, Возраст: %d лет\n", p.name, p.age)
}

// метод birthday для увеличения возраста
func (p *Person) Birthday() {
	p.age++
}

// структура Circle
type Circle struct {
	radius float64
}

// метод для вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// интерфейс Shape и его реализации
type Shape interface {
	Area() float64
}

// структура Rectangle
type Rectangle struct {
	width  float64
	height float64
}

// реализация интерфейса Shape для Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// функция для вывода площадей объектов
func printAreas(shapes []Shape) {
	for i, shape := range shapes {
		area := shape.Area()
		fmt.Printf("Фигура %d: площадь = %.2f\n", i+1, area)
	}
	fmt.Println()
}

// интерфейс Stringer и структура Book
type Stringer interface {
	String() string
}
type Book struct {
	title  string
	author string
	year   int
}

// реализация интерфейса Stringer для Book
func (b Book) String() string {
	return fmt.Sprintf("Книга: \"%s\" (автор: %s, год: %d)", b.title, b.author, b.year)
}

func main() {
	fmt.Println("=== Задача 1 ===")
	person1 := Person{name: "Анна", age: 25}
	person1.PrintInfo()
	person2 := Person{name: "Иван", age: 30}
	person2.PrintInfo()
	fmt.Println()

	fmt.Println("=== Задача 2 ===")
	fmt.Printf("До дня рождения: ")
	person1.PrintInfo()
	person1.Birthday()
	fmt.Printf("После дня рождения: ")
	person1.PrintInfo()
	fmt.Println()

	fmt.Println("=== Задача 3 ===")
	circle := Circle{radius: 5.0}
	area := circle.Area()
	fmt.Printf("Круг с радиусом %.2f: площадь = %.2f\n", circle.radius, area)
	fmt.Println()

	fmt.Println("=== Задача 4 ===")
	rectangle := Rectangle{width: 4.0, height: 6.0}

	var shape1 Shape = circle
	var shape2 Shape = rectangle

	fmt.Printf("Площадь круга через интерфейс: %.2f\n", shape1.Area())
	fmt.Printf("Площадь прямоугольника через интерфейс: %.2f\n", shape2.Area())
	fmt.Println()

	fmt.Println("=== Задача 5 ===")
	shapes := []Shape{
		Circle{radius: 3.0},
		Rectangle{width: 5.0, height: 2.0},
		Circle{radius: 7.5},
		Rectangle{width: 3.0, height: 3.0},
	}
	printAreas(shapes)

	fmt.Println("=== Задача 6 ===")
	book := Book{
		title:  "Война и мир",
		author: "Лев Толстой",
		year:   1869,
	}

	// Использование интерфейса Stringer
	var stringer Stringer = book
	fmt.Println(stringer.String())
}
