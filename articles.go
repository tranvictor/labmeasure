package labmeasure

import (
	"encoding/json"
)

type ImageList []struct {
	Link string `json:"link,omitempty"`
}

type Article struct {
	Body                      string    `json:"text,omitempty"`
	BodyTotalTime             float32   `json:"body_total_time,omitempty"`
	Title                     string    `json:"title,omitempty"`
	TitleTotalTime            float32   `json:"title_total_time,omitempty"`
	Medias                    ImageList `json:"media,omitempty"`
	ImageTotalTime            float32   `json:"image_total_time,omitempty"`
	ImageComputationTime      float32   `json:"image_computation_time,omitempty"`
	ExtractionType            string    `json:"extraction_type,omitempty"`
	ExtractionTotalTime       float32   `json:"extraction_total_time,omitempty"`
	ExtractionComputationTime float32   `json:"extraction_computation_time,omitempty"`
}

func (a Article) HasBody() bool {
	return a.Body != ""
}

func (a Article) HasTitle() bool {
	return a.Title != ""
}

func (a Article) Images() []string {
	result := []string{}
	for _, image := range a.Medias {
		if image.Link != "" {
			result = append(result, image.Link)
		}
	}
	return result
}

type Articles map[string]Article

func BuildArticles(jsstring string) Articles {
	json_bytes := []byte(jsstring)
	var articles Articles
	json.Unmarshal(json_bytes, &articles)
	return articles
}
