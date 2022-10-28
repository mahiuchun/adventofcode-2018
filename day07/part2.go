package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

const NumWorkers = 5

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
	task_time := make(map[string]int)
	for key := range indegree {
		tasks = append(tasks, key)
		task_time[key] = 60 + int(key[0]-'A'+1)
	}
	sort.Strings(tasks)

	var worker_task [NumWorkers]string
	var worker_rem [NumWorkers]int
	t := 0
	for {
		free := 0
		for i := 0; i < NumWorkers; i++ {
			if worker_rem[i] > 0 {
				worker_rem[i] -= 1
				if worker_rem[i] == 0 {
					task := worker_task[i]
					for _, v := range after[task] {
						indegree[v] -= 1
					}
				}
			}
		}
		for i := 0; i < NumWorkers; i++ {
			if worker_rem[i] == 0 {
				for _, task := range tasks {
					if d, prs := indegree[task]; prs && d == 0 {
						worker_rem[i] = task_time[task]
						worker_task[i] = task
						delete(indegree, task)
						break
					}
				}
			}
			if worker_rem[i] == 0 {
				free += 1
			}
		}
		if free == NumWorkers {
			break
		}
		t += 1
	}
	fmt.Println(t)
}
