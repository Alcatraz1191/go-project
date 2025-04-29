package main

import "testing"

func TestDo(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Do()
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

