package discord

import (
	"evaluation/main.go/calculatrice"
	"evaluation/main.go/config"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Run() {
	//calculatrice = evaluation.calculatrice()

	goBot, err := discordgo.New("Bot " + config.Config.Token)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	user, err := goBot.User("@me")
	BotID = user.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()

	if err != nil {
		return
	}
	fmt.Println("Bot is running !")
	if err != nil {
		fmt.Printf("user: %v\n", user)
		log.Fatal(err.Error())
		return
	}
}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	new_text := config.Split(m.Content)
	if len(new_text) == 4 {

		list, _ := s.GuildMembersSearch(m.GuildID, new_text[1], 10)
		if BotID == m.Author.ID || m.ChannelID != "1032264204289323068" {
			return
		}
		for _, v := range list {
			if v.User.Username == new_text[2] {
				s.ChannelMessageSend(m.ChannelID, v.Mention())
			}
		}
		fmt.Printf("m.Content: %v\n", m.Content)
		if StartWith(m.Content, "calculatrice") {
			calculatrice.Calculatrice(calcul[' '], rune)

		}
	}
}

func StartWith(content string, patern string) bool {
	panic("unimplemented")
	if len(patern) > len(content) {
		return false
	}
	for i, v := range patern {
		if rune(content[i]) != v {
			return false
		}
	}
	return true
}
