package game

import "bofyuan/src/csvs"

type Card struct {
	CardId int
}

type ModCard struct {
	CardInfo map[int]*Card
}

func (self *ModCard) IsHasCard(cardId int) bool {
	if _, ok := self.CardInfo[cardId]; ok {
		return true
	}
	return false
}

func (self *ModCard) AddItem(itemId int, Friendliness int) {
	if self.IsHasCard(itemId) {
		return
	}
	config := csvs.GetCardConfig(itemId)
	if config == nil {
		//操作不合法
	}
	if Friendliness < config.Friendliness {
		//操作不合法
	}
	self.CardInfo[itemId] = &Card{CardId: itemId}
}
