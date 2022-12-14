package main

import (
	"fmt"
	"math/rand"
	"time"
)

const lengthNew = 12

func generationArrSorted() (array [lengthNew]int) {
	rand.Seed(time.Now().UnixNano())
	var tempNumber, count, randomNumber int
	randomNumber = 3
	for index1 := 0; index1 < lengthNew; index1++ {
		tempNumber = rand.Intn(randomNumber-tempNumber) + tempNumber
		for index2, _ := range array {
			if array[index2] > tempNumber {
				count++
			}
		}
		if count == 0 {
			array[index1] = tempNumber
			randomNumber += 2
		} else {
			index1--
		}
		count = 0
	}
	return
}

func numberSearchNew(array [lengthNew]int, desiredNumberNew int) (result int) {
	result = -1
	min, max, middle := 0, length-1, 0
	for max >= min {
		middle = (max + min) / 2
		if array[middle] == desiredNumberNew {
			result = middle
			break
		} else if array[middle] > (max+min)/2 {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	return
}

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться).")
	fmt.Println()

	var desiredNumberNew, result int

	array := generationArrSorted()
	fmt.Printf("Используемый массив:\n%v\n", array)

	fmt.Print("Введите число: ")
	_, _ = fmt.Scan(&desiredNumberNew)

	result = numberSearchNew(array, desiredNumberNew)

	if result == -1 {
		fmt.Println("Такого числа нет в массиве!")
	} else {
		fmt.Printf("Индекс: %v\n", result)
	}
}
