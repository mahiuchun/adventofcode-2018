package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	A = 1
	B = 2
	C = 3
)

type Op func(code [4]int, regs []int)

var ins_set = map[string]Op{
	"addr": Op(func(code [4]int, regs []int) {
		regs[code[C]] = regs[code[A]] + regs[code[B]]
	}),
	"addi": Op(func(code [4]int, regs []int) {
		regs[code[C]] = regs[code[A]] + code[B]
	}),
	"mulr": Op(func(code [4]int, regs []int) {
		regs[code[C]] = regs[code[A]] * regs[code[B]]
	}),
	"muli": Op(func(code [4]int, regs []int) {
		regs[code[C]] = regs[code[A]] * code[B]
	}),
	"banr": Op(func(code [4]int, regs []int) {
		regs[code[C]] = regs[code[A]] & regs[code[B]]
	}),
	"bani": Op(func(code [4]int, regs []int) {
		regs[code[C]] = regs[code[A]] & code[B]
	}),
	"borr": Op(func(code [4]int, regs []int) {
		regs[code[C]] = regs[code[A]] | regs[code[B]]
	}),
	"bori": Op(func(code [4]int, regs []int) {
		regs[code[C]] = regs[code[A]] | code[B]
	}),
	"setr": Op(func(code [4]int, regs []int) {
		regs[code[C]] = regs[code[A]]
	}),
	"seti": Op(func(code [4]int, regs []int) {
		regs[code[C]] = code[A]
	}),
	"gtir": Op(func(code [4]int, regs []int) {
		if code[A] > regs[code[B]] {
			regs[code[C]] = 1
		} else {
			regs[code[C]] = 0
		}
	}),
	"gtri": Op(func(code [4]int, regs []int) {
		if regs[code[A]] > code[B] {
			regs[code[C]] = 1
		} else {
			regs[code[C]] = 0
		}
	}),
	"gtrr": Op(func(code [4]int, regs []int) {
		if regs[code[A]] > regs[code[B]] {
			regs[code[C]] = 1
		} else {
			regs[code[C]] = 0
		}
	}),
	"eqir": Op(func(code [4]int, regs []int) {
		if code[A] == regs[code[B]] {
			regs[code[C]] = 1
		} else {
			regs[code[C]] = 0
		}
	}),
	"eqri": Op(func(code [4]int, regs []int) {
		if regs[code[A]] == code[B] {
			regs[code[C]] = 1
		} else {
			regs[code[C]] = 0
		}
	}),
	"eqrr": Op(func(code [4]int, regs []int) {
		if regs[code[A]] == regs[code[B]] {
			regs[code[C]] = 1
		} else {
			regs[code[C]] = 0
		}
	})}

type Line struct {
	ins  string
	code [4]int
}

func main() {
	r, err := os.Open("day21/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	ipr, _ := strconv.Atoi(scanner.Text()[4:])
	lines := make([]Line, 0)
	regs := make([]int, 6)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		var line = Line{ins: tokens[0]}
		for i := 1; i <= 3; i++ {
			line.code[i], _ = strconv.Atoi(tokens[i])
		}
		lines = append(lines, line)
	}
	ip := 0
	for ip >= 0 && ip < len(lines) {
		line := lines[ip]
		ins_set[line.ins](line.code, regs)
		regs[ipr]++
		ip = regs[ipr]
		// NOTE: this is input-specific
		if ip == len(lines)-2 {
			fmt.Println(regs[4])
			break
		}
	}
}
