package main

import (
	"container/heap"
	"testing"
)

func TestPointQueue(t *testing.T) {
	testTable := []struct {
		items    []Item
		expected Item
	}{
		{
			[]Item{
				{point: Point{col: 1, row: 1}, priority: 1},
				{point: Point{col: 2, row: 2}, priority: 0},
				{point: Point{col: 3, row: 3}, priority: 1},
			},

			Item{point: Point{col: 2, row: 2}, priority: 0},
		},
		{
			[]Item{
				{point: Point{col: 1, row: 1}, priority: 0},
				{point: Point{col: 2, row: 2}, priority: 2},
				{point: Point{col: 3, row: 3}, priority: 43},
			},

			Item{point: Point{col: 1, row: 1}, priority: 0},
		},
		{
			[]Item{
				{point: Point{col: 2, row: 2}, priority: 23},
				{point: Point{col: 3, row: 3}, priority: 3422},
				{point: Point{col: 1, row: 1}, priority: 1},
			},

			Item{point: Point{col: 1, row: 1}, priority: 1},
		},
	}
	p := NewPointQueue()
	for _, testCase := range testTable {
		for _, item := range testCase.items {
			heap.Push(p, item)
		}

		t.Logf("Calling test with items: (%v), result %d\n", testCase.items, testCase.expected)

		result := heap.Pop(p).(Item)
		if result.point.col != testCase.expected.point.col {
			t.Errorf("Incorrect result. Expect %d, got %d", testCase.expected, result)
		}
	}
}

func TestPointQueue_IsEmpty(t *testing.T) {
	pq := NewPointQueue()
	if isEmpty := pq.IsEmpty(); !isEmpty {
		t.Errorf("Incorrect result. Expect %v, got %v", true, isEmpty)
	}
}
