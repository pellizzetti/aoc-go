package year2015

import (
	"sort"
	"strconv"
)

type Day02 struct{}

func (p Day02) PartA(lines []string) any {
	total := 0
	for i := 0; i < len(lines); i++ {
		dimensions := lines[i]
		if len(dimensions) == 0 {
			continue
		}
		fs := 0
		ss := 0
		for s := 0; s < len(dimensions); s++ {
			c := dimensions[s]
			if c == 'x' {
				if fs == 0 {
					fs = s
				} else {
					ss = s
				}
			}
		}
		l, _ := strconv.Atoi(dimensions[:fs])
		w, _ := strconv.Atoi(dimensions[fs+1 : ss])
		h, _ := strconv.Atoi(dimensions[ss+1:])
		lw := l * w
		wh := w * h
		hl := h * l
		extra := lw
		if wh < extra {
			extra = wh
		}
		if hl < extra {
			extra = hl
		}
		total += (2*lw + 2*wh + 2*hl) + extra
	}
	return total
}

func (p Day02) PartB(lines []string) any {
	total := 0
	for i := 0; i < len(lines); i++ {
		dimensions := lines[i]
		if len(dimensions) == 0 {
			continue
		}
		fs := 0
		ss := 0
		for s := 0; s < len(dimensions); s++ {
			c := dimensions[s]
			if c == 'x' {
				if fs == 0 {
					fs = s
				} else {
					ss = s
				}
			}
		}
		l, _ := strconv.Atoi(dimensions[:fs])
		w, _ := strconv.Atoi(dimensions[fs+1 : ss])
		h, _ := strconv.Atoi(dimensions[ss+1:])
		shortest := []int{l, w, h}
		sort.Ints(shortest)

		total += shortest[0]*2 + shortest[1]*2 + (l * w * h)
	}
	return total
}
