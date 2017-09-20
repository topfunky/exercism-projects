package accumulate

const testVersion = 1

// Accumulate executes a function `f` on every item in `list`.
// It returns a new slice.
func Accumulate(list []string, f func(string) string) (result []string) {
	result = make([]string, len(list))
	for index, s := range list {
		result[index] = f(s)
	}
	return
}
