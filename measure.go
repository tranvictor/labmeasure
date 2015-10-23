// Package labmeasure stands for Lion Article Builder measure tool.
// It takes an JSON file as input and calculate some statistic such as
// accuracy, average precision, average recall to measure
// LAB performance in compared to Diffbot data.
//
// The JSON input should follow this format:
// { "<URL string>": { "text": "<Article Body>" } }
//
// Accuracy is calculated based on precision and recall thresholds.
package labmeasure

import (
	"fmt"
	"io/ioutil"
)

func getarticles(filename string) Articles {
	bytes, err := ioutil.ReadFile(filename)
	if err == nil {
		return BuildArticles(string(bytes))
	} else {
		panic(fmt.Sprintf("%q", err))
	}
}

type empty struct{}

func initFinalStat(length int, comparers []Comparer) FinalStat {
	result := FinalStat{
		map[string]Recorders{},
		map[string]Stater{},
		map[string]AggregateStater{},
	}
	for _, comparer := range comparers {
		result.InitRecorders(comparer.Name(), length)
	}
	return result
}

type Batch struct {
	LArticle Article
	Index    int
}

func splitArticles(larticles Articles, num int) []map[string]Batch {
	result := []map[string]Batch{}
	for i := 0; i < num; i++ {
		result = append(result, map[string]Batch{})
	}
	index := 0
	for url, larticle := range larticles {
		result[index%num][url] = Batch{
			larticle,
			index,
		}
		index += 1
	}
	return result
}

func analyze(darticles, larticles Articles, conf Config, result *FinalStat, comparers ...Comparer) {
	concurrency := 100
	sem := make(chan empty, concurrency)
	concurrencyBatches := splitArticles(larticles, concurrency)
	for pno, batch := range concurrencyBatches {
		go func(
			pno int, result *FinalStat,
			batches map[string]Batch, darticles *Articles) {
			for url, batch := range batches {
				larticle := batch.LArticle
				index := batch.Index
				if darticle, exist := (*darticles)[url]; exist {
					for _, comparer := range comparers {
						record := comparer.Compare(
							darticle, larticle, conf)
						if record != nil {
							record.SetURL(url)
							result.AddRecordFor(
								comparer.Name(), index, record)
						}
					}
				}
			}
			fmt.Printf("Done parallelism %d\n", pno)
			sem <- empty{}
		}(pno, result, batch, &darticles)
	}
	for i := 0; i < len(concurrencyBatches); i++ {
		<-sem
	}
	for _, comparer := range comparers {
		comparerName := comparer.Name()
		result.AddStat(
			comparerName,
			comparer.Calculate(
				result.GetRecords(comparerName), conf))
	}
}

func aggregate(larticles Articles, conf Config, st *FinalStat, aggregators ...Aggregator) {
	for _, aggregator := range aggregators {
		aggregateStater := aggregator.Calculate(larticles, conf)
		st.AddAggregateStat(aggregator.Name(), aggregateStater)
	}
}

func Measure(conf Config) FinalStat {
	darticles := getarticles(conf.DiffbotDataFile)
	larticles := getarticles(conf.LabDataFile)

	fmt.Printf("%d \n", len(darticles))
	fmt.Printf("%d \n", len(larticles))

	st := initFinalStat(len(larticles), []Comparer{BodyComparer{}, TitleComparer{}, ImageComparer{}})
	analyze(darticles, larticles, conf, &st,
		BodyComparer{}, TitleComparer{}, ImageComparer{})
	aggregate(larticles, conf, &st, TimeAggregator{})
	SaveImageMetaCaches(conf.ImageCaches)
	return st
}
