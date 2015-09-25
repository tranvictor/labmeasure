package labmeasure

type Recorders []Recorder

type Stater interface{}

type FinalStat struct {
	recorders map[string]Recorders
	stats     map[string]Stater
}

func (st *FinalStat) AddRecordFor(name string, index int, record Recorder) {
	st.recorders[name][index] = record
}

func (st FinalStat) GetRecords(name string) Recorders {
	return st.recorders[name]
}

func (st *FinalStat) AddStat(name string, stat Stater) {
	st.stats[name] = stat
}

func (st *FinalStat) InitRecorders(name string, length int) {
	st.recorders[name] = make(Recorders, length)
}
