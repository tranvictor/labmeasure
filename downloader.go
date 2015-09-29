package labmeasure

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/quirkey/magick"
	"io"
	"net/http"
	"os"
	"time"
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
	if err != nil {
		return false
	}
	defer image.Destroy()
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
		timeout := time.Duration(1 * time.Minute)
		client := http.Client{
			Timeout: timeout,
		}
		fmt.Printf("Going to get: %s\n", url)
		response, e := client.Get(url)
		if e != nil {
			return false
		}
		defer response.Body.Close()
		file, err := os.Create(filePath)
		if err != nil {
			return false
		}
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return false
		}
		file.Close()
		fmt.Printf("--> Done: %s to %s\n", url, filePath)
	}
	return isQualified(filePath)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func download(urls []string) DownloadedImages {
	result := DownloadedImages{}
	for _, url := range urls {
		hash := getMD5Hash(url)
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
