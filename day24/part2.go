package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Group struct {
	army   string
	no     int
	units  int
	hp     int
	weak   []string
	immune []string
	attack int
	typ    string
	init   int
	target *Group
}

func (grp *Group) String() string {
	return fmt.Sprintf("%v Group %v", grp.army, grp.no)
}

type ByEp []*Group

func (a ByEp) Len() int      { return len(a) }
func (a ByEp) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByEp) Less(i, j int) bool {
	ep_i := a[i].units * a[i].attack
	ep_j := a[j].units * a[j].attack
	if ep_i > ep_j {
		return true
	} else if ep_i < ep_j {
		return false
	}
	return a[i].init > a[j].init
}

type ByInit []*Group

func (a ByInit) Len() int      { return len(a) }
func (a ByInit) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByInit) Less(i, j int) bool {
	return a[i].init > a[j].init
}

type Army struct {
	name   string
	groups []*Group
}

func (a *Army) Left() int {
	tot := 0
	for _, grp := range a.groups {
		tot += grp.units
	}
	return tot
}

func (a *Army) Copy() *Army {
	res := &Army{name: a.name}
	for _, grp := range a.groups {
		grp := grp
		ngrp := *grp
		res.groups = append(res.groups, &ngrp)
	}
	return res
}

func ReadWeakImmune(s string, group *Group) {
	if s == "" {
		return
	}
	parts := strings.Split(s, "; ")
	for _, part := range parts {
		words := strings.Split(part, " ")
		if words[0] == "weak" {
			for _, typ := range words[2:] {
				group.weak = append(group.weak, strings.TrimRight(typ, ","))

			}
		} else if words[0] == "immune" {
			for _, typ := range words[2:] {
				group.immune = append(group.immune, strings.TrimRight(typ, ","))

			}
		} else {
			panic("something is wrong")
		}
	}
}

func ReadArmy(scanner *bufio.Scanner) (army Army) {
	scanner.Scan()
	army.name = scanner.Text()
	army.name = army.name[:len(army.name)-1]
	group_re := regexp.MustCompile(`(?P<units>\d+) units each with (?P<hp>\d+) hit points(?: \((?P<weak_immune>[a-z ,;]+)\))? with an attack that does (?P<attack>\d+) (?P<type>\w+) damage at initiative (?P<init>\d+)`)
	no := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		m := group_re.FindStringSubmatch(line)
		group := Group{army: army.name, no: no}
		group.units, _ = strconv.Atoi(m[group_re.SubexpIndex("units")])
		group.hp, _ = strconv.Atoi(m[group_re.SubexpIndex("hp")])
		group.attack, _ = strconv.Atoi(m[group_re.SubexpIndex("attack")])
		group.typ = m[group_re.SubexpIndex("type")]
		group.init, _ = strconv.Atoi(m[group_re.SubexpIndex("init")])
		ReadWeakImmune(m[group_re.SubexpIndex("weak_immune")], &group)
		army.groups = append(army.groups, &group)
		no++
	}
	return
}

func Clean(a *Army) {
	n := 0
	for i := range a.groups {
		if a.groups[i].units > 0 {
			a.groups[n] = a.groups[i]
			n++
		}
	}
	a.groups = a.groups[:n]
}

func Damage(attacker *Group, defender *Group) int {
	res := attacker.units * attacker.attack
	for _, typ := range defender.weak {
		if attacker.typ == typ {
			res *= 2
		}
	}
	for _, typ := range defender.immune {
		if attacker.typ == typ {
			res *= 0
		}
	}
	return res
}

func Fight(a1 *Army, a2 *Army) (progress bool) {
	var groups []*Group
	for _, grp := range a1.groups {
		grp := grp
		groups = append(groups, grp)
	}
	for _, grp := range a2.groups {
		grp := grp
		groups = append(groups, grp)
	}
	// target selection
	sort.Sort(ByEp(groups))
	chosen := make(map[string]int)
	for _, grp := range groups {
		grp := grp
		if grp.units == 0 {
			panic("something is wrong")
		}
		var enemy *Army
		if grp.army == a1.name {
			enemy = a2
		} else {
			enemy = a1
		}
		grp.target = nil
		b_damage := 0
		b_ep := 0
		b_init := 0
		for i := range enemy.groups {
			if _, prs := chosen[enemy.groups[i].String()]; prs {
				continue
			}
			damage := Damage(grp, enemy.groups[i])
			if damage > b_damage {
				grp.target = enemy.groups[i]
				b_damage = Damage(grp, grp.target)
				b_ep = grp.target.units * grp.target.attack
				b_init = grp.target.init
			} else if damage < b_damage {
				continue
			}
			ep := enemy.groups[i].units * enemy.groups[i].attack
			if ep > b_ep {
				grp.target = enemy.groups[i]
				b_damage = Damage(grp, grp.target)
				b_ep = grp.target.units * grp.target.attack
				b_init = grp.target.init
			} else if ep < b_ep {
				continue
			}
			init := enemy.groups[i].init
			if init > b_init {
				grp.target = enemy.groups[i]
				b_damage = Damage(grp, grp.target)
				b_ep = grp.target.units * grp.target.attack
				b_init = grp.target.init
			}
		}
		if b_damage == 0 {
			grp.target = nil
		}
		if grp.target != nil {
			chosen[grp.target.String()] = 1
		}
	}
	// attacking
	sort.Sort(ByInit(groups))
	for _, grp := range groups {
		grp := grp
		if grp.units == 0 || grp.target == nil {
			continue
		}
		target := grp.target
		damage := Damage(grp, target)
		kill := damage / target.hp
		if kill > target.units {
			kill = target.units
		}
		if kill > 0 {
			progress = true
		}
		target.units -= kill
	}
	Clean(a1)
	Clean(a2)
	return
}

func Battle(a1 *Army, a2 *Army, boost int) bool {
	// Assume a1 is immune system
	for _, grp := range a1.groups {
		grp := grp
		grp.attack += boost
	}
	for len(a1.groups) > 0 && len(a2.groups) > 0 {
		if !Fight(a1, a2) {
			break
		}
	}
	return len(a1.groups) > 0 && len(a2.groups) == 0
}

func main() {
	r, err := os.Open("day24/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(r)
	a1 := ReadArmy(scanner)
	a2 := ReadArmy(scanner)
	lo := 0
	hi := 987654321
	for lo < hi {
		mid := (lo + hi) / 2
		c1, c2 := a1.Copy(), a2.Copy()
		if Battle(c1, c2, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	Battle(&a1, &a2, lo)
	fmt.Println(a1.Left())
}
