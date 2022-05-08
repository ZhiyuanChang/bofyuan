package game

type Player struct {
	ModPlayer *ModPlayer
	ModIcon   *ModIcon
	ModCard   *ModCard
}

func NewTestPlayer() *Player {
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModPlayer.Name = "goudan"
	//********************
	player.ModPlayer.Icon = 0
	//************************
	return player
}

func (self *Player) RecvSetIcon(iconId int) {
	self.ModPlayer.SetIcon(iconId, self)
}

func (self *Player) RecvSetCard(CardId int) {
	self.ModPlayer.SetCard(CardId, self)
}
func (self *Player) RecvSetName(name string) {
	self.ModPlayer.SetName(name, self)
}
func (self *Player) RecvSetSign(sign string) {
	self.ModPlayer.SetSign(sign, self)
}
