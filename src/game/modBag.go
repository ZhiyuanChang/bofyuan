package game

import (
	"bofyuan/src/csvs"
	"fmt"
)

type Item struct {
	ID      int
	ItemNum int64
}

type ModBag struct {
	Bag map[int]*Item
}

func (self *ModBag) AddItem(itemId int, player *Player) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "物品不存在")
	}
	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		fmt.Println(itemId, "普通")
	case csvs.ITEMTYPE_ROLE:
		fmt.Println(itemId, "角色")
	case csvs.ITEMTYPE_ICON:
		player.ModIcon.AddItem(itemId)
	case csvs.ITEMTYPE_CARD:
		player.ModCard.AddItem(itemId, 1111111)
	default:
		self.AddItemToBag(itemId, 1)
	}
}

func (self *ModBag) AddItemToBag(itemId int, num int64) {
	if _, ok := self.Bag[itemId]; ok {
		self.Bag[itemId].ItemNum += num

	}
	config := csvs.GetItemConfig(itemId)
	if config == nil {
		//错误处理
	}
	self.Bag[itemId] = &Item{
		ID:      itemId,
		ItemNum: num,
	}
}
func (self *ModBag) RemoveItemToBagGM(itemId int, num int64) {
	_, ok := self.Bag[itemId]
	if ok {
		self.Bag[itemId].ItemNum -= num
	} else {
		self.Bag[itemId] = &Item{ID: itemId, ItemNum: 0 - num}
	}
	config := csvs.GetItemConfig(itemId)
	if config != nil {
		fmt.Println("扣除物品", config.ItemName, "----数量：", num, "----当前数量：", self.Bag[itemId].ItemNum)
	}
}
func (self *ModBag) RemoveItem(itemId int, num int64) {
	itemConfig := csvs.GetItemConfig(itemId)
	if itemConfig == nil {
		fmt.Println(itemId, "物品不存在")
		return
	}

	switch itemConfig.SortType {
	case csvs.ITEMTYPE_NORMAL:
		self.RemoveItemToBagGM(itemId, num)
	default: //同普通
		//self.AddItemToBag(itemId, 1)
	}
}
