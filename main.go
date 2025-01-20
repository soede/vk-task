package main

import (
	"fmt"
	"os"
)

// Для решения задачи используется алгоритм Дейкстры и приоритетная очередь,
// реализованная с использованием стандартной библиотеки Go.
// Для запуска кода нужно установить Golang до 1.18,
// потому что используется any в реализации очереди (queue.go – 53, 47 строка).
// Но если делать этого не хочется, то просто замените на interface{}
func main() {
	var rows, cols int
	var err error

	_, err = fmt.Scan(&rows)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if rows < 1 {
		fmt.Fprintln(os.Stderr, "Количество строк не может быть меньше 1")
		os.Exit(1)
	}

	_, err = fmt.Scan(&cols)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if cols < 1 {
		fmt.Fprintln(os.Stderr, "Количество столбцов не может быть меньше 1")
		os.Exit(1)
	}

	var filed = make([][]int, rows)
	for i := range filed {
		filed[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			_, err = fmt.Scan(&filed[i][j])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			if filed[i][j] < 0 || filed[i][j] > 9 {
				fmt.Fprintln(os.Stderr, "По условию клетки могут весить от 0 до 9")
				os.Exit(1)
			}
		}
	}

	var sr, sc int
	var er, ec int

	_, err = fmt.Scan(&sr)
	_, err = fmt.Scan(&sc)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, err = fmt.Scan(&er)
	_, err = fmt.Scan(&ec)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	start := Point{sr, sc}
	end := Point{er, ec}

	path := dijkstra(filed, start, end)
	if path == nil {
		fmt.Println("Путь не был найден")
	} else {
		fmt.Printf("\n")
		for _, point := range path {
			fmt.Println(point.String())
		}
		fmt.Println(".")
	}

}
