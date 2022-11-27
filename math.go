package aocutil

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](left, right T) T {
	if left < right {
		return left
	} else {
		return right
	}
}

func Max[T constraints.Ordered](left, right T) T {
	if left > right {
		return left
	} else {
		return right
	}
}

func Clamp[T constraints.Ordered](value, lower, upper T) T {
	if value < lower {
		return lower
	} else if value > upper {
		return upper
	} else {
		return value
	}
}

func ClampLower[T constraints.Ordered](value, lower T) T {
	if value < lower {
		return lower
	} else {
		return value
	}
}

func ClampUpper[T constraints.Ordered](value, upper T) T {
	if value > upper {
		return upper
	} else {
		return value
	}
}
