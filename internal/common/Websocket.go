package common

import (
	"encoding/json"
	"log"
	"runtime"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocket struct {
	conn *websocket.Conn
}
type webSocketPayload struct {
	OP             int             `json:"op"`
	Data           json.RawMessage `json:"d"`
	SequenceNumber *int            `json:"s,omitempty"`
	EventName      string          `json:"t,omitempty"`
}
type helloWebSocket struct {
	Heartbeats uint `json:"heartbeat_interval"`
}

const WEBSOCKET_URL = "wss://gateway.discord.gg/?v=10&encoding=json"

func (w WebSocket) Connect(BotToken string, Intents int, WebSocketChannel chan webSocketPayload) {
	conn, _, err := websocket.DefaultDialer.Dial(WEBSOCKET_URL, nil)
	w.conn = conn
	if err != nil {
		log.Fatalln(err)
	}
	defer w.conn.Close()

	_, message, err := w.conn.ReadMessage()

	if err != nil {
		log.Fatalln(err)
	}

	var objMsg webSocketPayload
	if err := json.Unmarshal(message, &objMsg); err != nil {
		log.Fatalln(err)
	}
	if objMsg.OP != 10 {
		log.Fatalln("First message isn't an Hello (op 10)")
	}
	var helloWebSocket helloWebSocket
	if err := json.Unmarshal(objMsg.Data, &helloWebSocket); err != nil {
		log.Fatalln(err)
	}
	heartbeat_ticker := time.NewTicker(time.Duration(helloWebSocket.Heartbeats * uint(time.Millisecond)))

	go func() {
		for {
			<-heartbeat_ticker.C
			heartbeat := &webSocketPayload{
				OP:   1,
				Data: nil,
			}
			if err := w.conn.WriteJSON(heartbeat); err != nil {
				log.Fatalln(err)
			}
		}
	}()

	identifyData := &IdentifyWebSocketData{
		Token: BotToken,
		Properties: struct {
			OS      string `json:"os"`
			Browser string `json:"browser"`
			Device  string `json:"device"`
		}{
			OS:      runtime.GOOS,
			Browser: "godiscord",
			Device:  "godiscord",
		},
		Intents: Intents,
	}
	identifyPayload := &webSocketPayload{
		OP:   2,
		Data: marshal(identifyData),
	}

	if err := w.conn.WriteJSON(identifyPayload); err != nil {
		log.Fatalln(err)
	}

	var discordRes webSocketPayload

	for {
		_, msg, err := w.conn.ReadMessage()
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
func (w WebSocket) SendEvent(WSEventType int, Data any) error {
	payload := &webSocketPayload{
		OP:   WSEventType,
		Data: marshal(Data),
	}
	return w.conn.WriteJSON(payload)
}
func marshal(v any) json.RawMessage {
	o, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return o
}
