package network

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"medbrat-server/configs"
	"medbrat-server/usecase"
	"time"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients and their tokens.
	clients map[*Client]string

	// Current usecases by tokens.
	usecases map[string]*usecase.GlobalUseCase

	// Inbound packets from the clients.
	broadcast chan Packet

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

type Packet struct {
	client  *Client
	content []byte
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan Packet),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]string),
		usecases:   make(map[string]*usecase.GlobalUseCase),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = ""
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case packet := <-h.broadcast:
			var response []byte
			command, token, body, err := parseCommand(packet.content)
			switch {
			case err != nil:
				log.Println(err)

			case command == "start":
				token := generateToken()
				h.clients[packet.client] = token

				response = createNewSessionJson(token)
				response = append(response, newline...)

				newUsecase := usecase.NewGlobal()
				h.usecases[token] = &newUsecase
				startMessages := newUsecase.Start()

				jsonMessages := createMessagesJson(startMessages)
				response = append(response, jsonMessages...)

			case command == "answer":
				log.Println("Answer!" + token + body.(string))
			}

			packet.client.send <- response
		}
	}
}

// Generate and return JWT token
func generateToken() (tokenString string) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token.Claims = claims
	tokenString, _ = token.SignedString(configs.SigningKey)
	return
}
