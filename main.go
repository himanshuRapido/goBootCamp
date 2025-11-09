package main

import (
	"fmt"
	// "github.com/fatih/color"
)

type ProductInfo struct {
	Product  string
	rating   int
	feedback string
	comments string
}

func getProducts() []ProductInfo {
	return []ProductInfo{
		{Product: "Watch", comments: "Good product", feedback: "I like this product", rating: 5},
		{Product: "Tv", comments: "Excellent product", feedback: "No complaints", rating: 4},
		{Product: "Mobile", comments: "Average product", feedback: "Some issues to be addressed", rating: 3},
		{Product: "Laptop", comments: "Not good product", feedback: "Many issues to be addressed", rating: 2},
		{Product: "iPhone", comments: "Please improve the product", feedback: "Not satisfied", rating: 1},
	}
}

func main() {
	// fmt.Println("Hello, World!")
	// color.Yellow("Hello, World!")

	products := getProducts()
	for i := 0; i < len(products); i++ {
		fmt.Printf("\n%d. Product: %s\n   Feedback: %s\n   Rating: %d⭐\n   Comments: %s\n",
			i+1, products[i].Product, products[i].feedback, products[i].rating, products[i].comments)
	}
	fmt.Println("Happy to see your feedback!")
	name, email, phone, feedback, rating := getUserPrompt()
	fmt.Println("Name: ", name)
	fmt.Println("Email: ", email)
	fmt.Println("Phone: ", phone)
	fmt.Println("Feedback: ", feedback)
	fmt.Println("Rating: ", rating)
	fmt.Printf("Thanks for your feedback! You rated %d⭐\n", rating)
}

func getUserPrompt() (string, string, int, string, int) {
	var (
		name     string
		email    string
		phone    int
		feedback string
		rating   int
	)
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter your email: ")
	fmt.Scanln(&email)
	fmt.Println("Enter your phone: ")
	fmt.Scanln(&phone)
	fmt.Println("Enter your feedback: ")
	fmt.Scanln(&feedback)
	fmt.Println("Enter your rating: ")
	fmt.Scanln(&rating)

	return name, email, phone, feedback, rating

}
