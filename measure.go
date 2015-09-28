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
	}
	for _, comparer := range comparers {
		result.InitRecorders(comparer.Name(), length)
	}
	return result
}

func analyze(darticles, larticles Articles, conf Config, comparers ...Comparer) FinalStat {
	result := initFinalStat(len(larticles), comparers)
	sem := make(chan empty, len(larticles))
	index := 0
	for url, larticle := range larticles {
		go func(
			index int, url string, result *FinalStat,
			larticle Article, darticles *Articles) {
			fmt.Printf("%d\n", index)
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
			sem <- empty{}
		}(index, url, &result, larticle, &darticles)
		index += 1
	}
	for i := 0; i < len(larticles); i++ {
		<-sem
	}
	for _, comparer := range comparers {
		comparerName := comparer.Name()
		result.AddStat(
			comparerName,
			comparer.Calculate(
				result.GetRecords(comparerName), conf))
	}
	return result
}

func Measure(conf Config) FinalStat {
	darticles := getarticles(conf.DiffbotDataFile)
	larticles := getarticles(conf.LabDataFile)

	fmt.Printf("%d \n", len(darticles))
	fmt.Printf("%d \n", len(larticles))

	st := analyze(darticles, larticles, conf, BodyComparer{}, TitleComparer{})
	return st
}
