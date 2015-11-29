package labmeasure

import (
	"reflect"
	"testing"
)

func TestBuildArticle(t *testing.T) {
	cases := []struct {
		in   string
		want Articles
	}{
		{
			`{"http://url.com": {
			  "text": "body text",
			  "body_total_time": 60.0,
			  "title": "title",
			  "title_total_time": 60.0,
			  "cleaner_total_time": 40.0,
			  "date": "Thu, 09 Jul 2015 00:00:00 GMT",
			  "ak": [1,4],
			  "published_date_total_time": 30.0,
			  "media": [{"link": "something"}],
			  "image_total_time": 100.1,
			  "image_computation_time": 40.4,
			  "extraction_type": "OG",
			  "extraction_total_time": 120.1,
			  "extraction_computation_time": 80.0
			}}`,
			Articles{
				"http://url.com": Article{
					"body text",
					60.0,
					"title",
					60.0,
					40.0,
					"Thu, 09 Jul 2015 00:00:00 GMT",
					[]int32{1, 4},
					30.0,
					ImageList{
						{"something"},
					},
					100.1,
					40.4,
					"OG",
					120.1,
					80.0,
				},
			},
		},
		{
			`{"http://url.com": {"text": "", "ak": [1, 3], "title": "title"}}`,
			Articles{
				"http://url.com": Article{
					"", 0,
					"title", 0,
					0,
					"",
					[]int32{1, 3},
					0,
					ImageList(nil), 0, 0,
					"", 0, 0,
				},
			},
		},
		{
			`{"http://url.com": {"ak": [1]}}`,
			Articles{
				"http://url.com": Article{
					"", 0,
					"", 0,
					0, "",
					[]int32{1},
					0,
					ImageList(nil), 0, 0,
					"", 0, 0},
			},
		},
	}

	for _, c := range cases {
		got := BuildArticles(c.in)
		// time.Date(2015, 7, 9, 0, 0, 0, time.UTC),
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf(
				"BuildArticles(%q) == \n%q, want \n%q",
				c.in, got, c.want)
		}
	}
}
