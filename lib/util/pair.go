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

func AlphaSort(arr []Pair, low int, high int) {
	if high <= low {
		return
	}
	j := partition(arr, low, high)
	AlphaSort(arr, low, j-1)
	AlphaSort(arr, j+1, high)
}

var partition = func(arr []Pair, low int, high int) int {
	i, j := low+1, high
	for true {
		for arr[i].Name < arr[low].Name {
			i++
			if i == high {
				break
			}
		}
		for arr[low].Name < arr[j].Name {
			j--
			if j == low {
				break
			}
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[low], arr[j] = arr[j], arr[low]
	return j
}
