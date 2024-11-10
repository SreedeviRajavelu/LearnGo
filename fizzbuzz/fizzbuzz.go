/* Print nos btw 1 - 20 but
- if num is divisible by 3 print "fizz"
- if num is divisible by 5, print "buzz"
- if num is divisible by 3 and 5, print "fizz buzz"

*/

package main

import (
	"fmt"
)

func fizzbuzz() {
	// i:=20
	for i := 0; i < 21; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("fizz")
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz()
}
