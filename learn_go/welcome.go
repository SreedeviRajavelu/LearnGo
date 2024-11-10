package main // every code lives inside the package
// all files in the same folder on the disc should be in the same package

// package main is special - it will make your code compile to an executable and not a library

import (
	"fmt" // fmt package is for formatted output
)

// go is a simple language and most functionalities are inside other packages
// we need to import

func main() {
	fmt.Println("Welcome Gophers !")
	mean() // call the mean function
} // always prefix the package name before the function we call

// every function or symbol you use from another package start with an upper case
