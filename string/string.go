package main

import (
	"fmt"
)

func string() {
	book := "Simply Fly"
	fmt.Println(book)
	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	//uint8 is a byte

	// strings in go are immutable
	//book[0] = 116

	// slice(start,end) 0 based, empty range
	fmt.Println(book[7:9])

	fmt.Println(book[7:])

	fmt.Println(book[:6])

	// use + to concatenate string
	fmt.Println("Captain Gopinath wrote " + book)

	// unicode
	fmt.Println("It was 1/2 price")

	fmt.Println("------------")

	fmt.Println("TRIBUTE")
	// multi line
	tribute := `
	Soorarai Pottru
	Captain Gopinath
	Suriya
	Sudha Kongara
	Maara Theme
	...
	`
	fmt.Println(tribute)

}

func main() {
	string()
}
