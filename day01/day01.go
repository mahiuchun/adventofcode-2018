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
	tot := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		tot += i
	}
	fmt.Println(tot)
}
