package main

import (
	"container/heap"
	"fmt"
)

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

func main() {
	var c node
	var ti int64
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for {
		_, err := fmt.Scanf("%d%d%d%d", &c.id, &c.st, &c.re, &c.pr)
		if err != nil {
			break
		}

		for pq.Len() > 0 && ti+pq[0].re <= int64(c.st) {
			b := heap.Pop(&pq).(node)
			fmt.Printf("%d %d\n", b.id, ti+int64(b.re))
			ti += int64(b.re)
		}

		if pq.Len() > 0 {
			d := heap.Pop(&pq).(node)
			d.re = d.re - int64(c.st) + ti
			heap.Push(&pq, d)
		}

		heap.Push(&pq, c)
		ti = int64(c.st)
	}

	for pq.Len() > 0 {
		f := heap.Pop(&pq).(node)
		ti += int64(f.re)
		fmt.Printf("%d %d\n", f.id, ti)
	}
}
