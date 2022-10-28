package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	r, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var changes []int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		changes = append(changes, i)
	}
	sofar := 0
	m := make(map[int]int)
	for {
		found := false
		for _, i := range changes {
			sofar += i
			_, prs := m[sofar]
			if prs {
				fmt.Println(sofar)
				found = true
				break
			}
			m[sofar] = 1
		}
		if found {
			break
		}
	}
}
