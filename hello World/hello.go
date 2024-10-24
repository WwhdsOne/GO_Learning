package main

type node struct {
	id, st, re, pr int64
}

type PriorityQueue []node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].pr == pq[j].pr {
		return pq[i].st > pq[j].st
	}
	return pq[i].pr < pq[j].pr
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func calc(x, y int) int {
	return x + y
}

func main() {
	//x := 10
	var x int

}
