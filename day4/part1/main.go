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

	fmt.Println("#############################")

	for i := 0; i < len(wall); i++ {
		// fmt.Println(string(wall[i]))
		for j := 0; j < len(wall[i]); j++ {
			fmt.Print(string(wall[i][j]))
		}
		fmt.Println()
	}

}
