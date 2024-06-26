package utils

import (
	"sort"
)

type StringSlice []string

func (l *StringSlice) Add(list ...string) {
	*l = append(*l, list...)
}

func (l *StringSlice) Delete(i int) {
	*l = append((*l)[:i], (*l)[i+1:]...)
}

func (l StringSlice) Contains(s string) bool {
	for _, e := range l {
		if e == s {
			return true
		}
	}
	return false
}

func (l StringSlice) Sort() {
	sort.Strings(l)
}
