package p

import "testing"

func TestAddBad(t *testing.T) { // want "test function TestAddBad should call t.Parallel()" "test function TestAddBad should use table-driven tests"
	if got := Add(1, 2); got != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", got)
	}
}

func TestAddGood(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 1, 2, 3},
		{"negative", -1, -2, -3},
		{"zero", 0, 0, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Add(tt.a, tt.b); got != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
