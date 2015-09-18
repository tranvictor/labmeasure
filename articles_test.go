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
			"{\"http://url.com\": {\"text\": \"body text\"}}",
			Articles{
				"http://url.com": Article{"body text"},
			},
		},
		{
			"{\"http://url.com\": {\"text\": \"\"}}",
			Articles{
				"http://url.com": Article{""},
			},
		},
		{
			"{\"http://url.com\": {}}",
			Articles{
				"http://url.com": Article{""},
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
