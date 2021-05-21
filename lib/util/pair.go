package util

import "sort"

type Pair struct {
	Name  string
	Score float32
}

type PairList []Pair

//Len()
func (l PairList) Len() int {
	return len(l)
}

//Less(): 成绩将有低到高排序
func (l PairList) Less(i, j int) bool {
	return l[i].Score > l[j].Score
}

//Swap()
func (l PairList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}


//100-->--->1
func (l *PairList) Sort() {
	sort.Sort(l)
}

//1-->100
func (l *PairList) RSort() {
	sort.Sort(sort.Reverse(l))
}

