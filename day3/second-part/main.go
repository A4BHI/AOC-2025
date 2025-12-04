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

	input = strings.TrimSpace(input)
	banks := strings.Split(input, "\n")

	var result []string
	max := 12

	for _, bank := range banks {
		bank = strings.TrimSpace(bank)
		toRemove := len(bank) - max
		stack := []rune{}

		for _, ch := range bank {
			fmt.Println(len(stack))
			for len(stack) > 0 && toRemove > 0 && ch > stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
				toRemove--
			}
			stack = append(stack, ch)
		}

		stack = stack[:max]
		result = append(result, string(stack))
	}

	var sum int
	for _, v := range result {
		num, _ := strconv.Atoi(v)
		sum += num
	}

	fmt.Println(result)
	fmt.Println(sum)
}
