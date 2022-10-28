package main

import (
	"fmt"
	"strconv"
)

const Input = "890691"

func main() {
	recipes := []int{3, 7}
	elves := []int{0, 1}
	for {
		sum := 0
		ll := len(recipes)
		for i := range elves {
			sum += recipes[elves[i]]
		}
		combined := strconv.Itoa(sum)
		for i := range combined {
			n, _ := strconv.Atoi(combined[i : i+1])
			recipes = append(recipes, n)
		}
		for i := range elves {
			elves[i] += 1 + recipes[elves[i]]
			elves[i] %= len(recipes)
		}
		index := -1
		for j := ll; j <= len(recipes); j++ {
			accept := true
			for i := 0; i < len(Input); i++ {
				jj := j - len(Input) + i
				if jj < 0 || recipes[jj] != int(Input[i]-'0') {
					accept = false
					break
				}
			}
			if accept {
				index = j - len(Input)
				break
			}
		}
		if index >= 0 {
			fmt.Println(index)
			break
		}
	}
}
