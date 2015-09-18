package labmeasure

import (
	"encoding/json"
)

type Article struct {
	Body string `json:"text,omitempty"`
}

func (a Article) has_body() bool {
	return a.Body != ""
}

type Articles map[string]Article

func BuildArticles(jsstring string) Articles {
	json_bytes := []byte(jsstring)
	var articles Articles
	json.Unmarshal(json_bytes, &articles)
	return articles
}
