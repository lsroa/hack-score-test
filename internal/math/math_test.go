package math

import (
	"testing"
)

type TestCase struct {
	name  string
	input Number
	want  Number
}

func TestIncrease(t *testing.T) {
	base := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	tests := []TestCase{
		{"Base case", Number{"1"}, Number{"2"}},
		{"First digit is top", Number{"9", "2", "9"}, Number{"9", "3", "0"}},
		{"All digits are top", Number{"9", "9", "9"}, Number{"1", "0", "0", "0"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Increase(test.input, base)
			if !Equal(got, test.want) {
				t.Errorf("Increase(%v) = %v, want %v", test.input, got, test.want)
			}
		})
	}

	t.Run("Binary base", func(t *testing.T) {
		got := Increase(Number{"0", "1", "1", "1"}, Number{"0", "1"})
		want := Number{"1", "0", "0", "0"}

		if !Equal(want, got) {
			t.Errorf("\n %v got \n %v want", got, want)
		}
	})
}

func TestFindValue(t *testing.T) {
	got := FindValue("01", "01100")
	want := 12

	if got != want {
		t.Errorf("\n %v got \n %v want", got, want)
	}
}
