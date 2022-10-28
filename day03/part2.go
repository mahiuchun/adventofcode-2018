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

	re := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	scanner := bufio.NewScanner(r)
	claims := make([]string, 0, 1349)
	for scanner.Scan() {
		claims = append(claims, scanner.Text())
	}
	for _, claim := range claims {
		matches := re.FindStringSubmatch(claim)
		left, _ := strconv.Atoi(matches[2])
		top, _ := strconv.Atoi(matches[3])
		w, _ := strconv.Atoi(matches[4])
		h, _ := strconv.Atoi(matches[5])
		for i := top; i < top+h; i++ {
			for j := left; j < left+w; j++ {
				mat[i][j] += 1
			}
		}
	}

	for _, claim := range claims {
		matches := re.FindStringSubmatch(claim)
		num, _ := strconv.Atoi(matches[1])
		left, _ := strconv.Atoi(matches[2])
		top, _ := strconv.Atoi(matches[3])
		w, _ := strconv.Atoi(matches[4])
		h, _ := strconv.Atoi(matches[5])

		overlap := false
		for i := top; i < top+h; i++ {
			for j := left; j < left+w; j++ {
				if mat[i][j] > 1 {
					overlap = true
					break
				}
			}
			if overlap {
				break
			}
		}

		if !overlap {
			fmt.Println(num)
		}
	}
}
