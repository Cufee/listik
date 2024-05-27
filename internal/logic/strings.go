package logic

func StringIfElse(condition bool, onTrue, onFalse string) string {
	if condition {
		return onTrue
	}
	return onFalse
}
