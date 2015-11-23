package labmeasure

type PPubdateRecord struct {
	URL                  string
	DiffbotPubdateString string
	LabPubdateString     string
	Acceptable           bool
}

func (pr *PPubdateRecord) SetURL(url string) {
	pr.URL = url
}
