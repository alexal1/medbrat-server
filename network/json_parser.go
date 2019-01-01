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

func parseCommand(content []byte) (commandName string, token string, body interface{}, err error) {
	var jsonCommand map[string]interface{}
	if jsonErr := json.Unmarshal(content, &jsonCommand); jsonErr != nil {
		err = errors.New("Cannot parse JSON command: " + string(content))
		return
	}

	var ok bool

	commandName, ok = jsonCommand["name"].(string)
	if ok == false {
		err = errors.New("Incorrect JSON: " + string(content))
		return
	}

	token, ok = jsonCommand["token"].(string)
	if ok == false {
		return
	}

	body, ok = jsonCommand["body"]
	if ok == false {
		return
	}

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
