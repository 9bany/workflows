package slack

import (
	"context"
	"log"

	"github.com/9bany/workflows/slack/cmds"
	"github.com/9bany/workflows/utils"
	"github.com/shomali11/slacker"
)

type SlackBot struct {
	slacker *slacker.Slacker
}

func NewBot(config utils.Config) *SlackBot {
	return &SlackBot{
		slacker: slacker.NewClient(config.BotToken, config.AppToken),
	}
}

func (bot *SlackBot) Start() error {

	bot.initDefault()
	bot.initCommands()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.slacker.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (bot *SlackBot) initDefault() {
	bot.slacker.Init(func() {
		log.Println("Connected!")
	})

	bot.slacker.Err(func(err string) {
		log.Println(err)
	})

	bot.slacker.DefaultEvent(func(event interface{}) {
		log.Println(event)
	})

}

func (bot *SlackBot) initCommands() {

	newsComand := cmds.NewNewsCommand()

	// jobs
	bot.slacker.Job(newsComand.JobDefinition())
	// commands
	bot.slacker.Command(cmds.PingCommandDefinition())
	bot.slacker.Command(newsComand.CommandDefinition())
	// help
	bot.slacker.Help(cmds.HelpCommandDefinition())
}
