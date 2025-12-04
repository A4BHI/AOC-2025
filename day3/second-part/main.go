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
	allBanks := strings.Split(input, "\n")

	var result []string
	for _, bank := range allBanks {
		stack := []rune{}
		max := 12

		for i, ch := range bank {
			bank = strings.TrimSpace(bank)
			remaining := len(bank) - i - 1

			for len(stack) > 0 && ch > stack[len(stack)-1] && len(stack)-1+remaining >= max {
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, ch)
		}

		if len(stack) > max {
			stack = stack[:max]
		}

		result = append(result, string(stack))
	}
	var sum int
	for _, v := range result {
		bank, _ := strconv.Atoi(v)
		sum += bank
	}
	fmt.Println(result)
	fmt.Println(sum)

	// stack = append(stack, 1, 2, 3, 5)
	// stack = stack[:len(stack)-1]
	// fmt.Println(stack)

}
