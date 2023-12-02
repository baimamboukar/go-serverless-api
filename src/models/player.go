package models

import "gorm.io/gorm"

type KenganPlayer struct {
	gorm.Model
	Name        string `json:"name"`
	Nickname    string `json:"nickname"`
	Affiliation string `json:"affiliation"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
	Occupation  string `json:"occupation"`
	Wins        int    `json:"wins"`
	Losses      int    `json:"losses"`
	Height      string `json:"height"`
	Weight      string `json:"weight"`
	MartialArt  string `json:"martialArt"`
	BackStory   string `json:"backStory"`
	Abilities   string `json:"abilities"`
}
