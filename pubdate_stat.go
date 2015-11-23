package labmeasure

type PubdateStat struct {
	Examined         int
	Correct          int
	Incorrect        int
	IncorrectRecords []PPubdateRecord
	Configuration    Config
}

func (ts PubdateStat) GetIncorrectRecords() interface{} {
	return ts.IncorrectRecords
}

func (ts PubdateStat) Accuracy() float32 {
	if ts.Examined == 0 {
		return 0.0
	}
	return float32(ts.Correct) / float32(ts.Examined)
}
