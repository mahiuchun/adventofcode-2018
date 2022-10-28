package main

import "fmt"

const Input = 6303
const Size = 300

func main() {
	grid := make([][]int, Size)
	part := make([][]int, Size+1)
	for i := 0; i < Size; i++ {
		grid[i] = make([]int, Size)
		part[i] = make([]int, Size+1)
	}
	part[Size] = make([]int, Size+1)
	for y := 1; y <= Size; y++ {
		for x := 1; x <= Size; x++ {
			rack_id := x + 10
			power := rack_id*y + Input
			power *= rack_id
			power /= 100
			power %= 10
			power -= 5
			grid[y-1][x-1] = power
			part[y][x] = power + part[y-1][x] + part[y][x-1] - part[y-1][x-1]
		}
	}

	best, bestx, besty, bestsz := 0, -1, -1, -1
	for y := 1; y <= Size; y++ {
		for x := 1; x <= Size; x++ {
			for sz := 1; sz <= Size; sz++ {
				if y+sz-1 > Size || x+sz-1 > Size {
					break
				}
				sum := part[y+sz-1][x+sz-1] - part[y+sz-1][x-1] - part[y-1][x+sz-1] + part[y-1][x-1]
				if sum > best {
					best = sum
					bestx = x
					besty = y
					bestsz = sz
				}
			}
		}
	}
	fmt.Printf("%v,%v,%v\n", bestx, besty, bestsz)
}
