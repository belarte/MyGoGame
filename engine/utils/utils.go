package utils

// CompareEpsilon compare two floating point numbers
// with regards to a given epsilon.
func CompareEpsilon(left, right float64) bool {
	return left-right < 0.000001
}
