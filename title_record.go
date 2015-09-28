package labmeasure

type PTitleRecord struct {
	URL          string
	DiffbotTitle string
	LabTitle     string
	Acceptable   bool
}

func (tr *PTitleRecord) SetURL(url string) {
	tr.URL = url
}
