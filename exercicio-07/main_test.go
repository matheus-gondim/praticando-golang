package main

import (
	"reflect"
	"testing"
)

func Test_calculateMedia(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateMedia(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateMedia() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculateMedia() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseData(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseData(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseData() = %v, want %v", got, tt.want)
			}
		})
	}
}
