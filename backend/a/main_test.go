package a

import "testing"
import "mypkg/b"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1+2=3", args{1, 2}, 3},
		{"2+3=5", args{2, 3}, 5},
		{"3+4=7", args{3, 4}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestHello(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1+2=3", args{1, 2}, 3},
		{"2+3=5", args{2, 3}, 5},
		{"3+4=7", args{3, 4}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
