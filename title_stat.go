package labmeasure

type TitleStat struct {
	Examined         int
	Correct          int
	Incorrect        int
	Records          []TitleCompareRecord
	IncorrectRecords []TitleCompareRecord
	Configuration    Config
}
