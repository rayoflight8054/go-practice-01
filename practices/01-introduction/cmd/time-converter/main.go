package main

import "fmt"

func main() {
	var mins int

	fmt.Print("Введите количество минут: ")
	_, err := fmt.Scan(&mins)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if mins < 0 {
		fmt.Println("Ошибка: число не должно быть отрицательным.")
		return
	}
	fmt.Println("Полных часов:", mins/60)
	fmt.Println("Оставшихся минут:", mins%60)
}
