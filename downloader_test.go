package labmeasure

import (
	"reflect"
	"testing"
)

func TestDownload(t *testing.T) {
	cases := []struct {
		in   []string
		want DownloadedImages
	}{
		{
			[]string{
				"http://www4.pictures.zimbio.com/mp/SPUxm3_wzxNx.jpg",
			},
			DownloadedImages{
				[]CacheImage{
					{
						"http://www4.pictures.zimbio.com/mp/SPUxm3_wzxNx.jpg",
						"\xd4\xf3\xa9}\xb45f\xb3\xf5\x12\x1b\xd1\xc2\xc8q\xe2",
					},
				},
			},
		},
		{
			[]string{
				"http://www2.pictures.zimbio.com/mp/oeIE1sD0rLRx.jpg",
			},
			DownloadedImages{
				[]CacheImage{
					{
						"http://www2.pictures.zimbio.com/mp/oeIE1sD0rLRx.jpg",
						"\xb3\xfd A\xf6\xfc\x91^\x8d\x96\xf1\x9fH$q\xc5",
					},
				},
			},
		},
	}
	for _, c := range cases {
		got := download(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf(
				"download(%q) ==\n%q, want \n%q",
				c.in, got, c.want,
			)
		}
	}
}
