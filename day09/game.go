package game

import "container/list"

func cw(l *list.List, e *list.Element, n int) *list.Element {
	p := e
	for i := 0; i < n; i++ {
		p = p.Next()
		if p == nil {
			p = l.Front()
		}
	}
	return p
}

func ccw(l *list.List, e *list.Element, n int) *list.Element {
	p := e
	for i := 0; i < n; i++ {
		p = p.Prev()
		if p == nil {
			p = l.Back()
		}
	}
	return p
}

func HighScore(np, nm int) int {
	circle := list.New()
	circle.PushBack(0)
	curr := circle.Front()
	player := 0
	tally := make(map[int]int)
	for i := 1; i <= nm; i++ {
		if i%23 != 0 {
			curr = circle.InsertAfter(i, cw(circle, curr, 1))
		} else {
			tally[player] += i
			to_remove := ccw(circle, curr, 7)
			curr = cw(circle, to_remove, 1)
			val := circle.Remove(to_remove).(int)
			tally[player] += val
		}
		player += 1
		player %= np
	}
	best := 0
	for _, score := range tally {
		if score > best {
			best = score
		}
	}
	return best
}
