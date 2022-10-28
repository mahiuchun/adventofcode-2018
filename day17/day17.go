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

var (
	springX = 500
	springY = 0
)

type Scan struct {
	minX int
	maxX int
	minY int
	maxY int
}

func ParseRange(s string, min *int, max *int) {
	parts := strings.Split(s, "..")
	if len(parts) == 1 {
		*min, _ = strconv.Atoi(parts[0])
		*max = *min
	} else {
		*min, _ = strconv.Atoi(parts[0])
		*max, _ = strconv.Atoi(parts[1])
	}
}

func Dfs(x, y int, pic [][]rune, limit *Scan, count *int) {
	if y+1 <= limit.maxY && pic[y+1-limit.minY][x-limit.minX] == '.' {
		pic[y+1-limit.minY][x-limit.minX] = '|'
		*count += 1
		Dfs(x, y+1, pic, limit, count)
		x_left, x_right := x, x
		for x_left >= limit.minX {
			if pic[y+1-limit.minY][x_left-limit.minX] == '#' {
				break
			}
			x_left--
		}
		for x_right <= limit.maxX {
			if pic[y+1-limit.minY][x_right-limit.minX] == '#' {
				break
			}
			x_right++
		}
		if x_left >= limit.minX && x_right <= limit.maxX && y+2 <= limit.maxY {
			support := true
			for x := x_left + 1; x <= x_right-1; x++ {
				if pic[y+2-limit.minY][x-limit.minX] != '#' && pic[y+2-limit.minY][x-limit.minX] != '~' {
					support = false
					break
				}
			}
			if support {
				for x := x_left + 1; x <= x_right-1; x++ {
					pic[y+1-limit.minY][x-limit.minX] = '~'
				}
			}
		}
	}
	if y+1 <= limit.maxY && (pic[y+1-limit.minY][x-limit.minX] == '#' || pic[y+1-limit.minY][x-limit.minX] == '~') {
		if x-1 >= limit.minX && pic[y-limit.minY][x-1-limit.minX] == '.' {
			pic[y-limit.minY][x-1-limit.minX] = '|'
			*count += 1
			Dfs(x-1, y, pic, limit, count)
		}
		if x+1 <= limit.maxX && pic[y-limit.minY][x+1-limit.minX] == '.' {
			pic[y-limit.minY][x+1-limit.minX] = '|'
			*count += 1
			Dfs(x+1, y, pic, limit, count)
		}
	}
}

func main() {
	r, err := os.Open("day17/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]Scan, 0)
	scanner := bufio.NewScanner(r)
	min_x, max_x, min_y, max_y := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ", ")
		if strings.HasPrefix(parts[0], "y=") {
			parts[0], parts[1] = parts[1], parts[0]
		}
		var scan Scan
		ParseRange(parts[0][2:], &scan.minX, &scan.maxX)
		ParseRange(parts[1][2:], &scan.minY, &scan.maxY)
		data = append(data, scan)
		if scan.minX < min_x {
			min_x = scan.minX
		}
		if scan.minY < min_y {
			min_y = scan.minY
		}
		if scan.maxX > max_x {
			max_x = scan.maxX
		}
		if scan.maxY > max_y {
			max_y = scan.maxY
		}
	}
	min_x--
	max_x++

	w := max_x - min_x + 1
	h := max_y - min_y + 1
	pic := make([][]rune, h)
	for i := 0; i < h; i++ {
		pic[i] = make([]rune, w)
		for j := 0; j < w; j++ {
			pic[i][j] = '.'
		}
	}

	for _, scan := range data {
		for i := scan.minY - min_y; i <= scan.maxY-min_y; i++ {
			for j := scan.minX - min_x; j <= scan.maxX-min_x; j++ {
				pic[i][j] = '#'
			}
		}
	}

	springY = min_y - 1
	count := 0
	Dfs(springX, springY, pic, &Scan{minX: min_x, minY: min_y, maxX: max_x, maxY: max_y}, &count)
	fmt.Println(count)
}
