package utils

func CompareEpsilon(left, right float64) bool {
	return left-right < 0.000001
}
