package string

func replaceSpace(s string) string {
	var a string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			a += "%20"
		} else {
			a += string(s[i])
		}
	}
	return a
}
