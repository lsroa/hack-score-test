package math

import (
	"strings"
)

type Number = []string

func Equal(a Number, b Number) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Increase(n Number, base []string) Number {
	top := base[len(base)-1]
	bottom := base[0]
	carry := 1

	for i := len(n) - 1; i >= 0; i-- {
		value := n[i]
		pos := GetPos(base, value)

		if value != top {
			n[i] = base[pos+carry]
			carry = 0
		}

		if value == top && carry == 1 {
			n[i] = bottom
			if i == 0 {
				n = append([]string{base[1]}, n...)
			}
		}
	}

	return n
}

func GetPos(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func FindValue(base string, input string) int {
	base_slice := strings.Split(base, "")
	input_slice := strings.Split(input, "")
	current := []string{base_slice[0]}
	count := 0

	for !Equal(current, input_slice) {
		current = Increase(current, base_slice)
		count += 1
	}

	return count

}
