package classes

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"godiscord.foo.ng/lib/internal/types"
)

type WebSocket struct {
	conn         *websocket.Conn
	ready        chan struct{}
	mutex        sync.Mutex
	lastPingTime time.Time
	Ping         int64
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
		Ping:  0,
	}
	return w
}

func (w *WebSocket) Connect(BotToken string, Intents []types.GatewayIntent, WebSocketChannel chan webSocketPayload) {
	var intents int
	for _, intent := range Intents {
		intents += int(intent)
	}
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

	fmt.Println(helloWebSocket.Heartbeats)

	heartbeat_ticker := time.NewTicker(time.Duration(helloWebSocket.Heartbeats) * time.Millisecond)
	defer heartbeat_ticker.Stop()

	w.conn.SetPongHandler(func(appData string) error {
		w.mutex.Lock()
		w.Ping = time.Since(w.lastPingTime).Milliseconds()
		w.mutex.Unlock()

		return nil
	})

	go func() {
		defer heartbeat_ticker.Stop()

		for range heartbeat_ticker.C {
			heartbeat := &webSocketPayload{
				OP:   1,
				Data: nil,
			}

			err := w.conn.WriteJSON(heartbeat)
			if err != nil {
				log.Println("❌ Heartbeat failed:", err)
				// tu pourrais aussi envoyer un signal de reconnexion ici
				return
			}

			w.mutex.Lock()
			w.lastPingTime = time.Now()
			w.mutex.Unlock()

			err = w.conn.WriteControl(
				websocket.PingMessage,
				[]byte("ping"),
				time.Now().Add(5*time.Second),
			)
			if err != nil {
				log.Println("❌ Erreur lors de l'envoi du ping :", err)
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
		Intents: intents,
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
