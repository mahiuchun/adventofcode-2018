package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func At(pic [][]rune, x, y int) rune {
	if y < 0 || y >= len(pic) {
		return '.'
	}
	if x < 0 || x >= len(pic[y]) {
		return '.'
	}
	return pic[y][x]
}

func Adj(pic [][]rune, x, y int) (a_t int, a_l int) {
	a_t = 0
	a_l = 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if At(pic, x+dx, y+dy) == '|' {
				a_t++
			} else if At(pic, x+dx, y+dy) == '#' {
				a_l++
			}
		}
	}
	return
}

func main() {
	r, err := os.Open("day18/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	pic := make([][]rune, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		pic = append(pic, []rune(scanner.Text()))
	}

	for i := 0; i < 10; i++ {
		next := make([][]rune, 0)
		for j := 0; j < len(pic); j++ {
			next = append(next, make([]rune, len(pic[j])))
			copy(next[j], pic[j])
		}
		for y := 0; y < len(pic); y++ {
			for x := 0; x < len(pic[y]); x++ {
				a_t, a_l := Adj(pic, x, y)
				switch pic[y][x] {
				case '.':
					if a_t >= 3 {
						next[y][x] = '|'
					}
				case '|':
					if a_l >= 3 {
						next[y][x] = '#'
					}
				case '#':
					if a_t < 1 || a_l < 1 {
						next[y][x] = '.'
					}
				default:
					panic("???")
				}
			}
		}
		pic = next
	}
	n_t, n_l := 0, 0
	for y := 0; y < len(pic); y++ {
		for x := 0; x < len(pic[y]); x++ {
			switch pic[y][x] {
			case '|':
				n_t++
			case '#':
				n_l++
			}
		}
	}
	fmt.Printf("%v * %v = %v\n", n_t, n_l, n_t*n_l)
}
