package Week09

//709. 转换成小写字母
func toLowerCase(str string) string {
	b := []rune(str)
	for i, r := range str {
		if r >= 'A' && r <= 'Z' {
			b[i] = rune(r + 32)
		}
	}
	return string(b)
}

//58. 最后一个单词的长度
func lengthOfLastWord(s string) (ans int) {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		ans++
		index--
	}
	return
}
