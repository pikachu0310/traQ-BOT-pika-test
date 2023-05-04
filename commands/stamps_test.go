package commands

import (
	"reflect"
	"testing"
)

func TestFirst(t *testing.T) {
	type args[T any] struct {
		slice []T
		num   int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		testCase[int]{
			"ecchi",
			args[int]{
				slice: []int{1, 2, 3},
				num:   2,
			},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := First(tt.args.slice, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}
