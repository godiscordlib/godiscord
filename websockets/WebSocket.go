package websockets

import (
	"encoding/json"
	"godiscord/types"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const WEBSOCKET_URL = "wss://gateway.discord.gg/?v=10&encoding=json"

type WebSocket struct {
}

func (w *WebSocket) Connect(BotToken string, Intents int, WebSocketChannel chan types.WebSocketPayload) {
	conn, _, err := websocket.DefaultDialer.Dial(WEBSOCKET_URL, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	_, message, err := conn.ReadMessage()

	if err != nil {
		log.Fatalln(err)
	}

	var objMsg types.WebSocketPayload
	if err := json.Unmarshal(message, &objMsg); err != nil {
		log.Fatalln(err)
	}
	if objMsg.OP != 10 {
		log.Fatalln("First message isn't an Hello (op 10)")
	}
	var helloWebSocket types.HelloWebSocket
	if err := json.Unmarshal(objMsg.Data, &helloWebSocket); err != nil {
		log.Fatalln(err)
	}
	heartbeat_ticker := time.NewTicker(time.Duration(helloWebSocket.Heartbeats * uint(time.Millisecond)))

	go func() {
		for {
			<-heartbeat_ticker.C
			heartbeat := &types.WebSocketPayload{
				OP:   1,
				Data: nil,
			}
			if err := conn.WriteJSON(heartbeat); err != nil {
				log.Fatalln(err)
			}
		}
	}()

	identifyData := &types.IdentifyWebSocketData{
		Token: BotToken,
		Properties: struct {
			OS      string `json:"os"`
			Browser string `json:"browser"`
			Device  string `json:"device"`
		}{
			OS:      "linux",
			Browser: "godiscord",
			Device:  "godiscord",
		},
		Intents: Intents,
	}
	identifyPayload := &types.WebSocketPayload{
		OP:   2,
		Data: marshal(identifyData),
	}

	if err := conn.WriteJSON(identifyPayload); err != nil {
		log.Fatalln(err)
	}

	var discordRes types.WebSocketPayload

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		err = json.Unmarshal(msg, &discordRes)
		if err != nil {
			log.Println(err)
			break
		}

		WebSocketChannel <- discordRes
	}

}

func marshal(v any) json.RawMessage {
	o, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return o
}
