package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func diffby1(a string, b string) (bool, string) {
	if len(a) != len(b) {
		panic("IDs have different lengths!")
	}
	count := 0
	var buf []byte
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count += 1
		} else {
			buf = append(buf, a[i])
		}
	}
	return count == 1, string(buf)
}

func main() {
	r, err := os.Open("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var ids []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}
	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			pred, s := diffby1(ids[i], ids[j])
			if pred {
				fmt.Println(s)
			}
		}
	}
}
