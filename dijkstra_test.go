package main

import (
	"testing"
)

func TestPath1(t *testing.T) {
	field := [][]int{
		{1, 2, 0},
		{2, 0, 1},
		{9, 1, 0},
	}
	start := Point{0, 0}
	end := Point{2, 1}
	path := dijkstra(field, start, end)
	t.Logf("Calling TestPath1 with field: %v, start: %v, end:%v , result %v\n", field, start, end, path)

	expect := []Point{
		{0, 0},
		{1, 0},
		{2, 0},
		{2, 1},
	}
	if len(path) != 4 {
		t.Errorf("Incorrect len(path). Expect %d, got %d", 4, len(path))
	}
	if path[1] != expect[1] || path[2] != expect[2] || path[3] != expect[3] {
		t.Errorf("Incorrect path. Expect %v, got %d", expect, path)
	}

}

func TestPath2(t *testing.T) {
	field := [][]int{
		{1, 2, 0},
		{4, 1, 1},
		{0, 1, 0},
	}
	start := Point{0, 0}
	end := Point{2, 1}
	path := dijkstra(field, start, end)
	t.Logf("Calling TestPath2 with field: %v, start: %v, end:%v , result %v\n", field, start, end, path)

	expect := []Point{
		{0, 0},
		{0, 1},
		{1, 1},
		{2, 1},
	}
	if len(path) != 4 {
		t.Errorf("Incorrect len(path). Expect %d, got %d", 4, len(path))
	}
	if path[1] != expect[1] || path[2] != expect[2] || path[3] != expect[3] {
		t.Errorf("Incorrect path. Expect %v, got %d", expect, path)
	}

}

func TestPath3(t *testing.T) {
	field := [][]int{
		{1, 5, 1},
		{0, 7, 1},
		{0, 4, 4},
	}
	start := Point{0, 0}
	end := Point{2, 2}
	path := dijkstra(field, start, end)
	t.Logf("Calling TestPath3 with field: %v, start: %v, end:%v , result %v\n", field, start, end, path)

	expect := []Point{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 2},
		{2, 2},
	}
	if len(path) != 5 {
		t.Errorf("Incorrect len(path). Expect %d, got %d", 5, len(path))
	}
	if path[1] != expect[1] || path[2] != expect[2] || path[3] != expect[3] {
		t.Errorf("Incorrect path. Expect %v, got %d", expect, path)
	}

}

func TestErrors1(t *testing.T) {
	field := [][]int{
		{1, 2, 0},
		{0, 0, 0},
		{0, 1, 1},
	}
	start := Point{0, 0}
	end := Point{2, 2}
	path := dijkstra(field, start, end)
	t.Logf("Calling TestErrors1 with field: %v, start: %v, end:%v , result %v\n", field, start, end, path)

	if path != nil {
		t.Errorf("Incorrect path. Expect %v, got %v", nil, path)
	}
	if len(path) != 0 {
		t.Errorf("Incorrect len(path). Expect %d, got %d", 0, len(path))
	}

}
func TestErrors2(t *testing.T) {
	field := [][]int{
		{0, 0},
		{0, 0},
	}
	start := Point{0, 0}
	end := Point{-32, 312}
	path := dijkstra(field, start, end)
	t.Logf("Calling TestErrors2 with field: %v, start: %v, end:%v , result %v\n", field, start, end, path)

	if len(path) != 0 {
		t.Errorf("Incorrect len(path). Expect %d, got %d", 5, len(path))
	}
	if path != nil {
		t.Errorf("Incorrect path. Expect %v, got %d", nil, path)
	}
}
