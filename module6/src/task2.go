package main

import "fmt"

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Сумма двух чисел по единице.")
	fmt.Println("------------------------------")

	var firstNum, secondNum int

	// Ввод данных.
	fmt.Print("Введите первое число: ")
	fmt.Scan(&firstNum)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&secondNum)

	sumNum := firstNum + secondNum
	fmt.Print("Сумма чисел равна: ", sumNum, ".\n")
	for firstNum != sumNum {
		firstNum++
		fmt.Print("Первое число: ", firstNum, "; \n")
	}

	fmt.Print("Цикл закончен: первое число выросло до = ", firstNum, " и теперь ровно сумме введенных чисел.")
}