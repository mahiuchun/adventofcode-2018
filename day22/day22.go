package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	r, err := os.Open("day22/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	line1 := scanner.Text()
	depth, _ := strconv.Atoi(strings.Split(line1, ": ")[1])
	scanner.Scan()
	line2 := scanner.Text()
	target := strings.Split(line2, ": ")[1]
	coord_s := strings.Split(target, ",")
	xt, _ := strconv.Atoi(coord_s[0])
	yt, _ := strconv.Atoi(coord_s[1])
	tot := 0
	mat := make([][]int, yt+1)
	for y := 0; y <= yt; y++ {
		mat[y] = make([]int, xt+1)
		for x := 0; x <= xt; x++ {
			geologic := -1
			switch {
			case x == 0 && y == 0:
				geologic = 0
			case x == xt && y == yt:
				geologic = 0
			case y == 0:
				geologic = x * 16807
			case x == 0:
				geologic = y * 48271
			default:
				geologic = mat[y][x-1] * mat[y-1][x]
			}
			erosion := (geologic + depth) % 20183
			tot += erosion % 3
			mat[y][x] = erosion
		}
	}
	fmt.Println(tot)
}
