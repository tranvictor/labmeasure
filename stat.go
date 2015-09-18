package labmeasure

type Stat struct {
	Examined         int
	Correct          int
	Incorrect        int
	Records          []CompareRecord
	IncorrectRecords []CompareRecord
	TotalRecall      float32
	TotalPrecision   float32
	Configuration    Config
}

func (st Stat) Accuracy() float32 {
	return float32(st.Correct) / float32(st.Examined)
}

func (st Stat) AverageRecall() float32 {
	return float32(st.TotalRecall) / float32(st.Examined)
}

func (st Stat) AveragePrecision() float32 {
	return float32(st.TotalPrecision) / float32(st.Examined)
}

func (st Stat) PrecisionThreshold() float32 {
	return st.Configuration.PrecisionThreshold
}

func (st Stat) RecallThreshold() float32 {
	return st.Configuration.RecallThreshold
}

func (st *Stat) Calculate() {
	for _, record := range st.Records {
		if record.URL != "" {
			st.Examined += 1
			st.TotalPrecision += record.Precision
			st.TotalRecall += record.Recall
			if record.Acceptable {
				st.Correct += 1
			} else {
				st.Incorrect += 1
				st.IncorrectRecords = append(st.IncorrectRecords, record)
			}
		}
	}
}
