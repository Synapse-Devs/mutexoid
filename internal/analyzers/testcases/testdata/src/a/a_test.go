package a

import "testing"

func TestAdd_OnlyPositive(t *testing.T) { // want "table-driven test \"TestAdd_OnlyPositive\" should contain at least one positive and one negative test case"
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		{
			name:    "simple addition",
			a:       1,
			b:       2,
			want:    3,
			wantErr: false,
		},
		{
			name:    "another success",
			a:       10,
			b:       20,
			want:    30,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd_OnlyNegative(t *testing.T) { // want "table-driven test \"TestAdd_OnlyNegative\" should contain at least one positive and one negative test case"
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		{
			name:    "overflow error 1",
			a:       9223372036854775807,
			b:       1,
			want:    0,
			wantErr: true,
		},
		{
			name:    "overflow error 2",
			a:       9223372036854775807,
			b:       2,
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd_Good(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		{
			name:    "simple addition",
			a:       1,
			b:       2,
			want:    3,
			wantErr: false,
		},
		{
			name:    "overflow error",
			a:       9223372036854775807,
			b:       1,
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
