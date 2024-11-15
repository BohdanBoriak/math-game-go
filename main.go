package main

import (
	"encoding/json"
	"fmt"
	"math-game/domain"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

const (
	totalPoints       int = 100
	pointsPerQuestion int = 100
)

var id uint64 = 1

func main() {
	users := getUsers()
	for _, user := range users {
		if user.Id >= id {
			id = user.Id + 1
		}
	}
	fmt.Println("Вітаємо у грі!")

	for {
		menu()
		point := ""
		fmt.Scan(&point)
		switch point {
		case "1":
			user := play()
			users = getUsers()
			users = append(users, user)
			sortAndSave(users)
		case "2":
			users = getUsers()
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

func sortAndSave(users []domain.User) {
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].TimeSpent < users[j].TimeSpent
	})

	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("sortAndSave -> os.OpenFile: %s", err)
		return
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		fmt.Printf("sortAndSave -> encoder.Encode: %s", err)
		return
	}
}

func getUsers() []domain.User {
	file, err := os.Open("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create("users.json")
			if err != nil {
				fmt.Printf("getUsers -> os.Create: %s\n", err)
				return nil
			}
			return nil
		}
		fmt.Printf("getUsers -> os.Open: %s\n", err)
		return nil
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
	}(file)

	var users []domain.User
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		return nil
	}

	return users
}
