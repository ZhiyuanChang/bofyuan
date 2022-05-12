package game

const (
	TASK_STATE_INIT   = 0
	TASK_STATE_DOING  = 1
	TASK_STATE_FINISH = 2
)

type Player struct {
	ModPlayer     *ModPlayer
	ModIcon       *ModIcon
	ModCard       *ModCard
	ModUniqueTask *ModUniqueTask
	ModRole       *ModRole
	ModItem       *ModBag
}

func NewTestPlayer() *Player {
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModUniqueTask = NewModUniqueTask()
	player.ModRole = new(ModRole)
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
func (self *Player) ReduceWorldLevel() {
	self.ModPlayer.ReduceWorldLevel(self)
}
func (self *Player) SetBirth(birth int) {
	self.ModPlayer.SetBirth(birth, self)
}
func (self *Player) SetShowCard(cards []int) {
	self.ModPlayer.SetShowCard(cards, self)
}
func (self *Player) SetShowTeam(role []*ShowRole) {
	self.ModPlayer.SetShowTeam(role, self)
}
func (self *Player) SetHideShowTeam(isHide int) {
	self.ModPlayer.SetHideShowTeam(isHide, self)
}
