package csvs

import "fmt"

type ConfigBanWord struct {
	Id  int
	Txt string
}

var ConfigBanwordSlice []*ConfigBanWord

func init() {
	ConfigBanwordSlice = append(ConfigBanwordSlice, &ConfigBanWord{
		Id:  1,
		Txt: "外挂",
	}, &ConfigBanWord{
		Id:  2,
		Txt: "辅助",
	}, &ConfigBanWord{
		Id:  1,
		Txt: "微信",
	}, &ConfigBanWord{
		Id:  1,
		Txt: "代练",
	}, &ConfigBanWord{
		Id:  1,
		Txt: "赚钱",
	}, &ConfigBanWord{
		Id:  1,
		Txt: "测试",
	})
	/*Loadcsv(ConfigBanwordSlice, "banword.csv")*/
	fmt.Println("csv_banword初始化")
}

func GetBaseBanWord() []string {
	relstring := make([]string, 0)
	for _, v := range ConfigBanwordSlice {
		relstring = append(relstring, v.Txt)
	}
	return relstring
}
