package main

import "testing"

func Test_randInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randInt(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("randInt() = %v, want %v", got, tt.want)
			}
		})
	}
}