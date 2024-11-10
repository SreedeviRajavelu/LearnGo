// go slices
package main

import (
	"fmt"
)

func main() {
	//same type
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// length
	fmt.Println(len(loons)) // 3

	fmt.Println("--------")
	// 0 indexing
	fmt.Println(loons[1]) // daffy
}
