# walao-net
A proof of concept golang and discord bot based botnet


## Reasoning
Fun, I don't know. I thought it'd be interesting to do something like this. <br>
Turns out it was really easy to make something like this. But that's probably because the hard part is
making it secure, and hard to reverse engineer. Anyways, go is fun... but I should really get to
doing more rust programming lol.

## Building
Download the src file from github, and run the follow in that directory: <br>
`go mod init github.com/Contengent/walao-net` <br>
`go mod tidy` <br>
`go mod build`

## Usage
If you wanna play around with this, download it and build it after adding your token, guildID, and catagoryID.
It will make a new channel under the specified catagory, named: "session-[number]" everytime it's run.
It'll delete this if channel you exit the session properly. <br>
Commands are really barebones:
- !run [command] - Runs a command
- !exit - Exits the session

GuildID and CatagoryID can be found by enabling developer tools, then right clicking the server/catagory,
and clicking "Copy Server/Channel ID" <br>
(Note: I don't know why, but the catagory id is called "Channel ID" when you try to copy it lol)

### I DO NOT CLAIM RESPONSABILITY FOR ANY MISUE
If you do something profoundly stupid with this repository, you should consider seeking medical attention regaurding your mental health. This is no where near good or secure enough to be part of a real botnet, and it never was designed to be. This is a proof of concept using a discord bot, inside a server, as a C2 server. TL;DR: If you wanna commit federal crimes, and you use this repository, and you inevitably get caught: You can't blame me.
