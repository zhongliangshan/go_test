package rpcdemo

import (
	"errors"
)

// 定义结构
type QuickRpcDemo struct {
}

type Args struct {
	Inputs []int
}

type Args2 struct {
	A, B int
}

func findPos(arr []int, start, end int) int {
	val := arr[start]
	pos := start
	start++
	for start < end {
		if arr[start] < val {
			arr[start], arr[pos] = arr[pos], arr[start]
			pos++
		}
		start++

	}

	return pos
}

func quickSort2(arr []int, start, end int) []int {
	if start < end {
		pos := findPos(arr, start, end)
		if pos != start || pos == end {
			quickSort2(arr, start, pos)
			quickSort2(arr, pos, end)
		}

	}
	return arr
}

func (QuickRpcDemo) QuickSort(args Args, result *[]int) error {
	if len(args.Inputs) <= 0 {
		return errors.New("empty array")
	}

	*result = quickSort2(args.Inputs, 0, len(args.Inputs))

	return nil
}

func (QuickRpcDemo) Div(args Args2, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)

	return nil
}
