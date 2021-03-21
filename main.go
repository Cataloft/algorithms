package main

import (
	"fmt"
)

func main() {
	var d int
	fmt.Println("Введите алгоритм, который хотите использовать")
	fmt.Scan(&d)
	switch d {
	case 1:
		var num, del int
		fmt.Println("Введите число, потом число, которое хотите удалить из исходного числа")
		fmt.Scan(&num, &del)
		fmt.Println("Итоговое число: " ,deleteNum(num, del))
	}
}