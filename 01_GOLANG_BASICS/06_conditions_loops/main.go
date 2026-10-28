package main

import (
	"fmt"
)

func main() {
	score := 91

	if score > 90 {
		fmt.Println("Grade A")
	} else if score > 75 {
		fmt.Println("Grade B")
	} else if score > 45 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	item := 3
	price := 50

	if total := item * price; total > 200 {
		fmt.Println("Eligible for shipping")
	} else {
		fmt.Println("Not Eligible for shipping")
	}

	N := 10
	sum := 0

	for i := 1; i <= N; i++ {
		sum += i
	}

	fmt.Println("Sum of 10 n numbers ", sum)

	day := 5

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("wednesday")
	default:
		fmt.Println("unknown day")
	}

}
