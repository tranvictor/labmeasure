package labmeasure

type PPubdateRecord struct {
	URL                  string
	DiffbotPubdateString string
	LabPubdateString     string
	Acceptable           bool
	Ak                   []int32
}

func (pr *PPubdateRecord) SetURL(url string) {
	pr.URL = url
}
