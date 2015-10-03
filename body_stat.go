package labmeasure

type BodyStat struct {
	Examined         int
	BothEmpty        int
	DiffbotEmpty     int
	LabEmpty         int
	Qualified        int
	Acceptable       int
	Unacceptable     int
	IncorrectRecords []PBodyRecord
	TotalRecall      float32
	TotalPrecision   float32
	Configuration    Config
}

func (st BodyStat) GetIncorrectRecords() interface{} {
	return st.IncorrectRecords
}

func (st BodyStat) Correct() int {
	return st.Acceptable + st.BothEmpty
}

func (st BodyStat) Incorrect() int {
	return st.Unacceptable + st.DiffbotEmpty + st.LabEmpty
}

func (st BodyStat) Accuracy() float32 {
	return float32(st.Correct()) / float32(st.Examined)
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
