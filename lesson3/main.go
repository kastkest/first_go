package main

import (
	"fmt"
)

func main() {
	var (
		first  float32
		second float32
		total  int64
	)

	fmt.Scan(&first, &second)
	total = int64(first + second)
	if total%2.0 == 0 {
		fmt.Println("ЧЁТНОЕ")
	} else {
		fmt.Println("НЕЧЁТНОЕ")
	}
}
