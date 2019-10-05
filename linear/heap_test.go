package linear

import "testing"

func TestHeaps(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Tets#1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Heaps()
		})
	}
}
