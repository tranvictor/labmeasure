package labmeasure

import (
	"encoding/json"
)

type ImageList []struct {
	Link string `json:"link,omitempty"`
}

type Article struct {
	Body   string    `json:"text,omitempty"`
	Title  string    `json:"title,omitempty"`
	Images ImageList `json:"images,omitempty"`
}

func (a Article) HasBody() bool {
	return a.Body != ""
}

func (a Article) HasTitle() bool {
	return a.Title != ""
}

type Articles map[string]Article

func BuildArticles(jsstring string) Articles {
	json_bytes := []byte(jsstring)
	var articles Articles
	json.Unmarshal(json_bytes, &articles)
	return articles
}
