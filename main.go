package main

import (
	"fmt"
	"math-game/domain"
	"math/rand"
	"strconv"
	"time"
)

const (
	totalPoints       int = 100
	pointsPerQuestion int = 100
)

var id uint64 = 1

func main() {
	var users []domain.User
	fmt.Println("Вітаємо у грі!")

	for {
		menu()
		point := ""
		fmt.Scan(&point)
		switch point {
		case "1":
			user := play()
			users = append(users, user)
		case "2":
			for _, user := range users {
				fmt.Printf("Id: %v Name: %s Time: %v\n",
					user.Id, user.Name, user.TimeSpent)
			}
		case "3":
			return
		default:
			fmt.Println("Зробіть коректний вибір")
		}
	}
}

func menu() {
	fmt.Println("1. Почати гру")
	fmt.Println("2. Рейтинг")
	fmt.Println("3. Вийти")
}

func play() domain.User {
	for i := 3; i >= 1; i-- {
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
	fmt.Println("Введіть своє ім'я: ")

	name := ""
	fmt.Scan(&name)

	user := domain.User{
		Id:        id,
		Name:      name,
		TimeSpent: timeSpent,
	}
	id++

	return user
}
