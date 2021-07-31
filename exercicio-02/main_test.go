package main

import "testing"

func Test_calculatorOfDistanceBetweenTwoPoints(t *testing.T) {
	type args struct {
		p1 Points
		p2 Points
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"", args{Points{X: 1, Y: 1}, Points{X: 1, Y: 1}}, 0.0},
		{"", args{Points{X: 1.7, Y: 12}, Points{X: 19, Y: 0}}, 21.05445321066306},
		{"", args{Points{X: 1, Y: 2.6}, Points{X: 1.8, Y: 1.2}}, 1.61245154965971},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculatorOfDistanceBetweenTwoPoints(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("calculatorOfDistanceBetweenTwoPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
