package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Bot struct {
	Pos [3]int
	R   int
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
	var strongest int
	var index int
	for scanner.Scan() {
		var bot Bot
		matches := re.FindStringSubmatch(scanner.Text())
		for i := 0; i <= 2; i++ {
			bot.Pos[i], _ = strconv.Atoi(string(matches[i+1]))
		}
		bot.R, _ = strconv.Atoi(string(matches[4]))
		bots = append(bots, bot)
		if bots[strongest].R < bots[index].R {
			strongest = index
		}
		index += 1
	}
	tot := 0
	for _, bot := range bots {
		if bots[strongest].HasInRange(&bot) {
			tot++
		}
	}
	fmt.Println(tot)
}
