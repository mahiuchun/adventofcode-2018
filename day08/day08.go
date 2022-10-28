package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func WalkImpl(numbers []int) (len int, sum int) {
	len = 0
	sum = 0
	off := 2
	for i := 0; i < numbers[0]; i++ {
		l, s := WalkImpl(numbers[off:])
		off += l
		sum += s
	}
	len = off
	for i := off; i < off+numbers[1]; i++ {
		sum += numbers[i]
		len += 1
	}
	return
}

func Walk(numbers []int) (sum int) {
	_, sum = WalkImpl(numbers)
	return
}

func main() {
	r, err := os.Open("day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	numbers := make([]int, 0)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, x)
	}

	fmt.Println(Walk(numbers))
}
