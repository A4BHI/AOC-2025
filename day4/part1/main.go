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
	var k int
	for i := 0; i < len(wall); i++ {

		for j := 0; j < len(wall[i]); j++ {
			if wall[i][j] == '@' {

				k = j
				if j == 0 {
					begin = len(wall[i][:j])
				} else {
					if j == 0 {
						begin = len(wall[i][:j])
					} else {
						begin = len(wall[i][:j-1])
					}

				}

				if j == len(wall[i])-1 {
					limit = len(wall[i][:j])
				} else {

					if j == 0 {
						limit = len(wall[i][:j])
					} else {
						limit = len(wall[i][:j+1])
					}
				}

				for m := begin; m < limit-1; m++ {
					fmt.Println("Iteration", i, "Begin: ", begin, " Limit: ", limit-1)
					if wall[i][begin] == '@' || wall[i][begin] == 'X' {

						counter++

					}

					if wall[i][limit] == '@' || wall[i][limit] == 'X' {
						counter++

					}

					if i != 0 {
						if wall[i-1][k] == '@' || wall[i-1][k] == 'X' {
							counter++
						}

						if wall[i-1][begin] == '@' || wall[i-1][begin] == 'X' {
							counter++
						}
						if wall[i-1][limit] == '@' || wall[i-1][limit] == 'X' {
							counter++
						}

					}

					if i != len(wall)-1 {
						if wall[i+1][k] == '@' || wall[i+1][k] == 'X' {
							counter++
						}

						if wall[i+1][begin] == '@' || wall[i+1][begin] == 'X' {
							counter++
						}
						if wall[i+1][limit] == '@' || wall[i+1][limit] == 'X' {
							counter++
						}

					}

				}

			}
		}
		fmt.Println("Counter: ", counter)

		if counter < 4 {
			wall[i][k] = 'X'

		}

		counter = 0

	}

	for _, row := range wall {
		fmt.Println(string(row))
	}
	var sum int
	for i := 0; i < len(wall); i++ {
		for j := 0; j < len(wall[i]); j++ {
			if wall[i][j] == 'X' {
				sum++
			}
		}

	}
	fmt.Println(sum)

}
