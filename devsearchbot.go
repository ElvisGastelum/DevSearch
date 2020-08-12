package devsearchbot

import (
	"github.com/elvisgastelum/devsearchbot/controller"
	router "github.com/elvisgastelum/devsearchbot/http"
)

var (
	devSearchController controller.DevSearchController = controller.NewDevSearchController()
	httpRouter          router.Router                  = router.NewMuxRouter()
)

type bot struct{}

// Bot is the public interface to create the bot 
type Bot interface {
	Start() error
}

// NewDevSearchBot return the instance of dev search
// Simple example use:
//
// bot := devsearchbot.NewBot()
//
// bot.Start()
func NewDevSearchBot() Bot {
	return &bot{}
}

// Start the server of dev search bot
func (b *bot) Start() error {
	const port string = ":3000"

	httpRouter.Post("/slack/slash-commands/devz-search", devSearchController.SlashCommands)

	httpRouter.Post("/slack/actions/devz-search", devSearchController.Actions)

	err := httpRouter.Serve(port)
	if err != nil {
		return err
	}
	
	return nil
}
