package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	r, err := os.Open("day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	input := scanner.Text()
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
	fmt.Println(len(b.String()))
}
