package ocr

import (
	"encoding/json"
	"errors"
)

func parseOCRServiceResponse(content *[]byte) (text *string, err error) {
	var jsonContent struct {
		ParsedResults []struct {
			ParsedText string `json:"ParsedText"`
		} `json:"ParsedResults"`
	}

	if jsonErr := json.Unmarshal(*content, &jsonContent); jsonErr != nil {
		err = errors.New("cannot parse JSON content: " + string(*content))
		return
	}

	text = &jsonContent.ParsedResults[0].ParsedText
	return
}
