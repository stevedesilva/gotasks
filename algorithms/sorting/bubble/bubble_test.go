package bubble_test

import (
	"reflect"
	"testing"

	"github.com/stevedesilva/learngo/algorithms/sorting/bubble"
)

func TestBubbleSort_should_sort_sequence(t *testing.T) {
	type test struct {
		input []int
		want  []int
	}

	tests := []test{
		// {
		// 	input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		// 	want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		// },
		// {
		// 	input: []int{4, 1, 3, 2, 5, 0},
		// 	want:  []int{0, 1, 2, 3, 4, 5},
		// },
		{
			input: []int{1, 9, 2, 5},
			want:  []int{1, 2, 5, 9},
		},
	}

	for _, tc := range tests {
		actual := bubble.Sort(tc.input)
		if !reflect.DeepEqual(tc.want, actual) {
			t.Errorf("expected: %v, got: %v", tc.want, actual)
		}
	}

}
