package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	r, err := os.Open("day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	indegree := make(map[string]int)
	after := make(map[string][]string)

	re := regexp.MustCompile(`^Step (.) must be finished before step (.) can begin.$`)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		u := matches[1]
		v := matches[2]
		if _, prs := indegree[u]; !prs {
			indegree[u] = 0
		}
		indegree[v] += 1
		after[u] = append(after[u], v)
	}
	for _, val := range after {
		sort.Strings(val)
	}

	var tasks []string
	for key, _ := range indegree {
		tasks = append(tasks, key)
	}
	sort.Strings(tasks)

	todo := len(tasks)
	var order strings.Builder
	for todo > 0 {
		for _, t := range tasks {
			if d, prs := indegree[t]; prs && d == 0 {
				order.WriteString(t)
				for _, v := range after[t] {
					indegree[v] -= 1
				}
				todo -= 1
				delete(indegree, t)
				break
			}
		}
	}
	fmt.Println(order.String())
}
