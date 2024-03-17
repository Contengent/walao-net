package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func main() {
	/* Setting vars for bot */

	token := "your token goes here"
	guildID, catagoryID := "guildid here", "catagoryID here"

	heart := newHeart(token)
	channels := makeChannels(guildID, catagoryID)
	channels.makeSession(heart)
	heart.linkedChannelID = heart.getChannelID(channels)

	/* Setting up listener */
	// status handler :>
	heart.botSession.AddHandler(ready)
	heart.botSession.AddHandler(heart.messageCreate)

	//heart.botSession.AddHandler(messageCreate)
	// heart.sendMessage(channels, "noway dude")

	// Open the websocket and begin listening.
	heart.botSession.Identify.Intents = discordgo.IntentsGuildMessages

	err := heart.botSession.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	/* end */
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("A Walao-net bot is now running. Press RETURN to exit.")
	var debug_input string
	fmt.Scanln(&debug_input)

	heart.endSession()
}
