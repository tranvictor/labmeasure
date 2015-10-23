package labmeasure

type ElementStat struct {
	Count      int
	Sum        float32
	Average    float32
	Maximum    float32
	MaximumURL string
	Data       map[string]float32
}

func (es *ElementStat) AddData(url string, time float32) {
	es.Data[url] = time
}

type TimeStat struct {
	stat map[string]ElementStat
}

func (ts *TimeStat) Add(name string, es ElementStat) {
	ts.stat[name] = es
}

func (ts TimeStat) GetStat(name string) ElementStat {
	return ts.stat[name]
}
