package labmeasure

type BodyStat struct {
	Examined         int
	Correct          int
	Incorrect        int
	IncorrectRecords []BodyRecord
	TotalRecall      float32
	TotalPrecision   float32
	Configuration    Config
}

func (st BodyStat) Accuracy() float32 {
	return float32(st.Correct) / float32(st.Examined)
}

func (st BodyStat) AverageRecall() float32 {
	return float32(st.TotalRecall) / float32(st.Examined)
}

func (st BodyStat) AveragePrecision() float32 {
	return float32(st.TotalPrecision) / float32(st.Examined)
}

func (st BodyStat) PrecisionThreshold() float32 {
	return st.Configuration.PrecisionThreshold
}

func (st BodyStat) RecallThreshold() float32 {
	return st.Configuration.RecallThreshold
}
