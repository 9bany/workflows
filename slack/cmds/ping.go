package cmds

import "github.com/shomali11/slacker"

var commandString = "ping"

func PingCommandDefinition() (string, *slacker.CommandDefinition) {
	definition := &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Yoo yoh !!! What's up man ?")
		},
	}
	return commandString, definition
}
