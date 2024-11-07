package main

import (
	"fmt"
)

func main() {
	var (
		first  int
		second int
	)

	fmt.Scan(&first, &second)
	var line int = (first + second) * (first + second)

	fmt.Print(line)
}
