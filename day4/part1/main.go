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
	var begin int
	var limit int
	var counter int = 0
	for i := 0; i < len(wall); i++ {
		// fmt.Println(string(wall[i]))
		for j := 0; j < len(wall[i]); j++ {
			if wall[i][j] == '@' {

				k := j
				if j == 0 {
					begin = len(wall[i][:j])
				} else {
					begin = len(wall[i][:j-1])
				}

				if j == len(wall[i])-1 {
					limit = len(wall[i][:j])
				} else {
					limit = len(wall[i][:j+1])
				}

				for m := begin; m < limit-1; m++ {
					// fmt.Println(begin)
					// fmt.Println(limit)
					// if j == k {
					// 	continue
					// }
					// fmt.Println("Found0")
					// fmt.Println(string(wall[i][limit]))
					if wall[i][begin] == '@' || wall[i][begin] == 'X' {
						counter++
						// wall[i][begin] = 'B'
						// fmt.Println("Found2")

					}

					if wall[i][limit] == '@' || wall[i][limit] == 'X' {
						counter++
						// fmt.Println()
						// wall[i][limit] = 'E'
						// fmt.Println("Found")
					}

				}
				// fmt.Println(counter)

				if counter == 2 {
					wall[i][k] = 'X'
					// fmt.Println("Couter value 2")
				}

				if counter == 1 {
					// fmt.Println("Couter value 1")
				}

				counter = 0
				// fmt.Println(counter)
				// if wall[i][j+1]

			} else {
				fmt.Println("NOP")
			}

		}
		fmt.Println()
	}

	for _, row := range wall {
		fmt.Println(string(row))
	}
	// fmt.Println(wall)

}
