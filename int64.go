package mathy

import "errors"

func Min64(lt, rt int64) int64 {
	if lt <= rt {
		return lt
	}
	return rt
}

func Max64(lt, rt int64) int64 {
	if lt >= rt {
		return lt
	}
	return rt
}

func Clamp64(v, lt, rt int64) (int64, error) {
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

func Abs64(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}
