package publiclibrarytemplate_test

import (
	"testing"

	publiclibrarytemplate "github.com/PaddleHQ/go-public-library-template"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive numbers", a: 1, b: 2, want: 3},
		{name: "negative numbers", a: -1, b: -2, want: -3},
		{name: "zero", a: 0, b: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := publiclibrarytemplate.Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
