package models

import "gorm.io/gorm"

type KenganPlayer struct {
	gorm.Model
	Name          string `json:"name"`
	Age           uint   `json:"age"`
	Weight        uint   `json:"weight"`
	Height        uint   `json:"height"`
	Image         string `json:"image"`
	Description   string `json:"description"`
	FightingStyle string `json:"fighting_style"`
	Origin        string `json:"origin"`
	Status        string `json:"status"`
	Wins          uint   `json:"wins"`
	Losses        uint   `json:"losses"`
	Draws         uint   `json:"draws"`
	Company       string `json:"company"`
}
