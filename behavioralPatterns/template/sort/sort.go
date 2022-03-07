package main

import (
	"fmt"
	"sort"
)

type MyList []int

func (m MyList) Len() int {
	return len(m)
}

func (m MyList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	list := MyList{4, 3, 2, 1}
	fmt.Println(list)
	sort.Sort(list)
	fmt.Println(list)
}
