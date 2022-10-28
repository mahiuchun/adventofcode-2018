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

	m := make(map[int]map[string]int)
	for _, sample := range samples {
		cands := make(map[string]int)
		for ins, fn := range ins_set {
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
				cands[ins] = 1
			}
		}
		opcode := sample.Inst[0]
		if _, ok := m[opcode]; ok {
			for k := range m[opcode] {
				if _, ok := cands[k]; !ok {
					delete(m[opcode], k)
				}
			}
		} else {
			m[opcode] = make(map[string]int)
			for cand := range cands {
				m[opcode][cand] = 1
			}
		}
	}
	known := make(map[int]int)
	mm := make(map[int]string)
	for len(known) < len(ins_set) {
		var ins string
		for k, v := range m {
			if len(v) == 1 {
				known[k] = 1
				for kk := range m[k] {
					ins = kk
				}
				mm[k] = ins
				break
			}
		}
		for k := range m {
			if _, ok := known[k]; ok {
				continue
			}
			delete(m[k], ins)
		}
	}

	regs := make([]int, 4)
	scanner.Scan()
	for scanner.Scan() {
		nss := strings.Split(scanner.Text(), " ")
		var inst [4]int
		for i := 0; i < 4; i++ {
			inst[i], _ = strconv.Atoi(nss[i])
		}
		ins_set[mm[inst[0]]](inst, regs)
	}
	fmt.Println(regs[0])
}
