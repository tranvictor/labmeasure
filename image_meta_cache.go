package labmeasure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ImageMetaCache struct {
	FilePath string
	Width    int
	Height   int
	Type     string
	Phash    uint64
}

type ImageMetaCaches map[string]ImageMetaCache

func LoadImageMetaCachesPointer() *ImageMetaCaches {
	jsstring, err := ioutil.ReadFile("/Users/victor/image_caches/image_caches.json")
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}
	json_byte := []byte(jsstring)
	var result ImageMetaCaches
	json.Unmarshal(json_byte, &result)
	return &result
}

func SaveImageMetaCaches(cache *ImageMetaCaches) error {
	b, _ := json.MarshalIndent(cache, " ", "")
	err := ioutil.WriteFile("/Users/victor/image_caches/image_caches.json", b, 0666)
	return err
}
