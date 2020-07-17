package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   string
	Value int
}
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func main() {
	testMap := make(map[string]int)
	testMap["a"] = 1
	testMap["d"] = 4
	testMap["e"] = 5
	testMap["b"] = 2
	testMap["c"] = 3

	pairList := sortMapByValue(testMap)
	fmt.Printf("pairList : %v", pairList)
}
