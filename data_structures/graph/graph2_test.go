package graph

import (
	"testing"
)

/*
var data = map[int][]int{
		0: []int{1,2,5,6},
		1: []int{0},
		2: []int{0},
		3: []int{4,5},
		4: []int{3,5,6},
		5: []int{0,3,4},
		6: []int{0,4},
		7: []int{8},
		8: []int{7},
		9: []int{10, 11, 12},
		10: []int{9},
		11: []int{9,12},
		12: []int{9,11},
	}
*/

var testData = []byte(`{
	"total_nodes": 12,
	"directed": false,
	"mapping": {
		"0": [1,2,5,6],
		"1": [0],
		"2": [0],
		"3": [4,5],
		"4": [3,5,6],
		"5": [0,3,4],
		"6": [0,4],
		"7": [8],
		"8": [7],
		"9": [10, 11, 12],
		"10": [9],
		"11": [9, 12],
		"12": [9, 11]
	} 
}`)

func TestNew2(t *testing.T) {
	tests := []struct {
		Directed bool
	}{
		{true},
		{false},
	}

	for _, test := range tests {
		New2(test.Directed)
	}
}

func TestReadIn2(t *testing.T) {
	graph := ReadIn(testData)
	Print(graph)
}
