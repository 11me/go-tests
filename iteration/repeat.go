package iteration

func Repeat(char string, n int) string {
	var res string

	for i := 0; i < n; i++ {
		res += char
	}
	return res
}
