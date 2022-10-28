package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
)

type Pos struct {
	x int
	y int
}

func (pos *Pos) String() string {
	return fmt.Sprintf("%v,%v", pos.x, pos.y)
}

type Graph struct {
	nodes []Pos
	index map[string]int
	edges map[int]map[int]int
}

func (g *Graph) AddNode(n Pos) int {
	key := n.String()
	if idx, prs := g.index[key]; prs {
		return idx
	} else {
		g.index[key] = len(g.nodes)
		g.nodes = append(g.nodes, n)
		return g.index[key]
	}
}

func (g *Graph) AddEdge(u Pos, v Pos) {
	iu := g.AddNode(u)
	iv := g.AddNode(v)
	if iu == iv {
		panic("self-edge is not allowed")
	}
	if _, prs := g.edges[iu]; !prs {
		g.edges[iu] = make(map[int]int)
	}
	g.edges[iu][iv] = 1
	if _, prs := g.edges[iv]; !prs {
		g.edges[iv] = make(map[int]int)
	}
	g.edges[iv][iu] = 1
}

func (g *Graph) Bfs(start Pos) int {
	queue := list.New()
	queue.PushBack(start)
	dist := make(map[string]int)
	seen := make(map[string]int)
	dist[start.String()] = 0
	seen[start.String()] = 1
	tot := 0
	for queue.Len() > 0 {
		u := queue.Remove(queue.Front()).(Pos)
		d := dist[u.String()] + 1
		idx := g.index[u.String()]
		for iv, _ := range g.edges[idx] {
			v := g.nodes[iv]
			if _, prs := seen[v.String()]; prs {
				continue
			}
			queue.PushBack(v)
			dist[v.String()] = d
			seen[v.String()] = 1
			if d >= 1000 {
				tot += 1
			}
		}
	}
	return tot
}

func process(s string, start Pos, g *Graph) (int, []Pos) {
	i := 0
	n := len(s)
	opts := make([]Pos, 0)
	opts = append(opts, start)
	for i < n {
		c := s[i]
		switch c {
		case '(':
			nopts := make([]Pos, 0)
			pos_set := make(map[string]int)
			for {
				i += 1
				ll := 0
				for _, pos := range opts {
					l, lopts := process(s[i:], pos, g)
					ll = l
					for _, npos := range lopts {
						if _, prs := pos_set[npos.String()]; !prs {
							nopts = append(nopts, npos)
							pos_set[npos.String()] = 1
						}

					}
				}
				i += ll
				if s[i] == ')' {
					break
				}
			}
			opts = nopts
		case '|':
			fallthrough
		case ')':
			return i, opts
		case 'E':
			for j := range opts {
				npos := Pos{x: opts[j].x + 1, y: opts[j].y}
				g.AddEdge(opts[j], npos)
				opts[j] = npos
			}
		case 'W':
			for j := range opts {
				npos := Pos{x: opts[j].x - 1, y: opts[j].y}
				g.AddEdge(opts[j], npos)
				opts[j] = npos
			}
		case 'N':
			for j := range opts {
				npos := Pos{x: opts[j].x, y: opts[j].y - 1}
				g.AddEdge(opts[j], npos)
				opts[j] = npos
			}
		case 'S':
			for j := range opts {
				npos := Pos{x: opts[j].x, y: opts[j].y + 1}
				g.AddEdge(opts[j], npos)
				opts[j] = npos
			}
		}
		i += 1
	}
	return i, opts
}

func main() {
	r, err := os.Open("day20/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	my_regex := scanner.Text()
	var g Graph
	g.index = make(map[string]int)
	g.edges = make(map[int]map[int]int)
	orig := Pos{x: 0, y: 0}
	process(my_regex[1:len(my_regex)-1], orig, &g)
	fmt.Println(g.Bfs(orig))
}
