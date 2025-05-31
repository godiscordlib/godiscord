package new

import (
	"os"

	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/types"
)

func Client(Token string, Intents ...types.GatewayIntent) classes.Client {
	os.Setenv("GODISCORD_TOKEN", Token)
	return classes.Client{
		EventManager: classes.NewEventManager(),
		Intents:      Intents,
	}
}
