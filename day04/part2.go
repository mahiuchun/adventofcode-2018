package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Record struct {
	last   int
	asleep bool
	data   [60]bool
	tot    int
}

func main() {
	r, err := os.Open("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := make([]string, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sort.Strings(lines)

	re := regexp.MustCompile(`^\[1518-(\d\d)-(\d\d) (?:\d\d):(\d\d)\] (.+)$`)
	re2 := regexp.MustCompile(`^Guard #(\d+) begins shift$`)

	guards := make(map[string]int)
	asleep := make(map[string]*Record)

	curr_num := 0
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		month, _ := strconv.Atoi(matches[1])
		day, _ := strconv.Atoi(matches[2])
		minute, _ := strconv.Atoi(matches[3])
		msg := matches[4]
		key := fmt.Sprintf("%v-%v", month, day)
		if _, ok := asleep[key]; !ok {
			asleep[key] = &Record{last: -1}
		}
		switch msg {
		case "falls asleep":
			guards[key] = curr_num
			asleep[key].asleep = true
			asleep[key].last = minute
		case "wakes up":
			guards[key] = curr_num
			for i := asleep[key].last; i < minute; i++ {
				asleep[key].data[i] = true
				asleep[key].tot += minute - asleep[key].last
			}
			asleep[key].asleep = false
			asleep[key].last = minute
		default:
			matches := re2.FindStringSubmatch(msg)
			curr_num, _ = strconv.Atoi(matches[1])
		}
	}
	for key, rec := range asleep {
		if rec.last < 0 {
			delete(asleep, key)
		}
		if rec.asleep {
			panic("Asleep at the end of the shift!")
		}
	}

	guards_tally := make(map[int][]int)
	for key, rec := range asleep {
		num := guards[key]
		if guards_tally[num] == nil {
			guards_tally[num] = make([]int, 60)
		}
		for i := 0; i < 60; i++ {
			if rec.data[i] {
				guards_tally[num][i] += 1
			}
		}
	}

	max, maxid, maxmin := 0, -1, -1
	for key, tally := range guards_tally {
		for i := 0; i < 60; i++ {
			if tally[i] > max {
				max = tally[i]
				maxid = key
				maxmin = i
			}
		}
	}
	fmt.Printf("%v * %v = %v\n", maxid, maxmin, maxid*maxmin)
}
