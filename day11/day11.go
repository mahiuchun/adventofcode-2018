package main

import "fmt"

const Input = 6303
const Size = 300

func main() {
	grid := make([][]int, Size)
	for i := 0; i < Size; i++ {
		grid[i] = make([]int, Size)
	}
	for y := 1; y <= Size; y++ {
		for x := 1; x <= Size; x++ {
			rack_id := x + 10
			power := rack_id*y + Input
			power *= rack_id
			power /= 100
			power %= 10
			power -= 5
			grid[y-1][x-1] = power
		}
	}
	best, bestx, besty := 0, -1, -1
	for y := 1; y <= Size-2; y++ {
		for x := 1; x <= Size-2; x++ {
			sum := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					sum += grid[y+i-1][x+j-1]
				}
			}
			if sum > best {
				best = sum
				bestx = x
				besty = y
			}
		}
	}
	fmt.Printf("%v,%v\n", bestx, besty)
}
