package main

import "container/heap"

func dijkstra(field [][]int, start, end Point) []Point {
	rows, cols := len(field), len(field[0])

	inBounds := func(r, c int) bool {
		return r >= 0 && r < rows && c >= 0 && c < cols
	}

	if !inBounds(start.row, start.col) || !inBounds(start.row, start.col) {
		return nil
	}

	steps := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	costs := make(map[Point]int)
	previous := make(map[Point]*Point)
	pq := NewPointQueue()

	heap.Push(pq, Item{point: start, priority: 0})
	costs[start] = 0

	for !pq.IsEmpty() {
		current := heap.Pop(pq).(Item)
		point := current.point
		cost := current.priority

		if point == end {
			var path []Point
			for p := &point; p != nil; p = previous[*p] {
				path = append([]Point{*p}, path...)
			}
			return path
		}

		for _, st := range steps {
			near := Point{point.row + st.row, point.col + st.col}

			if inBounds(near.row, near.col) && field[near.row][near.col] != 0 {
				newCost := cost + field[near.row][near.col]

				if d, ok := costs[near]; !ok || newCost < d {
					costs[near] = newCost
					previous[near] = &point
					heap.Push(pq, Item{point: near, priority: newCost})
				}
			}
		}
	}

	return nil
}
