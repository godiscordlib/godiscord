package common

import (
	"encoding/json"
	"log"
	"runtime"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocket struct {
	conn  *websocket.Conn
	ready chan struct{}
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

func newWebSocket() *WebSocket {
	conn, _, err := websocket.DefaultDialer.Dial(WEBSOCKET_URL, nil)
	if err != nil {
		log.Fatalln(err)
	}
	w := &WebSocket{
		conn:  conn,
		ready: make(chan struct{}),
	}
	return w
}

func (w *WebSocket) Connect(BotToken string, Intents int, WebSocketChannel chan webSocketPayload) {
	_, message, err := w.conn.ReadMessage()
	if err != nil {
		log.Fatalln("Error reading Hello", err)
	}

	var objMsg webSocketPayload
	if err := json.Unmarshal(message, &objMsg); err != nil {
		log.Fatalln("Error unmarshal Hello:", err)
	}
	if objMsg.OP != 10 {
		log.Fatalln("first message isn't a Hello (op 10)")
	}

	var helloWebSocket helloWebSocket
	if err := json.Unmarshal(objMsg.Data, &helloWebSocket); err != nil {
		log.Fatalln("Error unmarshal Hello:", err)
	}

	heartbeat_ticker := time.NewTicker(time.Duration(helloWebSocket.Heartbeats) * time.Millisecond)
	defer heartbeat_ticker.Stop()

	go func() {
		for range heartbeat_ticker.C {
			heartbeat := &webSocketPayload{
				OP:   1,
				Data: nil,
			}
			if err := w.conn.WriteJSON(heartbeat); err != nil {
				log.Println("Error sending heartbeat :", err)
				return
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
		log.Fatalln("Error sending Identify :", err)
	}

	for {
		_, msg, err := w.conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		var discordRes webSocketPayload
		if err := json.Unmarshal(msg, &discordRes); err != nil {
			log.Println("Error unmarshal message:", err)
			break
		}

		WebSocketChannel <- discordRes
	}
}

func (w *WebSocket) SendEvent(WSEventType int, Data any) error {
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
