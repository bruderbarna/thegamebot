package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

const gameMessage string = "You lost the game!"
const ourGuildID = "137546139649048576"

func sendDirectMessage(discord *discordgo.Session, userID string, content string) (*discordgo.Message, error) {
	channel, err := discord.UserChannelCreate(userID)
	if err != nil {
		log.Fatalln("Failed to create user channel")
		return nil, errors.New("Failed to create user channel")
	}

	message, err := discord.ChannelMessageSend(channel.ID, content)
	if err != nil {
		log.Fatalln("Failed to create user channel")
		return nil, errors.New("Failed to create user channel")
	}

	return message, nil
}

func main() {
	if len(os.Args) != 1 {
		fmt.Println("usage: thegamebot <token>")
		return
	}
	token := os.Args[1]

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln("Failed to create discord session")
		return
	}

	for {
		guildMembers, err := discord.GuildMembers(ourGuildID, "0", 1000)
		if err != nil {
			log.Println("Failed to get guild")
			return
		}

		for _, member := range guildMembers {
			if member.User.Bot {
				continue
			}

			id := member.User.ID
			_, err := sendDirectMessage(discord, id, gameMessage)
			if err != nil {
				log.Println(err.Error())
			}
		}
		time.Sleep(time.Minute * 30)
	}
}
