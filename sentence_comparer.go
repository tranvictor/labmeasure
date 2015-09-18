package labmeasure

import (
	"regexp"
	"strings"
)

func getSentences(str string) (map[string]int, int) {
	result := make(map[string]int)
	re := regexp.MustCompile(`[^\.\?\!\;]+([\.\?\!\;]+|\z)`)
	sentences := re.FindAllString(str, -1)
	total := 0
	for _, s := range sentences {
		words := countWord(s)
		result[strings.Trim(s, " ")] = words
		total += words
	}
	return result, total
}

func compareBodyBySentence(diffbot, lab string, pt, rt float32) CompareRecord {
	record := CompareRecord{}
	normDiffbot := normalize(diffbot)
	dbsens, dbwords := getSentences(normDiffbot)
	normLab := normalize(lab)
	labsens, labwords := getSentences(normLab)
	lid := 0
	lnid := 0
	for s := range labsens {
		if size, exist := dbsens[s]; exist {
			lid += size
		} else {
			lnid += size
		}
	}
	record.DiffbotBody = normDiffbot
	record.LabBody = normLab
	record.LID = lid
	record.LNID = lnid
	record.DiffbotSize = dbwords
	record.LabSize = labwords
	record.Precision = float32(record.LID) / float32(record.LabSize)
	record.Recall = float32(record.LID) / float32(record.DiffbotSize)
	record.Acceptable = isAcceptable(record.Precision, record.Recall, pt, rt)
	return record
}
