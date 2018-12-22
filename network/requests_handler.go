package network

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"medbrat-server/configs"
	"time"
)

type responseStruct struct {
	Name string      `json:"name"`
	Body interface{} `json:"body"`
}

// Handle message from client
func handle(message []byte) (response []byte, err error) {
	var jsonCommand map[string]interface{}
	if err := json.Unmarshal(message, &jsonCommand); err != nil {
		return []byte{}, errors.New("Cannot parse JSON command: " + string(message))
	}

	name, ok := jsonCommand["name"].(string)
	if ok == false {
		return []byte{}, errors.New("Incorrect JSON: " + string(message))
	}
	switch name {
	case "start":
		return handleStart()
	default:
		return []byte{}, errors.New("Unknown command: " + name)
	}
}

func handleStart() (response []byte, err error) {
	type startBody struct {
		Token string `json:"token"`
	}

	if newToken, err := generateToken(); err != nil {
		return []byte{}, err
	} else {
		res := responseStruct{Name: "new_session", Body: startBody{Token: newToken}}
		jsonResponse, _ := json.Marshal(res)
		return jsonResponse, nil
	}
}

// Generate and return JWT token
func generateToken() (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token.Claims = claims
	return token.SignedString(configs.SigningKey)
}
