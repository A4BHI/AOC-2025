package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := `987654321111111
			  811111111111119
			  234234234234278
			  818181911112111`

	allBanks := strings.Split(input, "\n")
	stack := [][]int{}
	for _, bank := range allBanks {
		for _, battery := range bank {
			if len(stack) == 0 {
				temp, _ := strconv.Atoi(string(battery))
				stack = append(stack, temp)
			} else {
				t, _ := strconv.Atoi(string(battery))
				if t >= stack[len(stack)-1] {

					if len(stack[:len(stack)-1]) < 12 {
						currentbat, _ := strconv.Atoi(string(battery))
						if stack[len(stack)-1] <= currentbat {
							stack := stack[:len(stack)-1]
							tm, _ := strconv.Atoi(string(battery))
							stack = append(stack, tm)
						} else {
							tm, _ := strconv.Atoi(string(battery))
							stack = append(stack, tm)
						}

					} else {
						fmt.Println(stack)
					}

				}
			}

		}
	}
	fmt.Println(stack)

	// stack = append(stack, 1, 2, 3, 5)
	// stack = stack[:len(stack)-1]
	// fmt.Println(stack)

}
