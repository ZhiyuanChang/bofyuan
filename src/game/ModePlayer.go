package game

import (
	"bofyuan/src/csvs"
	"fmt"
	"time"
)

type ShowRole struct {
	RoleId    int
	RoleLevel int
}

type ModPlayer struct {
	UserId         int64
	Icon           int
	Card           int
	Name           string
	Sign           string
	PlayerLevel    int
	PlayerExp      int
	WorldLevel     int
	WorldLevelNow  int
	WorldLevelCool int64
	WorldTime      int64
	Birth          int
	ShowTeam       []*ShowRole //后期替换英雄模块
	HideShowTeam   int         //展示隐藏的开关
	ShowCard       []int
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
func (self *ModPlayer) AddExp(Exp int, player *Player) {
	self.PlayerExp += Exp

	for {
		config := csvs.GetNowLevelConfig(self.PlayerLevel)
		if config == nil {
			//记录并处理
			break
		} else if config.PlayerExp == 0 {
			return
		}
		if config.ChapterId > 0 && player.ModUniqueTask.IsTaskFInish(config.ChapterId) {
			break
		}
		if self.PlayerExp >= config.PlayerExp {
			self.PlayerLevel++
			self.PlayerExp -= config.PlayerExp
		}
	}

}

func (self *ModPlayer) ReduceWorldLevel(player *Player) {
	if self.WorldLevel < csvs.REDUCE_WORLD_LEVEL_START {
		fmt.Println("降低世界等级操作失败", self.PlayerLevel, "当前经验：", self.PlayerExp)
	}
	if self.WorldLevel-self.WorldLevelNow >= csvs.REDUCE_WORLD_LEVEL_MAX {
		fmt.Println("降低世界等级操作失败", self.PlayerLevel, "当前世界等级：", self.WorldLevelNow)
	}
	if time.Now().Unix() < self.WorldLevelCool {
		fmt.Println("降低世界等级操作失败--冷却中")
	}
	self.WorldLevelNow -= 1
	self.WorldLevelCool = time.Now().Unix() + int64(csvs.REDUCE_WORLD_LEVEL_COOL_TIME)
	fmt.Println("操作成功")
}
func (self *ModPlayer) ReturnWorldLevel(self2 *Player) {
	if self.WorldLevelNow == self.WorldLevel {
		fmt.Println("操作失败", self.PlayerLevel, "当前世界等级：", self.WorldLevelNow)
	}
	if time.Now().Unix() < self.WorldLevelCool {
		fmt.Println("操作失败", self.PlayerLevel, "当前世界等级：", self.WorldLevelNow)
	}
	self.WorldLevelNow += 1
	self.WorldLevelCool = time.Now().Unix() + int64(csvs.REDUCE_WORLD_LEVEL_COOL_TIME)
	fmt.Println("操作成功")
}

func (self *ModPlayer) SetBirth(birth int, player *Player) {
	month := birth / 100
	day := birth % 100
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day <= 0 || day > 31 {
			fmt.Println(month, "月没有", day, "日")
		}
	case 4, 6, 9, 11:
		if day <= 0 || day > 30 {
			fmt.Println(month, "月没有", day, "日")
		}
	case 2:
		if day <= 0 || day > 30 {
			fmt.Println(month, "月没有", day, "日")
		}
	default:
		fmt.Println(month, "月没有")
	}
	self.Birth = birth
	fmt.Println("设置成功")
}

func (self *ModPlayer) SetShowCard(cards []int, player *Player) {
	//应该信任那些数据
	//服务端是远端，特别是客户端的环境是不安全的
	if len(cards) >= 9 {
		//错误
		//记录
		//anti-cheating接入
	}
	newList := make([]int, 0)
	cardsExist := make(map[int]int, len(cards))
	for _, v := range self.ShowCard {
		if _, ok := cardsExist[v]; ok {
			//错误
			//记录
			//anti-cheating接入
		}
		if !player.ModPlayer.IsHasCard(v) {
			//错误
			//记录
			//anti-cheating接入
		}
		newList = append(newList, v)
		cardsExist[v] = 1
	}
	self.ShowCard = newList
}

func (self *ModPlayer) IsHasCard(v int) bool {
	return true
}

func (self *ModPlayer) SetHideShowTeam(hide int, player *Player) {
	if hide == csvs.LOGIC_FALSE || hide == csvs.LOGIC_TRUE {
		player.ModPlayer.HideShowTeam = hide
	}
}

func (self *ModPlayer) SetShowTeam(roles []*ShowRole, player *Player) {
	if len(roles) >= 9 {
		//错误
		//记录
		//anti-cheating接入
	}
	newList := make([]*ShowRole, 0)
	roleExist := make(map[*ShowRole]int, len(roles))
	for _, v := range roles {
		if _, ok := roleExist[v]; ok {
			//错误
			//记录
			//anti-cheating接入
		}
		if !player.ModRole.IsHasRole(v) {
			//错误
			//记录
			//anti-cheating接入
		}
		newList = append(newList, v)
		roleExist[v] = 1
	}
	player.ModPlayer.ShowTeam = newList
}

func (self *ModPlayer) SetIsGM(isGM int) {
	if isGM == csvs.LOGIC_FALSE || isGM == csvs.LOGIC_FALSE {
		self.IsGM = isGM
	}
}
func (self *ModPlayer) SetProhibit(prohibit int) {
	self.IsProhibit = prohibit
}
