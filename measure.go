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

type Comparer func(diffbot, lab string, pt, rt float32) CompareRecord

func analyze(darticles, larticles Articles, conf Config, comparer Comparer) Stat {
	result := Stat{
		0, 0, 0, make([]CompareRecord, len(larticles)),
		make([]CompareRecord, 0),
		0, 0, conf,
	}
	sem := make(chan empty, len(larticles))
	index := 0
	for url, larticle := range larticles {
		go func(
			index int, url string,
			larticle Article, darticles *Articles) {

			fmt.Printf("%d\n", index)
			if darticle, exist := (*darticles)[url]; exist {
				if darticle.has_body() && larticle.has_body() {
					record := comparer(
						darticle.Body, larticle.Body,
						conf.PrecisionThreshold,
						conf.RecallThreshold)
					record.URL = url
					result.Records[index] = record
				}
			}
			sem <- empty{}
		}(index, url, larticle, &darticles)
		index += 1
	}
	for i := 0; i < len(larticles); i++ {
		<-sem
	}
	result.Calculate()
	return result
}

func Measure(conf Config) Stat {
	darticles := getarticles(conf.DiffbotDataFile)
	larticles := getarticles(conf.LabDataFile)

	fmt.Printf("%d \n", len(darticles))
	fmt.Printf("%d \n", len(larticles))

	st := analyze(darticles, larticles, conf, compareBodyByWord)

	fmt.Printf("Number of examined articles: %d\n", st.Examined)
	fmt.Printf("Number of correct articles: %d\n", st.Correct)
	fmt.Printf("Number of incorrect articles: %d\n", st.Incorrect)
	fmt.Printf("Accuracy: %.2f\n", st.Accuracy())
	fmt.Printf("Average precision: %.2f\n", st.AveragePrecision())
	fmt.Printf("Average recall: %.2f\n", st.AverageRecall())
	fmt.Printf("Precision threshold: %f\n", conf.PrecisionThreshold)
	fmt.Printf("Recall threshold: %f\n", conf.RecallThreshold)

	return st
}
