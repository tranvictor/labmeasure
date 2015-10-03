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
