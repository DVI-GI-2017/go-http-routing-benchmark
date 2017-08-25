package getval

import (
	"time"
)

func String(v, d string) string {
	if v == "" {
		return d
	}

	return v
}

func Float64(v, d float64) float64 {
	if v == 0 {
		return d
	}

	return v
}

func Int(v, d int) int {
	if v == 0 {
		return d
	}

	return v
}

func Int64(v, d int64) int64 {
	if v == 0 {
		return d
	}

	return v
}

func Time(v, d time.Time) time.Time {
	if v.IsZero() {
		return d
	}

	return v
}

func Rune(v, d rune) rune {
	if v == 0 {
		return d
	}

	return v
}
