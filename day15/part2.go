package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	Wall   = -2
	Empty  = -1
	Elf    = 0
	Goblin = 1
	Attack = 3
	HP     = 200
)

type Unit struct {
	side  int
	hp    int
	moved bool
}

func Iabs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Outcome(caves [][]Unit, power []int) (outcome int, elf_die bool) {
	round := 0
	elf_die = false
	for {
		end := true
		for y := range caves {
			for x := range caves[y] {
				if caves[y][x].side < 0 || caves[y][x].moved {
					continue
				}
				// in range
				es := 1 - caves[y][x].side
				n_x, n_y := x, y
				in_range := make([]string, 0)
				adj := false
				for yy := 1; yy < len(caves)-1; yy++ {
					for xx := 1; xx < len(caves[yy])-1; xx++ {
						if caves[yy-1][xx].side == es || caves[yy+1][xx].side == es || caves[yy][xx-1].side == es || caves[yy][xx+1].side == es {
							if caves[yy][xx].side == Empty {
								in_range = append(in_range, fmt.Sprintf("%v,%v", xx, yy))
							} else if xx == x && yy == y {
								adj = true
								break
							}
						}
					}
					if adj {
						end = false
						break
					}
				}
				if !adj && len(in_range) > 0 {
					// shortest paths
					queue := list.New()
					u := fmt.Sprintf("%v,%v", x, y)
					queue.PushBack(u)
					dist := make(map[string]int)
					seen := make(map[string]int)
					pi := make(map[string]string)
					dist[u] = 0
					seen[u] = 1
					for queue.Len() > 0 {
						v := queue.Remove(queue.Front()).(string)
						coords := strings.Split(v, ",")
						xx, _ := strconv.Atoi(coords[0])
						yy, _ := strconv.Atoi(coords[1])
						d := dist[v] + 1
						var cands = []string{fmt.Sprintf("%v,%v", xx, yy-1), fmt.Sprintf("%v,%v", xx-1, yy), fmt.Sprintf("%v,%v", xx+1, yy), fmt.Sprintf("%v,%v", xx, yy+1)}
						for _, cand := range cands {
							coords := strings.Split(cand, ",")
							xx, _ := strconv.Atoi(coords[0])
							yy, _ := strconv.Atoi(coords[1])
							if _, prs := seen[cand]; prs {
								continue
							}
							if caves[yy][xx].side == Empty {
								dist[cand] = d
								seen[cand] = 1
								pi[cand] = v
								queue.PushBack(cand)
							}
						}
					}
					// find move
					best := math.MaxInt
					for _, cand := range in_range {
						if d, prs := dist[cand]; prs {
							if d < best {
								best = d
								dd := d
								v := cand
								for dd > 1 {
									v = pi[v]
									dd = dist[v]
								}
								coords := strings.Split(v, ",")
								n_x, _ = strconv.Atoi(coords[0])
								n_y, _ = strconv.Atoi(coords[1])
							}
						}
					}
					if best < math.MaxInt {
						end = false
						caves[n_y][n_x].side = caves[y][x].side
						caves[n_y][n_x].hp = caves[y][x].hp
						caves[n_y][n_x].moved = true
						caves[y][x].side = Empty
					}
				}
				// check & attack
				a_x, a_y, a_hp := -1, -1, HP+1
				for yy := range caves {
					for xx := range caves[yy] {
						if caves[yy][xx].side != es {
							continue
						}
						if Iabs(n_x-xx)+Iabs(n_y-yy) == 1 && caves[yy][xx].hp < a_hp {
							a_x = xx
							a_y = yy
							a_hp = caves[yy][xx].hp
						}
					}
				}
				if a_hp <= HP {
					caves[a_y][a_x].hp -= power[caves[n_y][n_x].side]
					if caves[a_y][a_x].hp <= 0 {
						if caves[a_y][a_x].side == Elf {
							elf_die = true
						}
						caves[a_y][a_x].side = Empty
					}
				}
			}
		}
		if end {
			break
		}
		round += 1
		for y := range caves {
			for x := range caves[y] {
				caves[y][x].moved = false
			}
		}
	}
	round -= 1 // why?
	left := 0
	for y := range caves {
		for x := range caves[y] {
			if caves[y][x].side >= Elf {
				left += caves[y][x].hp
			}
		}
	}
	outcome = round * left
	return
}

func main() {
	r, err := os.Open("day15/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var caves [][]Unit
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		row := make([]Unit, 0)
		for _, c := range line {
			switch c {
			case 'E':
				row = append(row, Unit{side: Elf, hp: HP, moved: false})
			case 'G':
				row = append(row, Unit{side: Goblin, hp: HP, moved: false})
			case '.':
				row = append(row, Unit{side: Empty})
			case '#':
				row = append(row, Unit{side: Wall})
			default:
				panic("Something is wrong!")
			}
		}
		caves = append(caves, row)
	}
	var power = []int{3, 3}
	lo := 4
	hi := HP
	for lo < hi {
		mid := (lo + hi) / 2
		power[Elf] = mid
		var temp [][]Unit
		for y := range caves {
			temp = append(temp, make([]Unit, len(caves[y])))
			copy(temp[y], caves[y])
		}
		_, elf_die := Outcome(temp, power)
		if !elf_die {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	power[Elf] = lo
	o, _ := Outcome(caves, power)
	fmt.Println(o)
}
