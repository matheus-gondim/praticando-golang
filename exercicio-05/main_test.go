package main

import "testing"

func Test_calculaMetragem(t *testing.T) {
	type args struct {
		dto Dto
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test 60 centimeters scratcher", args{Dto{20.0, 60.0, 0.5}}, 75.36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculaMetragem(tt.args.dto); got != tt.want {
				t.Errorf("calculaMetragem() = %v, want %v", got, tt.want)
			}
		})
	}
}
