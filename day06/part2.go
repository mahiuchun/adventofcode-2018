package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const Margin = 0

type Coord struct {
	x int
	y int
}

func Iabs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func main() {
	r, err := os.Open("day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	coords := make([]Coord, 0)
	scanner := bufio.NewScanner(r)
	minx, miny := math.MaxInt, math.MaxInt
	maxx, maxy := math.MinInt, math.MinInt
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		coords = append(coords, Coord{x: x, y: y})
		if x < minx {
			minx = x
		}
		if x > maxx {
			maxx = x
		}
		if y < miny {
			miny = y
		}
		if y > maxy {
			maxy = y
		}
	}

	count := 0
	for x := minx - Margin; x <= maxx+Margin; x++ {
		for y := miny - Margin; y <= maxy+Margin; y++ {
			tot := 0
			for _, coord := range coords {
				dist := Iabs(x-coord.x) + Iabs(y-coord.y)
				tot += dist
			}
			if tot < 10000 {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
