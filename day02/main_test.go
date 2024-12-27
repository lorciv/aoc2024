package main

import "testing"

func TestSafe(t *testing.T) {
	tests := []struct {
		record []int
		want   bool
	}{
		{
			record: []int{7, 6, 4, 2, 1},
			want:   true,
		},
		{
			record: []int{1, 2, 7, 8, 9},
			want:   false,
		},
		{
			record: []int{9, 7, 6, 2, 1},
			want:   false,
		},
		{
			record: []int{1, 3, 2, 4, 5},
			want:   false,
		},
		{
			record: []int{8, 6, 4, 4, 1},
			want:   false,
		},
		{
			record: []int{1, 3, 6, 7, 9},
			want:   true,
		},
	}

	for _, s := range tests {
		got := Safe(s.record)
		if got != s.want {
			t.Errorf("Safe(%v) = %v, want %v", s.record, got, s.want)
		}
	}
}
