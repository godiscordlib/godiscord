package new

import (
	"os"

	"github.com/godiscordlib/godiscord/pkg/classes"
	"github.com/godiscordlib/godiscord/pkg/types"
)

func Client(Token string, Intents ...types.GatewayIntent) classes.Client {
	os.Setenv("GODISCORD_TOKEN", Token)
	return classes.Client{
		EventManager: classes.NewEventManager(),
		Intents:      Intents,
	}
}
