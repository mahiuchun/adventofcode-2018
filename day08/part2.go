package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func WalkImpl(numbers []int) (length int, val int) {
	length = 0
	val = 0
	off := 2
	vals := make([]int, 0)
	for i := 0; i < numbers[0]; i++ {
		l, v := WalkImpl(numbers[off:])
		off += l
		vals = append(vals, v)
	}
	if len(vals) == 0 {
		for i := off; i < off+numbers[1]; i++ {
			val += numbers[i]
		}
	} else {
		for i := off; i < off+numbers[1]; i++ {
			if 0 < numbers[i] && numbers[i] <= len(vals) {
				val += vals[numbers[i]-1]
			}
		}
	}
	length = off + numbers[1]
	return
}

func Walk(numbers []int) (val int) {
	_, val = WalkImpl(numbers)
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
