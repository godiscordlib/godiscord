package types

type IdentifyWebSocketData struct {
	Token      string `json:"token"`
	Intents    int    `json:"intents"`
	Properties struct {
		OS      string `json:"os"`
		Browser string `json:"browser"`
		Device  string `json:"device"`
	} `json:"properties"`
}
