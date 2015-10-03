package labmeasure

type PBodyRecord struct {
	URL         string
	DiffbotBody string
	LabBody     string
	Type        string
	Precision   float32
	Recall      float32
	DiffbotSize int
	LabSize     int
	LID         int
	LNID        int
	Acceptable  bool
}

func isAcceptable(precision, recall float32, pt, rt float32) bool {
	return precision >= pt && recall >= rt
}

func (br *PBodyRecord) SetURL(url string) {
	br.URL = url
}

func (br *PBodyRecord) DecideType(diffbotBody, labBody string) {
	br.DiffbotBody = diffbotBody
	br.LabBody = labBody
	if diffbotBody == "" {
		br.Type = "DiffbotEmpty"
	}
	if labBody == "" {
		br.Type = "LabEmpty"
	}
	if diffbotBody == "" && labBody == "" {
		br.Type = "BothEmpty"
	}
	if diffbotBody != "" && labBody != "" {
		br.Type = "Qualified"
	}
}
