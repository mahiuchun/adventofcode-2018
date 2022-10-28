package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	r, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	twos := 0
	threes := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		m := make(map[rune]int)
		for _, c := range scanner.Text() {
			m[c] += 1
		}
		d_twos := 0
		d_threes := 0
		for _, v := range m {
			if v == 2 {
				d_twos = 1
			} else if v == 3 {
				d_threes = 1
			}
		}
		twos += d_twos
		threes += d_threes
	}
	fmt.Println(twos * threes)
}
