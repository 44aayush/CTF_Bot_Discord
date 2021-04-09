package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"aayush.com/ctfbot/brain"
	"aayush.com/ctfbot/handlers"
)

func main() {
	// Loading Discord bot API token from text file
	token, err := loadToken("token.dat")
	if err != nil {
		fmt.Println("Error loading token:", err)
		return
	}

	discord, err := discordgo.New("Bot "+token)
	if err != nil {
		fmt.Println("Error creating Discord client:", err)
		return
	}

	brain.Init("pool.json", "knowledgeStore.json")

	err = brain.Validate()
	if err != nil {
		fmt.Println("Failed to validate:", err)
		return
	}

	// Register handlers
	discord.AddHandler(handlers.OnReady)
	discord.AddHandler(handlers.OnMessageCreate)


	if err := discord.Open(); err != nil {
		fmt.Println("Error opening Discord connection:", err)
		return
	}
	defer discord.Close()

	time.Sleep(time.Second * 30)
}

func loadToken(filename string) (string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.TrimSpace(s) != "" {
			return s, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("%v did not contain a token ", filename)

}