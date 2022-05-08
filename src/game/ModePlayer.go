package game

import "fmt"

type ModPlayer struct {
	UserId      int64
	Icon        int
	Card        int
	Name        string
	Sign        string
	PlayerLevel int
	PlayerExp   int
	WorldLevel  int
	WorldTime   int64
	Birth       int
	ShowTeam    []int //后期替换英雄模块
	ShowCard    int
	//看不到的字段比这个要多很多
	IsProhibit int
	IsGM       int
}

func (self *ModPlayer) SetIcon(iconId int, player *Player) {

	if player.ModIcon.IsHasIcon(iconId) != true {
		//通知
		//记录
		//接入anti-cheat
	}
	player.ModPlayer.Icon = iconId
	fmt.Println()
}
func (self *ModPlayer) SetCard(cardId int, player *Player) {

	if player.ModCard.IsHasCard(cardId) != true {
		//通知
		//记录
		//接入anti-cheat
	}
	player.ModPlayer.Card = cardId
	fmt.Println()
}
func (self *ModPlayer) SetName(name string, player *Player) {

	//调用一个RPC接口
	if GetManageBanWord().IsBanWord(name) {
		return
	}
	player.ModPlayer.Name = name
	fmt.Println()
}
func (self *ModPlayer) SetSign(sign string, player *Player) {
	if GetManageBanWord().IsBanWord(sign) {
		return
	}
	player.ModPlayer.Sign = sign
	fmt.Println()
}
