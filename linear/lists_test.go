package linear

import "testing"

func TestList(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			List()
		})
	}
}
