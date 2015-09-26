package labmeasure

type Recorders []PRecorder

type Stater interface {
	GetIncorrectRecords() interface{}
}

type FinalStat struct {
	recorders map[string]Recorders
	stats     map[string]Stater
}

func (st *FinalStat) Stats() map[string]Stater {
	return st.stats
}

func (st *FinalStat) AddRecordFor(name string, index int, record PRecorder) {
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
