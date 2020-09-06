package parallelcomputing

import "testing"

func TestOddEvenMonitor(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OddEvenMonitor()
		})
	}
}
