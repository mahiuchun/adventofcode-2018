package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const Size = 1001

func main() {
	r, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	mat := make([][]int, Size)
	for i := 0; i < Size; i++ {
		mat[i] = make([]int, Size)
	}

	re := regexp.MustCompile(`^#\d+ @ (\d+),(\d+): (\d+)x(\d+)$`)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		left, _ := strconv.Atoi(matches[1])
		top, _ := strconv.Atoi(matches[2])
		w, _ := strconv.Atoi(matches[3])
		h, _ := strconv.Atoi(matches[4])
		for i := top; i < top+h; i++ {
			for j := left; j < left+w; j++ {
				mat[i][j] += 1
			}
		}
	}

	count := 0
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if mat[i][j] > 1 {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
