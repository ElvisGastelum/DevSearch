package devsearchbot

import (
	"github.com/elvisgastelum/devsearchbot/controller"
	router "github.com/elvisgastelum/devsearchbot/http"
)

var (
	devSearchController controller.DevSearchController = controller.NewDevSearchController()
	httpRouter          router.Router                  = router.NewMuxRouter()
)

// Bot is the instance of dev search
// Simple example use:
//
// bot := devsearchbot.Bot{}
//
// bot.Start()
type Bot struct{}

// Start the server of dev search
func (b *Bot) Start() {
	const port string = ":3000"

	httpRouter.Post("/slack/slash-commands/devz-search", devSearchController.SlashCommand)

	httpRouter.Post("/slack/actions/devz-search", devSearchController.Actions)

	httpRouter.Serve(port)
}
