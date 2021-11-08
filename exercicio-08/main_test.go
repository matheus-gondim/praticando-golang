package main

import (
	"reflect"
	"testing"
)

func TestSplitString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Test with an even sized string",
			args{"awsaws"},
			[]string{"aw", "sa", "ws"},
		},
		{
			"Test with an odd sized string",
			args{"abc"},
			[]string{"ab", "c_"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitString(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitString() = %v, want %v", got, tt.want)
			}
		})
	}
}
