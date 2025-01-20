package main

import (
	"container/heap"
	"fmt"
)

type Point struct {
	row, col int
}

func (p Point) String() string {
	return fmt.Sprintf("%d %d", p.row, p.col)
}

type Item struct {
	point    Point
	priority int
	index    int
}

type PointQueue []Item

func NewPointQueue() *PointQueue {
	var pq PointQueue
	heap.Init(&pq)
	return &pq
}

// оставил обычные получатели и с указателем, так как разработчики
// делают так в обновляемой документации:
// https://pkg.go.dev/container/heap#example__priorityQueue
func (pq PointQueue) Len() int {
	return len(pq)
}

func (pq PointQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PointQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PointQueue) Push(x any) {
	item := x.(Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PointQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
func (pq *PointQueue) IsEmpty() bool {
	return len(*pq) == 0
}
