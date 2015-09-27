package labmeasure

type BodyComparer struct {
}

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

func (bc BodyComparer) Name() string {
	return "BodyComparer"
}

func (bc BodyComparer) Compare(diffbot, lab Article, config Config) PRecorder {
	record := PBodyRecord{}
	if !diffbot.HasBody() || !lab.HasBody() {
		return nil
	}
	normDiffbot := normalize(diffbot.Body)
	dbwords, dbSize := getMapWordCount(normDiffbot)
	normLab := normalize(lab.Body)
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
	record.Acceptable = isAcceptable(
		record.Precision, record.Recall,
		config.PrecisionThreshold, config.RecallThreshold)
	return &record

}

func (bc BodyComparer) Calculate(recorders Recorders, config Config) Stater {
	st := BodyStat{
		0, 0, 0, make([]PBodyRecord, 0), 0, 0, config,
	}
	for _, record := range recorders {
		if record != nil {
			bodyRecord := record.(*PBodyRecord)
			if bodyRecord.URL != "" {
				st.Examined += 1
				st.TotalPrecision += bodyRecord.Precision
				st.TotalRecall += bodyRecord.Recall
				if bodyRecord.Acceptable {
					st.Correct += 1
				} else {
					st.Incorrect += 1
					st.IncorrectRecords = append(
						st.IncorrectRecords, *bodyRecord)
				}
			}
		}
	}
	return st
}
