package main

import (
	"fmt"
	"log"
	"math"
)

const PI = 3.1415

func main()  {
	var change uint8
	fmt.Println("Введите число - номер задачи (1-3):")
	fmt.Println("1 - Программа для вычисления площади прямоугольника.")
	fmt.Println("2 - Программа, вычисляющуя диаметр и длину окружности по заданной площади круга.")
	fmt.Println("3 - Вывод цифр, соответствующих количеству сотен, десятков и единиц в трехзначном числе.")
	_, err := fmt.Scan(&change)
	if err != nil {
		fmt.Println(err)
	}
	switch change {
	case 1:
		square()
	case 2:
		circle()
	case 3:
		digits()
	default:
		fmt.Println("Необходимо ввести число от 1 до 3")
	}

}

func square()  {
	var a, b float32
	fmt.Println("Введите поочередно длины сторон прямоугольника")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Сторона а = %v, Сторона b = %v, Площадь равна: %v", a, b, a * b)
}

func circle()  {
	var s float64
	fmt.Println("Введите площадь круга:")
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Println(err)
	}

	r := math.Sqrt(s / PI)
	c := PI * 2 * r
	d := r * 2

	fmt.Printf("длина окружности равна: %v\n", c)
	fmt.Printf("диаметр окружности равен: %v\n", d)
}

func digits()  {
	var abc uint16
	fmt.Println("Введите трехзначное число")
	_, err := fmt.Scan(&abc)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("количество сотен: %v\nколичество десятков: %v\nколичество единиц: %v", abc / 100, abc / 10 % 10,
		abc % 10 % 10)
}