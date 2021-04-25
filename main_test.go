package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T){
	var tests = []struct{
		n string
		want string
	}{
		{"12345", "54321"},
		{"0", "0"},
		{"111", "111"},
		{"", ""},
		{"-1", "1-"},
	}

	for _, tt := range tests{
		testName := fmt.Sprintf("%s", tt.n)
		t.Run(testName, func(t *testing.T) {
			ans := Reverse(tt.n)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}