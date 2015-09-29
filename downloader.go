package labmeasure

import (
	"crypto/md5"
	"fmt"
	"github.com/quirkey/magick"
	"io"
	"net/http"
	"os"
)

type CacheImage struct {
	url  string
	hash string
}

type DownloadedImages struct {
	cacheImages []CacheImage
}

func (o DownloadedImages) URLs() []string {
	result := []string{}
	for _, cacheImage := range o.cacheImages {
		result = append(result, cacheImage.url)
	}
	return result
}

func (o DownloadedImages) Hashes() []string {
	result := []string{}
	for _, cacheImage := range o.cacheImages {
		result = append(result, cacheImage.hash)
	}
	return result
}

func (o *DownloadedImages) AddDownloadedImage(image CacheImage) {
	o.cacheImages = append(o.cacheImages, image)
}

func (o DownloadedImages) Size() int {
	return len(o.cacheImages)
}

func isQualified(filePath string) bool {
	image, err := magick.NewFromFile(filePath)
	defer image.Destroy()
	if err != nil {
		return false
	}
	if image.Type() == "GIF" {
		return false
	}
	width := image.Width()
	height := image.Height()
	ratio := float32(width) / float32(height)
	if width < 320 || height < 240 {
		return false
	}
	if (160.0/240.0 > ratio) || (ratio > 640.0/240.0) {
		return false
	}
	return true
}

func httpDownload(url, filePath string) bool {
	// if the file is not downloaded
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		fmt.Printf("%s", url)
		response, e := http.Get(url)
		if e != nil {
			return false
		}
		defer response.Body.Close()
		file, err := os.Create(filePath)
		defer file.Close()
		if err != nil {
			return false
		}
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return false
		}

	}
	return isQualified(filePath)
}

func download(urls []string) DownloadedImages {
	result := DownloadedImages{}
	for _, url := range urls {
		h := md5.New()
		io.WriteString(h, url)
		hash := h.Sum(nil)
		filePath := "/Users/victor/image_caches/" + string(hash)
		qualified := httpDownload(url, filePath)
		if qualified {
			result.AddDownloadedImage(CacheImage{
				url, string(hash),
			})
		}
	}
	return result
}
