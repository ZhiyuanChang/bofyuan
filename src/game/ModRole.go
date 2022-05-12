package game

type ModRole struct {
}

func (self *ModRole) IsHasRole(v *ShowRole) bool {
	return true
}
func (self *ModRole) GetRoleLevel(roleId int) int {
	return 80
}
