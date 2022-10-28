package main

import (
	"fmt"
	"strconv"
	"strings"
)

const Input = 890691

func main() {
	recipes := []int{3, 7}
	elves := []int{0, 1}
	for t := 0; t < Input+10; t++ {
		sum := 0
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
	}
	var b strings.Builder
	for _, n := range recipes[Input : Input+10] {
		b.WriteString(strconv.Itoa(n))
	}
	fmt.Println(b.String())
}
