package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	Left = iota
	Straight
	Right
)

type Cart struct {
	X      int
	Y      int
	Dx     int
	Dy     int
	IntDir int
}

type ByPos []Cart

func (a ByPos) Len() int      { return len(a) }
func (a ByPos) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPos) Less(i, j int) bool {
	if a[i].Y != a[j].Y {
		return a[i].Y < a[j].Y
	} else {
		return a[i].X < a[j].X
	}
}

func main() {
	r, err := os.Open("day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var tracks []string
	var carts []Cart
	scanner := bufio.NewScanner(r)
	y := 0
	for scanner.Scan() {
		var b strings.Builder
		for x, c := range scanner.Text() {
			switch c {
			case '<':
				b.WriteRune('-')
				carts = append(carts, Cart{X: x, Y: y, Dx: -1, Dy: 0, IntDir: Left})
			case '>':
				b.WriteRune('-')
				carts = append(carts, Cart{X: x, Y: y, Dx: 1, Dy: 0, IntDir: Left})
			case '^':
				b.WriteRune('|')
				carts = append(carts, Cart{X: x, Y: y, Dx: 0, Dy: -1, IntDir: Left})
			case 'v':
				b.WriteRune('|')
				carts = append(carts, Cart{X: x, Y: y, Dx: 0, Dy: 1, IntDir: Left})
			default:
				b.WriteRune(c)
			}
		}
		tracks = append(tracks, b.String())
		y += 1
	}

	for {
		collided := false
		for i := range carts {
			carts[i].X += carts[i].Dx
			carts[i].Y += carts[i].Dy
			switch tracks[carts[i].Y][carts[i].X] {
			case '/':
				carts[i].Dx, carts[i].Dy = -carts[i].Dy, -carts[i].Dx
			case '\\':
				carts[i].Dx, carts[i].Dy = carts[i].Dy, carts[i].Dx
			case '+':
				switch carts[i].IntDir {
				case Left:
					carts[i].IntDir = Straight
					carts[i].Dx, carts[i].Dy = carts[i].Dy, -carts[i].Dx
				case Straight:
					carts[i].IntDir = Right
				case Right:
					carts[i].IntDir = Left
					carts[i].Dx, carts[i].Dy = -carts[i].Dy, carts[i].Dx
				}
			case '-':
			case '|':
			default:
				panic("Something is wrong")
			}
			for j := range carts {
				if j == i {
					continue
				}
				if carts[i].X == carts[j].X && carts[i].Y == carts[j].Y {
					collided = true
					break
				}
			}
			if collided {
				fmt.Printf("%v,%v\n", carts[i].X, carts[i].Y)
				break
			}
		}
		if collided {
			break
		}
		sort.Sort(ByPos(carts))
	}
}
