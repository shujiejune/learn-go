package main

import (
	"fmt"
)

type Product struct {
	id    int
	title string
	price float64
}

func newProduct(id int, title string, price float64) *Product {
	return &Product{
		id:    id,
		title: title,
		price: price,
	}
}

func main() {
	hobbies := [3]string{"cooking", "watching live stream", "writing fanfictions"}
	first2Hobbies := hobbies[:2]
	firstHobby1 := first2Hobbies[:1]
	firstHobby2 := first2Hobbies[0:1]
	last2Hobbies := first2Hobbies[1:3]

	fmt.Println(hobbies)
	fmt.Printf("First hobby: %v %v\n", firstHobby1, firstHobby2)
	fmt.Println(hobbies[1:])
	fmt.Println(last2Hobbies)

	fmt.Println("------")

	goals := []string{"build a backend server in Go", "master a new language"}
	fmt.Println(goals)
	goals[1] = "add a personal project to my resume"
	fmt.Println(goals)
	goals = append(goals, "nail a job")
	fmt.Println(goals)

	fmt.Println("------")

	prod1 := newProduct(1, "TUXEDO", 1105.0)
	prod2 := newProduct(2, "Neovim", 0.0)
	prodList := []Product{*prod1, *prod2}
	fmt.Println(prodList)
	prodList = append(prodList, *newProduct(3, "Weee", 35.0))
	fmt.Println(prodList)
}
