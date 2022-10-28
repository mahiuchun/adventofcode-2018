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

type Sample struct {
	Before [4]int
	Inst   [4]int
	After  [4]int
}

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

func main() {
	r, err := os.Open("day16/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	n_blank := 0
	scanner := bufio.NewScanner(r)
	var before, code, after [4]int
	var samples []Sample
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			n_blank += 1
			if n_blank > 1 {
				break
			}
			samples = append(samples, Sample{Before: before, Inst: code, After: after})
		} else {
			n_blank = 0
			switch {
			case strings.HasPrefix(line, "Before: "):
				nss := strings.Split(line[9:len(line)-1], ", ")
				for i := 0; i < 4; i++ {
					before[i], _ = strconv.Atoi(nss[i])
				}
			case strings.HasPrefix(line, "After: "):
				nss := strings.Split(line[9:len(line)-1], ", ")
				for i := 0; i < 4; i++ {
					after[i], _ = strconv.Atoi(nss[i])
				}
			default:
				nss := strings.Split(line, " ")
				for i := 0; i < 4; i++ {
					code[i], _ = strconv.Atoi(nss[i])
				}
			}
		}
	}

	tot := 0
	for _, sample := range samples {
		count := 0
		for _, fn := range ins_set {
			temp := make([]int, 4)
			for i := 0; i < 4; i++ {
				temp[i] = sample.Before[i]
			}
			fn(sample.Inst, temp)
			same := true
			for i := 0; i < 4; i++ {
				if temp[i] != sample.After[i] {
					same = false
				}
			}
			if same {
				count += 1
			}
		}
		if count >= 3 {
			tot++
		}
	}
	fmt.Println(tot)
}
