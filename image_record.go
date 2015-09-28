package labmeasure

type PImageRecord struct {
	URL           string
	DiffbotImages []string
	LabImages     []string
	Precision     float32
	Recall        float32
	DiffbotSize   int
	LabSize       int
	LID           int
	LNID          int
	Acceptable    bool
}

func (br *PImageRecord) SetURL(url string) {
	br.URL = url
}
