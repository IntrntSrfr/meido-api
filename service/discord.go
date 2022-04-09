package service

import "github.com/bwmarrin/discordgo"

type DiscordService interface {
}

type DiscordApiService struct {
	Token string
	Sess  *discordgo.Session
}

func NewDiscordService(token string) *DiscordApiService {
	s, err := discordgo.New(token)
	if err != nil {
		panic(err)
	}
	return &DiscordApiService{Token: token, Sess: s}
}
