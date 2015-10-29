package labmeasure

import (
	"reflect"
	"testing"
)

func createTestConfig() Config {
	return Config{
		"test_data/diffbot.json",
		"test_data/lab.json",
		0.5,
		0.5,
		LoadImageMetaCachesPointer(),
	}
}

func TestMeasure(t *testing.T) {
	cases := []struct {
		in   Config
		want FinalStat
	}{
		{
			createTestConfig(),
			FinalStat{
				map[string]Recorders{
					"BodyComparer": Recorders{
						&PBodyRecord{
							"http://testURL1.com",
							"this is a test body",
							"is a test body with additional text",
							"Qualified",
							4.0 / 7.0,
							0.8,
							5,
							7,
							4,
							3,
							true,
						},
					},
					"TitleComparer": Recorders{
						&PTitleRecord{
							"http://testURL1.com",
							"this is the title",
							"this is the title",
							true,
						},
					},
					"ImageComparer": Recorders{
						&PImageRecord{
							"http://testURL1.com",
							[]string{
								"http://www4.pictures.zimbio.com/mp/SPUxm3_wzxNx.jpg",
							},
							[]string{
								"http://www4.pictures.zimbio.com/mp/SPUxm3_wzxNx.jpg",
							},
							1.0,
							1.0,
							1,
							1,
							1,
							0,
							true,
						},
					},
				},
				map[string]Stater{
					"BodyComparer": BodyStat{
						1, 0, 0, 0, 1, 1, 0, []PBodyRecord{}, 0.8, 4.0 / 7.0,
						createTestConfig(),
					},
					"TitleComparer": TitleStat{
						1, 1, 0, []PTitleRecord{}, createTestConfig(),
					},
					"ImageComparer": ImageStat{
						1, 1, 1, 0, []PImageRecord{}, createTestConfig(),
					},
				},
				map[string]AggregateStater{
					"TimeAggregator": TimeStat{
						map[string]ElementStat{
							"BodyTotalTime": ElementStat{
								1, 100.5, 100.5, 100.5, "http://testURL1.com", map[string]float32{"http://testURL1.com": 100.5}},
							"TitleTotalTime": ElementStat{
								1, 100.5, 100.5, 100.5, "http://testURL1.com", map[string]float32{"http://testURL1.com": 100.5}},
							"CleanerTotalTime": ElementStat{
								1, 30.0, 30.0, 30.0, "http://testURL1.com", map[string]float32{"http://testURL1.com": 30.0}},
							"PublishedDateTotalTime": ElementStat{
								1, 40.0, 40.0, 40.0, "http://testURL1.com", map[string]float32{"http://testURL1.com": 40.0}},
							"ImageTotalTime": ElementStat{
								1, 120.0, 120.0, 120.0, "http://testURL1.com", map[string]float32{"http://testURL1.com": 120.0}},
							"ImageComputationTime": ElementStat{
								1, 60.0, 60.0, 60.0, "http://testURL1.com", map[string]float32{"http://testURL1.com": 60.0}},
							"ExtractionTotalTime": ElementStat{
								1, 321, 321, 321, "http://testURL1.com", map[string]float32{"http://testURL1.com": 321}},
							"ExtractionComputationTime": ElementStat{
								1, 161, 161, 161, "http://testURL1.com", map[string]float32{"http://testURL1.com": 161}},
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		got := Measure(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf(
				"BuildArticles(%q) == \n%q,\n want \n%q",
				c.in, got, c.want)
			// t.Errorf(
			// 	"%q", got.AggregateStats())
		}
	}
}
