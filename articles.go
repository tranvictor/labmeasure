package labmeasure

import (
	"encoding/json"
)

type ImageList []struct {
	Link string
}

type Article struct {
	Body   string `json:"text,omitempty"`
	Title  string
	Images ImageList `json:",omitempty"`
}

func (a Article) HasBody() bool {
	return a.Body != ""
}

type Articles map[string]Article

func BuildArticles(jsstring string) Articles {
	json_bytes := []byte(jsstring)
	var articles Articles
	json.Unmarshal(json_bytes, &articles)
	return articles
}
