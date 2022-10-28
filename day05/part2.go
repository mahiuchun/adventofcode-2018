package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"unicode"
)

func PolymerLen(input string) int {
	lst := list.New()
	for _, c := range input {
		lst.PushBack(c)
	}

	for p := lst.Front(); p != lst.Back(); {
		curr := p.Value.(rune)
		next := p.Next().Value.(rune)
		if curr != next && unicode.ToLower(curr) == unicode.ToLower(next) {
			var new_p *list.Element
			if p.Prev() != nil {
				new_p = p.Prev()
			} else {
				new_p = p.Next().Next()
			}
			lst.Remove(p.Next())
			lst.Remove(p)
			if new_p == nil {
				break
			}
			p = new_p
		} else {
			p = p.Next()
		}
	}

	b := new(strings.Builder)
	for p := lst.Front(); p != nil; p = p.Next() {
		b.WriteRune(p.Value.(rune))
	}
	return len(b.String())
}

func main() {
	r, err := os.Open("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	input := scanner.Text()

	min := math.MaxInt
	for i := 0; i < 26; i++ {
		remove := rune('a' + i)
		b := new(strings.Builder)
		for _, c := range input {
			if unicode.ToLower(c) != remove {
				b.WriteRune(c)
			}
		}
		cand := PolymerLen(b.String())
		if cand < min {
			min = cand
		}
	}
	fmt.Println(min)
}
