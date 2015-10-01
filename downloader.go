package labmeasure

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/kavu/go-phash"
	"github.com/quirkey/magick"
	"io"
	"net/http"
	"os"
	"time"
)

type CacheImage struct {
	URL      string
	FilePath string
}

type DownloadedImages struct {
	CacheImages []CacheImage
}

func (o DownloadedImages) URLs() []string {
	result := []string{}
	for _, cacheImage := range o.CacheImages {
		result = append(result, cacheImage.URL)
	}
	return result
}

func (o DownloadedImages) Hashes() []string {
	result := []string{}
	for _, cacheImage := range o.CacheImages {
		result = append(result, cacheImage.FilePath)
	}
	return result
}

func (o *DownloadedImages) AddDownloadedImage(image CacheImage) {
	o.CacheImages = append(o.CacheImages, image)
}

func (o DownloadedImages) Size() int {
	return len(o.CacheImages)
}

func isOk(t string, w int, h int) bool {
	if t == "GIF" {
		return false
	}
	ratio := float32(w) / float32(h)
	if w < 320 || h < 240 {
		return false
	}
	if (160.0/240.0 > ratio) || (ratio > 640.0/240.0) {
		return false
	}
	return true
}

func isQualified(filePath string, imageCaches *ImageMetaCaches) bool {
	if cache, exist := (*imageCaches)[filePath]; exist {
		return isOk(cache.Type, cache.Width, cache.Height)
	}
	image, err := magick.NewFromFile(filePath)
	if err != nil {
		return false
	}
	defer image.Destroy()
	result := isOk(image.Type(), image.Width(), image.Height())
	phash, err := phash.ImageHashDCT(filePath)
	(*imageCaches)[filePath] = ImageMetaCache{
		filePath,
		image.Width(),
		image.Height(),
		image.Type(),
		phash,
	}
	return result
}

func httpDownload(url, filePath string, config Config) bool {
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
		response.Body.Close()
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
	return isQualified(filePath, config.ImageCaches)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func download(urls []string, config Config) DownloadedImages {
	result := DownloadedImages{}
	for _, url := range urls {
		hash := getMD5Hash(url)
		filePath := "/Users/victor/image_caches/" + string(hash)
		qualified := httpDownload(url, filePath, config)
		if qualified {
			result.AddDownloadedImage(CacheImage{
				url, filePath,
			})
		}
	}
	return result
}
