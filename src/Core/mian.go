package main

import (
	"bofyuan/src/csvs"
	"bofyuan/src/game"
	"fmt"
	"time"
)

func main() {
	csvs.CheckLoadCsv()
	go game.GetManageBanWord().Run()
	player := game.NewTestPlayer()
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			if k := time.Now().Unix(); k%3 == 0 {
				player.RecvSetIcon(int(k))
				player.RecvSetName("专业代练")
				fmt.Println(player.ModPlayer.Icon)
				fmt.Println(player.ModPlayer.Name)
			}
		}
	}
}
