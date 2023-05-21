package server

import (
	"log"

	"github.com/9bany/workflows/slack"
	"github.com/9bany/workflows/utils"
)

type Server struct {
	slackBot *slack.SlackBot
}

func NewServer(config utils.Config) *Server {
	return &Server{
		slackBot: slack.NewBot(config),
	}
}

func (server Server) Start() {
	err := server.slackBot.Start()
	if err != nil {
		log.Panicln("Can not start server")
	}
}
