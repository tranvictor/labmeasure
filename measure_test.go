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
							4.0 / 7.0,
							0.8,
							5,
							7,
							4,
							3,
							true,
						},
					},
				},
				map[string]Stater{
					"BodyComparer": BodyStat{
						1, 1, 0, []PBodyRecord{}, 0.8, 4.0 / 7.0,
						createTestConfig(),
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
		}
	}
}
