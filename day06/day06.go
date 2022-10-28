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
	area := make(map[string]int)
	infinite := make(map[string]int)
	for x := minx - 1; x <= maxx+1; x++ {
		for y := miny - 1; y <= maxy+1; y++ {
			closest := math.MaxInt
			var closestid string
			for _, coord := range coords {
				key := fmt.Sprintf("%v,%v", coord.x, coord.y)
				dist := Iabs(x-coord.x) + Iabs(y-coord.y)
				switch {
				case dist < closest:
					closest = dist
					closestid = key
				case dist == closest:
					closestid = ""
				}
			}
			if closestid == "" {
				continue
			}
			area[closestid] += 1
			if x < minx || x > maxx || y < miny || y > maxy {
				infinite[closestid] = 1
			}
		}
	}
	best := 0
	for key, val := range area {
		if _, prs := infinite[key]; prs {
			continue
		}
		if val > best {
			best = val
		}
	}
	fmt.Println(best)
}
