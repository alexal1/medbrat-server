package network

import (
	"encoding/json"
	"errors"
	"medbrat-server/usecase"
)

type responseStruct struct {
	Name string      `json:"name"`
	Body interface{} `json:"body"`
}

func parseCommand(content []byte) (commandName string, token string, body map[string]interface{}, err error) {
	var jsonCommand map[string]interface{}
	if jsonErr := json.Unmarshal(content, &jsonCommand); jsonErr != nil {
		err = errors.New("cannot parse JSON command: " + string(content))
		return
	}

	var ok bool

	commandName, ok = jsonCommand["name"].(string)
	if ok == false {
		err = errors.New("incorrect JSON: " + string(content))
		return
	}

	token, ok = jsonCommand["token"].(string)
	if ok == false {
		return
	}

	body, ok = jsonCommand["body"].(map[string]interface{})
	if ok == false {
		return
	}

	return
}

func parseAnswer(body map[string]interface{}) (toId uint64, value interface{}, err error) {
	var ok bool

	toIdFloat, ok := body["to"].(float64)
	if ok == false {
		err = errors.New("cannot parse answer body key \"to\"")
		return
	} else {
		toId = uint64(toIdFloat)
	}

	value, ok = body["value"]
	if ok == false {
		err = errors.New("cannot parse answer body key \"value\"")
		return
	}

	return
}

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

func createNewSessionJson(token string) (jsonNewSession []byte) {
	type startBody struct {
		Token string `json:"token"`
	}

	newSession := responseStruct{Name: "new_session", Body: startBody{Token: token}}
	jsonNewSession, _ = json.Marshal(newSession)
	return
}

func createMessagesJson(messages []*usecase.Message) (jsonMessages []byte) {
	jsonMessages, _ = json.Marshal(messages)
	return
}
