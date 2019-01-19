package network

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"io/ioutil"
	"log"
	"medbrat-server/usecase"
	"mime/multipart"
	"net/http"
	"time"
)

type ocrSpace struct {
	url          string
	method       string
	apikey       string
	maxImageSize int
}

func NewOCR() usecase.OCR {
	return &ocrSpace{
		"https://api.ocr.space/parse/image",
		"POST",
		"b19cd740b088957",
		1024 * 1024,
	}
}

func (ocr *ocrSpace) RecognizeGeneralBloodTest(image *string) (blood *map[usecase.BloodComponent]float32) {
	base64Image := "data:image/jpg;base64," + *compressImage(image, ocr.maxImageSize)

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("language", "rus")
	_ = writer.WriteField("base64Image", base64Image)
	_ = writer.WriteField("detectOrientation", "true")
	_ = writer.WriteField("scale", "true")
	_ = writer.WriteField("isTable", "true")

	err := writer.Close()
	if err != nil {
		log.Println("error when closing writer", err)
		return
	}

	client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(ocr.method, ocr.url, payload)
	if err != nil {
		log.Println("error when creating request", err)
		return
	}

	req.Header.Add("apikey", ocr.apikey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		log.Println("error when making request", err)
		return
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Println("error when closing connection", err)
			return
		}
	}()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("error when reading response", err)
		return
	}

	log.Println(string(body))

	return
}

// Compress given image to the given size in bytes.
func compressImage(imageString *string, size int) *string {
	imageBytes, err := base64.StdEncoding.DecodeString(*imageString)
	if err != nil {
		log.Println("cannot decode base64 image", err)
		return nil
	}
	imageReader := bytes.NewReader(imageBytes)
	image, err := jpeg.Decode(imageReader)
	if err != nil {
		log.Println("cannot decode jpeg image", err)
		return nil
	}

	compress := func(quality int) (compressedImage *[]byte, compressedSize int) {
		options := jpeg.Options{
			Quality: quality,
		}

		var buff bytes.Buffer
		err := jpeg.Encode(&buff, image, &options)
		if err != nil {
			log.Println("cannot encode jpeg image", err)
			return
		}

		result := buff.Bytes()

		return &result, buff.Len()
	}

	// Binary search for the best quality value
	intervalStart := 1
	intervalEnd := 101
	var compressedImageBytes *[]byte
	var compressedImageSize int
	var currentQuality int
	for intervalStart != intervalEnd {
		currentQuality = (intervalStart + intervalEnd) / 2
		compressedImageBytes, compressedImageSize = compress(currentQuality)
		if compressedImageSize < size {
			intervalStart = currentQuality + 1
		} else {
			intervalEnd = currentQuality
		}
	}

	imageStringResult := base64.StdEncoding.EncodeToString(*compressedImageBytes)

	return &imageStringResult
}
