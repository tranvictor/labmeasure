package labmeasure

type TitleStat struct {
	Examined         int
	Correct          int
	Incorrect        int
	IncorrectRecords []PTitleRecord
	Configuration    Config
}

func (ts TitleStat) GetIncorrectRecords() interface{} {
	return ts.IncorrectRecords
}

func (ts TitleStat) Accuracy() float32 {
	if ts.Examined == 0 {
		return 0.0
	}
	return float32(ts.Correct) / float32(ts.Examined)
}
