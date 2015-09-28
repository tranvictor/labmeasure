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
