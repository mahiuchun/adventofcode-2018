package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// Observed by running first 200 generations
const Pre = 100
const Diff = 38
const Gens = 50000000000

func Lookup(m map[int]rune, i int) rune {
	if v, ok := m[i]; ok {
		return v
	} else {
		return '.'
	}
}

func main() {
	r, err := os.Open("day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	initial := scanner.Text()
	initial = initial[len("initial state: "):]
	scanner.Scan()

	rules := make(map[string]rune)
	for scanner.Scan() {
		rule := scanner.Text()
		rules[rule[0:5]] = rune(rule[9])
	}

	pots := make(map[int]rune)
	for i, c := range initial {
		pots[i] = c
	}

	for i := 0; i < Pre; i++ {
		next := make(map[int]rune)
		mini := math.MaxInt
		maxi := math.MinInt
		for k := range pots {
			if k < mini {
				mini = k
			}
			if k > maxi {
				maxi = k
			}
		}
		for j := mini - 2; j <= maxi+2; j++ {
			var b strings.Builder
			b.WriteRune(Lookup(pots, j-2))
			b.WriteRune(Lookup(pots, j-1))
			b.WriteRune(Lookup(pots, j))
			b.WriteRune(Lookup(pots, j+1))
			b.WriteRune(Lookup(pots, j+2))
			if c, ok := rules[b.String()]; ok {
				next[j] = c
			}
		}
		pots = next
	}

	sum := 0
	for k, v := range pots {
		if v == '#' {
			sum += k
		}
	}
	fmt.Println(sum + Diff*(Gens-Pre))
}
