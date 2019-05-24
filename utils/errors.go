package utils

// Check checks for the error and panics if there is one
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
