package labmeasure

func getMapWordCount(str string) (map[string]int, int) {
	result := make(map[string]int)
	words := getWords(str)
	total := 0
	for _, w := range words {
		result[w] += 1
		total += 1
	}
	return result, total
}

func compareBodyByWord(diffbot, lab string, pt, rt float32) CompareRecord {
	record := CompareRecord{}
	normDiffbot := normalize(diffbot)
	dbwords, dbSize := getMapWordCount(normDiffbot)
	normLab := normalize(lab)
	labwords, labSize := getMapWordCount(normLab)
	lid := 0
	lnid := 0
	for w, labcount := range labwords {
		if dbcount, exist := dbwords[w]; exist {
			if labcount < dbcount {
				lid += labcount
			} else {
				lid += dbcount
				lnid += labcount - dbcount
			}
		} else {
			lnid += labcount
		}
	}
	record.DiffbotBody = normDiffbot
	record.LabBody = normLab
	record.LID = lid
	record.LNID = lnid
	record.DiffbotSize = dbSize
	record.LabSize = labSize
	record.Precision = float32(record.LID) / float32(record.LabSize)
	record.Recall = float32(record.LID) / float32(record.DiffbotSize)
	record.Acceptable = isAcceptable(record.Precision, record.Recall, pt, rt)
	return record
}
