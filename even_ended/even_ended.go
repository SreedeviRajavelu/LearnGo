/*
An even ended num has first and last digit to be the same

Count how many even-ended nums can result from a multiplication of two 4 digit nums
*/

package main

import (
	"fmt"
)

func main() {
	count := 0

	//for every pair of 4 digit numbers
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ { // dont count twice the same combination
			n := a * b

			// a *b is even ended
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				// count = count +1
				count++
			}
		}
	}
	fmt.Println(count)
}
