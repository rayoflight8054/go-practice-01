package main

import "fmt"

func main() {
	const course = "Бэкенд-разработка на языке Go"
	var name string = "Умный"
	var group string = "ТЭП-03-26"
	var completed int = 0
	var interes bool = true

	fmt.Printf("Участник %s\n", name)
	fmt.Println("Группа:", group)
	fmt.Println("Курс", course)
	fmt.Println("Завершено работ:", completed)
	fmt.Println("Интерес к бэкенду", interes)
}
