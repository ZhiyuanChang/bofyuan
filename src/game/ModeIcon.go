package game

import "bofyuan/src/csvs"

type Icon struct {
	IconId int
}

type ModIcon struct {
	IconInfo map[int]*Icon
}

func (self *ModIcon) IsHasIcon(iconId int) bool {
	if _, ok := self.IconInfo[iconId]; ok {
		return true
	}
	return false
}

func (self *ModIcon) AddItem(itemId int) {
	if self.IsHasIcon(itemId) {
		return
	}
	if i := csvs.GetIconConfig(itemId); i == nil {
		//操作不合法
	}
	self.IconInfo[itemId] = &Icon{IconId: itemId}
}
