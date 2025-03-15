package main

import "testing"

func TestTagCompatible(t *testing.T) {
	tests := []struct {
		name string
		tag  string
		want bool
	}{
		{
			name: "22 should be compatible",
			tag:  "22",
			want: true,
		},
		{
			name: "22-alpine3.19 should be compatible",
			tag:  "22-alpine3.19",
			want: true,
		},
		{
			name: "22-slim should be compatible",
			tag:  "22-slim",
			want: true,
		},
		{
			name: "22.0 should not be compatible",
			tag:  "22.0",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TagCompatible(tt.tag); got != tt.want {
				t.Errorf("TagCompatible() = %v, want %v", got, tt.want)
			}
		})
	}
}
