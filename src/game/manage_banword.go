package game

import (
	"bofyuan/src/csvs"
	"fmt"
	"regexp"
	"time"
)

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

func (self *ManageBanWord) Run() {
	self.BanWordBase = csvs.GetBaseBanWord()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			if time.Now().Unix()%10 == 0 {
				fmt.Println("更新词库")
			} else {

			}
			fmt.Println("等待")
		}
	}
}
