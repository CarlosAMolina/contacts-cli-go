package main

type Contact struct {
	ID            int           `json:"id"`
	Addresses     []string      `json:"addresses"`
	Categories    []string      `json:"categories"`
	Emails        []string      `json:"emails"`
	Name          string        `json:"name"`
	Nicknames     []string      `json:"nicknames"`
	Note          string        `json:"note"`
	Phones        []Phone       `json:"phones"`
	SocialNetwork SocialNetwork `json:"socialNetwork"`
	Surname       string        `json:"surname"`
	Urls          []string      `json:"urls"`
}

type Phone struct {
	Description string `json:"description"`
	Number      string `json:"number"`
}

type SocialNetwork struct {
	DiscordAccounts   []Discord  `json:"discordAccounts"`
	FacebookAccounts  []string   `json:"facebookAccounts"`
	GithubAccounts    []string   `json:"githubAccounts"`
	InstagramAccounts []string   `json:"instagramAccounts"`
	LinkedinAccounts  []string   `json:"linkedinAccounts"`
	TelegramAccounts  []string   `json:"telegramAccounts"`
	TiktokAccounts    []string   `json:"tiktokAccounts"`
	TwitterAccounts   []string   `json:"twitterAccounts"`
	WallapopAccounts  []Wallapop `json:"wallapopAccounts"`
}

type Discord struct {
	Alias          string `json:"alias"`
	Discriminator  int    `json:"discriminator"`
	GlobalName     string `json:"globalName"`
	LegacyUserName string `json:"legacyUserName"`
	UserName       string `json:"userName"`
}

type Wallapop struct {
	URL  string `json:"url"`
	Note string `json:"note"`
}

type Data struct {
	Contacts []Contact `json:"contacts"`
}
