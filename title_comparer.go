package labmeasure

type TitleComparer struct {
}

func (tc TitleComparer) Name() string {
	return "TitleComparer"
}

func (tc TitleComparer) Compare(diffbot, lab Article, config Config) PRecorder {
	if !diffbot.HasTitle() || !lab.HasTitle() {
		return nil
	}
	record := PTitleRecord{}
	record.DiffbotTitle = diffbot.Title
	record.LabTitle = lab.Title
	if diffbot.Title == lab.Title {
		record.Acceptable = true
	} else {
		record.Acceptable = false
	}
	return &record
}

func (tc TitleComparer) Calculate(recorders Recorders, config Config) Stater {
	ts := TitleStat{
		0, 0, 0, make([]PTitleRecord, 0), config,
	}
	for _, record := range recorders {
		if record != nil {
			titleRecord := record.(*PTitleRecord)
			if titleRecord.URL != "" {
				ts.Examined += 1
				if titleRecord.Acceptable {
					ts.Correct += 1
				} else {
					ts.Incorrect += 1
					ts.IncorrectRecords = append(
						ts.IncorrectRecords, *titleRecord)
				}
			}
		}
	}
	return ts
}
