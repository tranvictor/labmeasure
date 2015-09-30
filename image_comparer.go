package labmeasure

import "github.com/kavu/go-phash"
import "fmt"

type ImageComparer struct {
}

func (o ImageComparer) Name() string {
	return "ImageComparer"
}

func phashEqual(labPath, diffbotPath string) bool {
	fmt.Printf("------------------------------------------\n")
	labHash, err := phash.ImageHashDCT(labPath)
	if err != nil {
		return false
	}
	diffbotHash, err := phash.ImageHashDCT(diffbotPath)
	if err != nil {
		return false
	}
	d, err := phash.HammingDistanceForHashes(labHash, diffbotHash)
	if err != nil {
		return false
	}
	fmt.Printf("Distance: %d between (%s, %s)\n", d, labPath, diffbotPath)
	fmt.Printf("#########################################\n")
	return d < 15
}

func compareImageList(diffbotImages, labImages DownloadedImages) (int, int) {
	lid := 0
	for _, labImage := range labImages.CacheImages {
		for _, diffbotImage := range diffbotImages.CacheImages {
			if labImage.URL == diffbotImage.URL {
				lid += 1
			} else {
				if phashEqual(labImage.Hash, diffbotImage.Hash) {
					lid += 1
				}
			}
		}
	}
	lnid := len(labImages.CacheImages) - lid
	return lid, lnid
}

func (o ImageComparer) Compare(diffbot, lab Article, config Config) PRecorder {
	record := PImageRecord{}
	localDiffbotImages := download(diffbot.Images())
	localLabImages := download(lab.Images())

	// fmt.Printf("localDiffbotImages: %q\n", localDiffbotImages)
	// fmt.Printf("localLabImages: %q\n", localLabImages)

	record.DiffbotImages = localDiffbotImages.URLs()
	record.LabImages = localLabImages.URLs()
	record.DiffbotSize = len(record.DiffbotImages)
	record.LabSize = len(record.LabImages)
	// both doesnt have any images
	if localDiffbotImages.Size()+localLabImages.Size() == 0 {
		record.Precision = 1.0
		record.Recall = 1.0
		record.LID = 0
		record.LNID = 0
		record.Acceptable = true
	} else {
		record.LID, record.LNID = compareImageList(localDiffbotImages, localLabImages)
		// fmt.Printf("LID - LNID: %d - %d", record.LID, record.LNID)
		// fmt.Printf("Image Record: %q \n", record)
		if record.LabSize == 0 {
			record.Precision = 1.0
		} else {
			record.Precision = float32(record.LID) / float32(record.LabSize)
		}
		if record.DiffbotSize == 0 {
			record.Recall = 1.0
		} else {
			record.Recall = float32(record.LID) / float32(record.DiffbotSize)
		}
		record.Acceptable = isAcceptable(record.Precision, record.Recall, 1, 0)
	}
	return &record
}

func (o ImageComparer) Calculate(recorders Recorders, config Config) Stater {
	st := ImageStat{
		0, 0, 0, 0, make([]PImageRecord, 0), config,
	}
	for _, record := range recorders {
		if record != nil {
			imageRecord := record.(*PImageRecord)
			if imageRecord.URL != "" {
				st.Examined += 1
				st.Qualified += 1
				if imageRecord.Acceptable {
					st.Correct += 1
				} else {
					st.Incorrect += 1
					st.IncorrectRecords = append(st.IncorrectRecords, *imageRecord)
				}
			}
		}
	}
	return st
}
