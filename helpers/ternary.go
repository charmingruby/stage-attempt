package helpers

func If[R any, L any](condition bool, right R, left L) any {
	if condition {
		return right
	}

	return left
}
