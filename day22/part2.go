package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	None  = 0
	Climb = 1
	Torch = 2
)

const (
	Rocky  = 0
	Wet    = 1
	Narrow = 2
)

var (
	memo  map[string]int
	xt    int
	yt    int
	depth int
)

type State struct {
	x    int
	y    int
	gear int
}

func (s *State) Same(o *State) bool {
	return s.x == o.x && s.y == o.y && s.gear == o.gear
}

func (s *State) Cannot() bool {
	if Lazy(s.x, s.y) == Rocky && s.gear == None {
		return true
	}
	if Lazy(s.x, s.y) == Wet && s.gear == Torch {
		return true
	}
	if Lazy(s.x, s.y) == Narrow && s.gear == Climb {
		return true
	}
	return false
}

func (s *State) Moves() []State {
	var res []State
	// change gear
	for g := 0; g < 3; g++ {
		if g == s.gear {
			continue
		}
		cand := State{x: s.x, y: s.y, gear: g}
		if cand.Cannot() {
			continue
		}
		res = append(res, cand)
	}
	// move
	for i := 1; i <= 1; i++ {
		if s.x-i < 0 {
			break
		}
		cand := State{x: s.x - i, y: s.y, gear: s.gear}
		if cand.Cannot() {
			break
		}
		res = append(res, cand)
	}
	for i := 1; i <= 1; i++ {
		cand := State{x: s.x + i, y: s.y, gear: s.gear}
		if cand.Cannot() {
			break
		}
		res = append(res, cand)
	}
	for i := 1; i <= 1; i++ {
		if s.y-i < 0 {
			break
		}
		cand := State{x: s.x, y: s.y - i, gear: s.gear}
		if cand.Cannot() {
			break
		}
		res = append(res, cand)
	}
	for i := 1; i <= 1; i++ {
		cand := State{x: s.x, y: s.y + i, gear: s.gear}
		if cand.Cannot() {
			break
		}
		res = append(res, cand)
	}
	if len(res) == 0 {
		panic(fmt.Sprintf("No move from %v", s))
	}
	return res
}

func Erosion(x, y int) int {
	key := fmt.Sprintf("%v,%v", x, y)
	if val, ok := memo[key]; ok {
		return val
	}
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
		geologic = Erosion(x-1, y) * Erosion(x, y-1)
	}
	erosion := (geologic + depth) % 20183
	memo[key] = erosion
	return memo[key]
}

func Lazy(x, y int) int {
	return Erosion(x, y) % 3
}

// Only if u and v are one move apart.
func (u *State) W(v *State) int {
	if u.x == v.x && u.y == v.y {
		return 7
	} else {
		return 1
	}
}

func Iabs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Imin(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func H(s *State) int {
	res := Iabs(s.x-xt) + Iabs(s.y-yt)
	if s.gear != Torch {
		res += 7
	}
	return res
}

func Shortest() int {
	init := State{x: 0, y: 0, gear: Torch}
	goal := State{x: xt, y: yt, gear: Torch}
	known := make(map[string]bool)
	dist := make(map[string]int)
	score := make(map[string]int)
	known[fmt.Sprintf("%v,%v,%v", init.x, init.y, init.gear)] = true
	dist[fmt.Sprintf("%v,%v,%v", init.x, init.y, init.gear)] = 0
	score[fmt.Sprintf("%v,%v,%v", init.x, init.y, init.gear)] = H(&init)
	// TODO: Use priority queue to make this run faster.
	open := make([]State, 0)
	for _, s := range init.Moves() {
		s := s
		s_key := fmt.Sprintf("%v,%v,%v", s.x, s.y, s.gear)
		dist[s_key] = s.W(&init)
		score[s_key] = dist[s_key] + H(&s)
		open = append(open, s)
	}
	last := &init
	for !last.Same(&goal) {
		var vnext *State
		snext := -1
		for _, s := range open {
			s := s
			s_key := fmt.Sprintf("%v,%v,%v", s.x, s.y, s.gear)
			if known[s_key] {
				continue
			}
			if snext < 0 || score[s_key] < snext {
				vnext = &s
				snext = score[s_key]
			}
		}
		if vnext == nil {
			panic("Open set is empty")
		}
		known[fmt.Sprintf("%v,%v,%v", vnext.x, vnext.y, vnext.gear)] = true
		last = vnext
		nextopen := make([]State, 0)
		for _, s := range open {
			s := s
			if known[fmt.Sprintf("%v,%v,%v", s.x, s.y, s.gear)] {
				continue
			}
			nextopen = append(nextopen, s)
		}
		open = nextopen
		for _, x := range vnext.Moves() {
			x := x
			keyx := fmt.Sprintf("%v,%v,%v", x.x, x.y, x.gear)
			if known[keyx] {
				continue
			}
			newd := dist[fmt.Sprintf("%v,%v,%v", vnext.x, vnext.y, vnext.gear)] + vnext.W(&x)
			if _, ok := dist[keyx]; ok {
				if newd < dist[keyx] {
					dist[keyx] = newd
					score[keyx] = dist[keyx] + H(&x)
					open = append(open, x)
				}
			} else {
				dist[keyx] = newd
				score[keyx] = dist[keyx] + H(&x)
				open = append(open, x)
			}
		}
	}
	return dist[fmt.Sprintf("%v,%v,%v", goal.x, goal.y, goal.gear)]
}

func main() {
	r, err := os.Open("day22/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	line1 := scanner.Text()
	depth, _ = strconv.Atoi(strings.Split(line1, ": ")[1])
	scanner.Scan()
	line2 := scanner.Text()
	target := strings.Split(line2, ": ")[1]
	coord_s := strings.Split(target, ",")
	xt, _ = strconv.Atoi(coord_s[0])
	yt, _ = strconv.Atoi(coord_s[1])
	memo = make(map[string]int)
	if Lazy(xt, yt) != Rocky {
		panic("Something is wrong")
	}
	fmt.Println(Shortest())
}
