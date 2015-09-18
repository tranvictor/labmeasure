package labmeasure

type CompareRecord struct {
	URL         string
	DiffbotBody string
	LabBody     string
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
