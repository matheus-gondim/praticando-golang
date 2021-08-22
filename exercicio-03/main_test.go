package main

import "testing"

func Test_getOlderAge(t *testing.T) {
	type args struct {
		ages []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{1, 2, 3, 10}}, 10},
		{"", args{[]int{1, 20, 3, 10}}, 20},
		{"", args{[]int{112, 2, 3, 10}}, 112},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOlderAge(tt.args.ages); got != tt.want {
				t.Errorf("getOlderAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
