package arankup

func ReplaceCharAt(str string, index int, newChar byte) string {
	strBytes := []byte(str)
	strBytes[index] = newChar
	return string(strBytes)
}
