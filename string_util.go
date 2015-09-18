package labmeasure

import (
	"regexp"
)

func normalize(s string) string {
	re := regexp.MustCompile(`[\n\r]+`)
	return re.ReplaceAllString(s, " ")
}

func countWord(s string) int {
	return len(getWords(s))
}

func getWords(s string) []string {
	re := regexp.MustCompile(`[ ,\(\)\<\>\'\"\{\}\[\]\-\_\.\?\!\;]+`)
	words := re.Split(s, -1)
	return words
}
