package main

import (
	"testing"
)

func Test_isProth(t *testing.T) {
	type args struct {
		in int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"positive test",
			args{
				41,
			},
			true,
		},
		{
			"negative test",
			args{
				4,
			},
			false,
		},
		{
			"negative test",
			args{
				25,
			},
			false,
		},
		{
			"negative test",
			args{
				73,
			},
			false,
		},
		{
			"negative test",
			args{
				9,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isProth(tt.args.in); got != tt.want {
				t.Errorf("isProth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"True test",
			args{
				4,
			},
			true,
		},
		{
			"False test",
			args{
				5,
			},
			false,
		},
		{
			"False test",
			args{
				0,
			},
			false,
		},
		{
			"True test",
			args{
				1024,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"negative test",
			args{
				25,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.n); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isProthPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"positive test",
			args{
				41,
			},
			true,
		},
		{
			"negative test",
			args{
				4,
			},
			false,
		},
		{
			"negative test",
			args{
				25,
			},
			false,
		},
		{
			"negative test",
			args{
				73,
			},
			false,
		},
		{
			"negative test",
			args{
				9,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isProthPrime(tt.args.n)

			if got != tt.want {
				t.Errorf("isProthPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
