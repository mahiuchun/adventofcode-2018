package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point [4]int

func Iabs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func IsClose(p1, p2 *Point) bool {
	dist := 0
	for i := 0; i < 4; i++ {
		dist += Iabs(p1[i] - p2[i])
	}
	return dist <= 3
}

func Dfs(g [][]int, m map[int]bool, u int) bool {
	if _, prs := m[u]; prs {
		return false
	}
	m[u] = true
	for _, v := range g[u] {
		Dfs(g, m, v)
	}
	return true
}

func main() {
	r, err := os.Open("day25/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var pts []Point
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		var pt Point
		for i, x := range strings.Split(line, ",") {
			pt[i], _ = strconv.Atoi(x)
		}
		pts = append(pts, pt)
	}

	graph := make([][]int, len(pts))
	for i := 0; i < len(pts); i++ {
		graph[i] = make([]int, 0)
	}
	for i := 0; i < len(pts); i++ {
		for j := i + 1; j < len(pts); j++ {
			if IsClose(&pts[i], &pts[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}

	n := 0
	m := make(map[int]bool)
	for i := 0; i < len(pts); i++ {
		if Dfs(graph, m, i) {
			n++
		}
	}
	fmt.Println(n)
}
