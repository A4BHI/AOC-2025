package main

import (
	"fmt"
	"strings"
)

func main() {

	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	slicesOfwalls := strings.Split(input, "\n")
	var wall [][]rune

	for _, slices := range slicesOfwalls {
		wall = append(wall, []rune(slices))
	}

	for _, row := range wall {
		fmt.Println(string(row))
	}

}
