package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	totalPoints       int = 100
	pointsPerQuestion int = 20
)

func main() {
	fmt.Println("Вітаємо у грі!")
	time.Sleep(3 * time.Second)

	for i := 5; i >= 1; i-- {
		fmt.Printf("Гра почнеться через: %v\n", i)
		time.Sleep(1 * time.Second)
	}

	startTime := time.Now()
	myPoints := 0
	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)

		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Не правильна відповідь!")
		} else {
			if ansInt == x+y {
				myPoints += pointsPerQuestion
				fmt.Println("Балів набрано: ", myPoints)
				fmt.Printf("Залишилось набрати: %v\n", totalPoints-myPoints)
			} else {
				fmt.Println("Спробуй ще Т_Т")
			}
		}
	}

	endTime := time.Now()
	timeSpent := endTime.Sub(startTime)
	fmt.Println("Вітаю! Ти впорався всього за: ", timeSpent)
	time.Sleep(5 * time.Second)
}
