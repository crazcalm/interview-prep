package graph

import (
	"testing"
)

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
