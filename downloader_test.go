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
						"d4f3a97db43566b3f5121bd1c2c871e2",
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
						"b3fd2041f6fc915e8d96f19f482471c5",
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
