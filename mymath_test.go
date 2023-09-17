package mymath

import (
    "math"
    "testing"
    "studentgit.kata.academy/OvchinnikovVlI/go-kata/course1/12.package/5.package_version/task1.12.5.1/mymath" 
)


func TestAbs(t *testing.T) {
	tests := []struct {
		input  float64
		output float64
	}{
		{-1, 1},
		{0, 0},
		{1, 1},
	}

	for _, tt := range tests {
		got := Abs(tt.input)
		if got != tt.output {
			t.Errorf("Abs(%v) = %v; want %v", tt.input, got, tt.output)
		}
	}
}

func TestAcos(t *testing.T) {
	tests := []struct {
		input  float64
		output float64
	}{
		{1, 0},
		{0, math.Pi / 2},
		{-1, math.Pi},
	}

	for _, tt := range tests {
		got := Acos(tt.input)
		if got != tt.output {
			t.Errorf("Acos(%v) = %v; want %v", tt.input, got, tt.output)
		}
	}
}

func TestAcosh(t *testing.T) {
	tests := []struct {
		input  float64
		output float64
	}{
		{1, 0},
		{math.Cosh(1), 1},
	}

	for _, tt := range tests {
		got := Acosh(tt.input)
		if got != tt.output {
			t.Errorf("Acosh(%v) = %v; want %v", tt.input, got, tt.output)
		}
	}
}

func TestAsin(t *testing.T) {
	tests := []struct {
		input  float64
		output float64
	}{
		{0, 0},
		{1, math.Pi / 2},
		{-1, -math.Pi / 2},
	}

	for _, tt := range tests {
		got := Asin(tt.input)
		if got != tt.output {
			t.Errorf("Asin(%v) = %v; want %v", tt.input, got, tt.output)
		}
	}
}

func TestAsinh(t *testing.T) {
	tests := []struct {
		input  float64
		output float64
	}{
		{0, 0},
		{math.Sinh(1), 1},
		{-math.Sinh(1), -1},
	}

	for _, tt := range tests {
		got := Asinh(tt.input)
		if got != tt.output {
			t.Errorf("Asinh(%v) = %v; want %v", tt.input, got, tt.output)
		}
	}
}
