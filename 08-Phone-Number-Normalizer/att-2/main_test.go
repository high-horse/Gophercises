package main

import (
	"testing"
)

type phoneBook struct  {
	sample string
	want string
}
func TestNormalize(t *testing.T){
	book := []phoneBook {
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"123-456-7890", "1234567890"},
		{"1234567892", "1234567892"},
		{"(123)456-7892", "1234567892"},
	}
	
	for _, ph := range book {
		t.Run(ph.want, func(t *testing.T) {
			actual := Normalize(ph.sample)
			if actual != ph.want {
				t.Errorf("want %v ; got %v",ph.want, actual)
			}
		})
	}
}