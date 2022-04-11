package main

import (
	"fmt"
	"strings"
)

type Pizza struct {
	Name string `json:"name"`
	Price float64 `json:"price"`


}

func test() {
		pizzas := []Pizza{
			{Name:"La buena pizza", Price:18.34},
			{Name:"La not buena pizza", Price:8.67},
					}
	for _, pizza := range pizzas {
		fmt.Printf("%v and %v \n", pizza.Name, pizza.Price)
	}		
	fmt.Println()

}

func main() {
	test()
	var books = []string{}
	books = append(books, "@added")
	isValidAddition := len(books[0]) > 3 && len(books[0]) < 40
	isvalidEmail := strings.Contains(books[0], "@")
	fmt.Printf("%v", isValidAddition)
	fmt.Printf("%v", isvalidEmail)


	fmt.Printf("%v", books)
}
