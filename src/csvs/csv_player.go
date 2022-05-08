package csvs

import "bofyuan/utils"

type ConfigPlayerLevel struct {
	PlayerLevel int `json:"PlayerLevel"`
	PlayerExp   int ` json:"PlayerExp"`
	WorldLevel  int `json:"WorldLevel"`
	ChapterId   int `json:"ChapterId"`
}

var (
	ConfigPlayerLevelSlice []*ConfigPlayerLevel
)

func init() {
	utils.GetCsvUtilMgr().LoadCsv("PlayerLevel", &ConfigPlayerLevelSlice)
}

func GetNowLevelConfig(level int) *ConfigPlayerLevel {
	if 0 > level || level >= len(ConfigBanwordSlice) {
		//错误记录
		return nil
	}
	return ConfigPlayerLevelSlice[level-1]
}
