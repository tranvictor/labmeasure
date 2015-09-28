package labmeasure

type ImageStat struct {
	Examined         int
	Qualified        int
	Correct          int
	Incorrect        int
	IncorrectRecords []PImageRecord
	Configuration    Config
}

func (o ImageStat) GetIncorrectRecords() interface{} {
	return o.IncorrectRecords
}

func (o ImageStat) Accuracy() float32 {
	if o.Qualified == 0 {
		return 0.0
	}
	return float32(o.Correct) / float32(o.Qualified)
}
