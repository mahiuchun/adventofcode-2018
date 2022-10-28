package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	px int
	py int
	vx int
	vy int
}

func main() {
	r, err := os.Open("day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`^position=\<\s*(\S+),\s*(\S+)\> velocity=\<\s*(\S+),\s*(\S+)\>$`)

	var points []Point
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var pt Point
		matches := re.FindStringSubmatch(scanner.Text())
		pt.px, _ = strconv.Atoi(matches[1])
		pt.py, _ = strconv.Atoi(matches[2])
		pt.vx, _ = strconv.Atoi(matches[3])
		pt.vy, _ = strconv.Atoi(matches[4])
		points = append(points, pt)
	}
	for t := 0; ; t++ {
		minx, miny, maxx, maxy := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
		for _, pt := range points {
			if pt.px < minx {
				minx = pt.px
			}
			if pt.px > maxx {
				maxx = pt.px
			}
			if pt.py < miny {
				miny = pt.py
			}
			if pt.py > maxy {
				maxy = pt.py
			}
		}

		w := maxx - minx + 1
		h := maxy - miny + 1
		if w <= 10 || h <= 10 {
			for row := 0; row < h; row++ {
				buf := make([]byte, w)
				for i := 0; i < w; i++ {
					buf[i] = ' '
				}
				for _, pt := range points {
					if pt.py-miny == row {
						buf[pt.px-minx] = '#'
					}
				}
			}
			fmt.Println(t)
			break
		}

		for i := range points {
			points[i].px += points[i].vx
			points[i].py += points[i].vy
		}
	}
}
