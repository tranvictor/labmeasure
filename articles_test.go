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
			`{"http://url.com": {"text": "", "title": "title"}}`,
			Articles{
				"http://url.com": Article{
					"", 0,
					"title", 0,
					0, 0,
					ImageList(nil), 0, 0,
					"", 0, 0,
				},
			},
		},
		{
			"{\"http://url.com\": {}}",
			Articles{
				"http://url.com": Article{
					"", 0,
					"", 0,
					0, 0,
					ImageList(nil), 0, 0,
					"", 0, 0},
			},
		},
	}

	for _, c := range cases {
		got := BuildArticles(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf(
				"BuildArticles(%q) == \n%q, want \n%q",
				c.in, got, c.want)
		}
	}
}
