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

type Bot struct {
	Pos  [3]int
	R    int
	DMin int
	DMax int
}

func Iabs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func (b *Bot) HasInRange(o *Bot) bool {
	dist := 0
	for i := 0; i < 3; i++ {
		dist += Iabs(b.Pos[i] - o.Pos[i])
	}
	return dist <= b.R
}

func main() {
	r, err := os.Open("day23/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`^pos=<(-?\d+),(-?\d+),(-?\d+)>, r=(-?\d+)$`)

	var bots []Bot
	scanner := bufio.NewScanner(r)
	var cands []int
	for scanner.Scan() {
		var bot Bot
		matches := re.FindStringSubmatch(scanner.Text())
		for i := 0; i <= 2; i++ {
			bot.Pos[i], _ = strconv.Atoi(string(matches[i+1]))
		}
		bot.R, _ = strconv.Atoi(string(matches[4]))
		bot.DMin = Iabs(bot.Pos[0]) + Iabs(bot.Pos[1]) + Iabs(bot.Pos[2]) - bot.R
		if bot.DMin < 0 {
			bot.DMin = 0
		}
		bot.DMax = Iabs(bot.Pos[0]) + Iabs(bot.Pos[1]) + Iabs(bot.Pos[2]) + bot.R
		cands = append(cands, bot.DMin)
		bots = append(bots, bot)
	}
	sort.Ints(cands)
	best_d := 0
	best_n := 0
	for _, cand := range cands {
		n := 0
		for _, bot := range bots {
			if bot.DMin <= cand && cand <= bot.DMax {
				n++
			}
		}
		if n > best_n {
			best_n = n
			best_d = cand
		}
	}
	fmt.Println(best_d)
}
