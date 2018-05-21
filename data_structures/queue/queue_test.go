package queue

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewLinkedListQueue()

	result := fmt.Sprintf("%T", queue)

	if !strings.Contains(result, "queue") {
		t.Errorf("Expected type Queue, but got %s", result)
	}
}

func TestEnqueue(t *testing.T) {
	tests := []struct {
		Data []int
		Len  int
	}{
		{[]int{}, 0},
		{[]int{1, 2, 3, 4, 5}, 5},
	}

	for _, test := range tests {
		queue := NewLinkedListQueue()
		for _, data := range test.Data {
			queue.Enqueue(data)
		}

		if queue.Len() != test.Len {
			t.Errorf("Expected %d items in the queue, but got %d", test.Len, queue.Len())
		}
	}
}

func TestDequeue(t *testing.T) {
	tests := []struct {
		Data []int
		Len  int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
	}

	for _, test := range tests {
		queue := NewLinkedListQueue()
		for _, data := range test.Data {
			queue.Enqueue(data)
		}

		for _, data := range test.Data {
			item, err := queue.Dequeue()
			if err != nil {
				t.Errorf("Unexpected error while dequeuing: %s", err.Error())
			}

			if item.(int) != data {
				t.Errorf("Expected Next item to be %d, but got %d", data, item.(int))
			}
		}
	}
}

func BenchmarkQueue(b *testing.B) {
	queue := NewLinkedListQueue()
	for i := 0; i < 1000000; i++ {
		queue.Enqueue(i)
	}

	for queue.Len() != 0 {
		queue.Dequeue()
	}
}
