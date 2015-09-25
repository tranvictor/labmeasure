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
			`{"http://url.com": {"text": "body text", "title": "title", "images": [{"link": "something"}]}}`,
			Articles{
				"http://url.com": Article{
					"body text",
					"title",
					ImageList{
						{"something"},
					},
				},
			},
		},
		{
			`{"http://url.com": {"text": "", "title": "title"}}`,
			Articles{
				"http://url.com": Article{
					"",
					"title",
					ImageList(nil),
				},
			},
		},
		{
			"{\"http://url.com\": {}}",
			Articles{
				"http://url.com": Article{"", "", ImageList(nil)},
			},
		},
	}

	for _, c := range cases {
		got := BuildArticles(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf(
				"BuildArticles(%q) == %q, want %q",
				c.in, got, c.want)
		}
	}
}
