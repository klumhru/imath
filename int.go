package imath

import "errors"

func Min(lt, rt int) int {
	if lt <= rt {
		return lt
	}
	return rt
}

func Max(lt, rt int) int {
	if lt >= rt {
		return lt
	}
	return rt
}

func Clamp(v, lt, rt int) (int, error) {
	if !(lt <= rt) {
		return 0, errors.New("lt must be less or eq to rt")
	}
	if v < lt {
		return lt, nil
	}
	if v > rt {
		return rt, nil
	}
	return v, nil
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
