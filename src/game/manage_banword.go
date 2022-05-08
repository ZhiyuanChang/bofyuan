package game

import "regexp"

var manageBanWord *ManageBanWord

type ManageBanWord struct {
	BanWordBase  []string
	BanWordExtra []string
}

func (self *ManageBanWord) IsBanWord(txt string) bool {
	for _, v := range self.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		if match {
			return match
		}
	}
	for _, v := range self.BanWordExtra {
		match, _ := regexp.MatchString(v, txt)
		if match {
			return match
		}
	}
	return true
}

func GetManageBanWord() *ManageBanWord {
	if manageBanWord == nil {
		manageBanWord = NewManageBanWord()
	}
	return manageBanWord
}
func NewManageBanWord() *ManageBanWord {
	return &ManageBanWord{}
}