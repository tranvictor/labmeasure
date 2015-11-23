package labmeasure

import "fmt"

type PubdateComparer struct {
}

func (pc PubdateComparer) Name() string {
	return "PubdateComparer"
}

func (pc PubdateComparer) Compare(diffbot, lab Article, config Config) PRecorder {
	record := PPubdateRecord{}
	record.DiffbotPubdateString = diffbot.PubdateString
	record.LabPubdateString = lab.PubdateString
	fmt.Printf("%q", diffbot.Pubdate())
	fmt.Printf("%q", lab.Pubdate())
	diffbotPubdateString := diffbot.Pubdate().Format("Jan 2 2006")
	labPubdateString := lab.Pubdate().Format("Jan 2 2006")
	if diffbotPubdateString == labPubdateString {
		record.Acceptable = true
	} else {
		record.Acceptable = false
	}
	return &record
}

func (pc PubdateComparer) Calculate(recorders Recorders, config Config) Stater {
	ts := PubdateStat{
		0, 0, 0, make([]PPubdateRecord, 0), config,
	}
	for _, record := range recorders {
		if record != nil {
			pubdateRecord := record.(*PPubdateRecord)
			if pubdateRecord.URL != "" {
				ts.Examined += 1
				if pubdateRecord.Acceptable {
					ts.Correct += 1
				} else {
					ts.Incorrect += 1
					ts.IncorrectRecords = append(
						ts.IncorrectRecords, *pubdateRecord)
				}
			}
		}
	}
	return ts
}
