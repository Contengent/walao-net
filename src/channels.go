package main

import (
	_ "fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Channels struct {
	guildID    string
	catagoryID string
}

func makeChannels(guildID, catagoryID string) Channels {
	return Channels{
		guildID,
		catagoryID,
	}
}

func createChannelData(channelName string, catagoryID string) discordgo.GuildChannelCreateData {
	return discordgo.GuildChannelCreateData{
		Name:     channelName,
		Type:     0,
		ParentID: catagoryID,
	}
}

func (channels Channels) makeSession(heart *Heart) {
	var sessionNum = 1
	var channelList, _ = heart.botSession.GuildChannels(channels.guildID)

	for _, channel := range channelList {
		if strings.Contains(channel.Name, "session-"+strconv.Itoa(sessionNum)) {
			sessionNum++
		}
	}

	heart.linkedChannel = "session-" + strconv.Itoa(sessionNum)

	heart.botSession.GuildChannelCreateComplex(channels.guildID, createChannelData(heart.linkedChannel, channels.catagoryID))

}
