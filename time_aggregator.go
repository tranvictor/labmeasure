package labmeasure

import "reflect"

type TimeAggregator struct{}

func (ta TimeAggregator) Calculate(articles Articles, config Config) AggregateStater {
	result := TimeStat{map[string]ElementStat{}}
	ta.statistic(articles, "BodyTotalTime", &result)
	ta.statistic(articles, "TitleTotalTime", &result)
	ta.statistic(articles, "ImageTotalTime", &result)
	ta.statistic(articles, "ImageComputationTime", &result)
	ta.statistic(articles, "ExtractionTotalTime", &result)
	ta.statistic(articles, "ExtractionComputationTime", &result)
	ta.statistic(articles, "CleanerTotalTime", &result)
	ta.statistic(articles, "PublishedDateTotalTime", &result)
	return AggregateStater(result)
}

func (ta TimeAggregator) statistic(articles Articles, name string, result *TimeStat) {
	elementStat := ElementStat{
		Maximum: -1,
		Data:    map[string]float32{},
	}
	for url, article := range articles {
		value := reflect.Indirect(reflect.ValueOf(&article)).FieldByName(name)
		time := float32(value.Float())
		if time > 0 {
			elementStat.Count++
			elementStat.Sum += time
			if elementStat.Maximum < time {
				elementStat.Maximum = time
				elementStat.MaximumURL = url
			}
			elementStat.AddData(url, time)
		}
	}
	if elementStat.Count > 0 {
		elementStat.Average = elementStat.Sum / float32(elementStat.Count)
	}
	result.Add(name, elementStat)
}

func (ta TimeAggregator) Name() string {
	return "TimeAggregator"
}
