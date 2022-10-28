package main

import (
	"bufio"
	"fmt"
	"game"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	r, err := os.Open("day09/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`^(\d+) players; last marble is worth (\d+) points$`)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	matches := re.FindStringSubmatch(scanner.Text())
	np, _ := strconv.Atoi(matches[1])
	nm, _ := strconv.Atoi(matches[2])
	fmt.Println(game.HighScore(np, nm))
}
