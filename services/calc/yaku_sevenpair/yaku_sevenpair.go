package yaku_sevenpair

import (
	"fmt"
	"log"

	"github.com/dQw4w/jpmj_calc/services/calc/win"
)

func Non_Yakuman_Special(sp win.Seven_Pairs_Win) (int, string) {
	var han int
	var msg string
	if sp.DoubleReach {
		han += 2
		msg += "ダブル立直 2飜\n"
	} else if sp.Reach {
		han++
		msg += "立直 1飜\n"
	}
	if sp.Ippatsu {
		han++
		msg += "一発 1飜\n"
	}
	if sp.RinShan {
		han++
		msg += "嶺上開花 1飜\n"
	}
	if sp.HaiTei {
		han++
		msg += "海底摸月 1飜\n"
	}
	if sp.HoTei {
		han++
		msg += "河底撈魚 1飜\n"
	}
	return han, msg
}

func MenchinTsumo(sp win.Seven_Pairs_Win) (int, string) {
	if !(sp.Menchin && sp.Tsumo) {
		return 0, ""
	}
	return 1, "門前清自摸和 1飜\n"
}

func Tanyao(sp win.Seven_Pairs_Win) (int, string) {
	for i := range sp.PairList {
		pair := sp.PairList[i]
		if pair.Suit == 'z' || pair.Rank == 1 || pair.Rank == 9 {
			return 0, ""
		}
	}
	return 1, "断么九 1飜\n"
}

func HonOldHead(sp win.Seven_Pairs_Win) (int, string) {
	for i := range sp.PairList {
		pair := sp.PairList[i]
		if pair.Suit == 'z' || pair.Rank == 1 || pair.Rank == 9 {
			continue
		}
		return 0, ""
	}

	return 2, "混老頭 2飜\n"

}

func Somete(sp win.Seven_Pairs_Win) (int, string) {
	var havezi bool = false
	var suit byte = 'x' // initial state, represents nothing
	for i := range sp.PairList {
		pair := sp.PairList[i]
		if pair.Suit == 'z' {
			havezi = true
		} else {
			if suit == 'x' {
				suit = pair.Suit
			} else if suit != pair.Suit {
				return 0, ""
			}
		}
	}
	if suit == 'x' {
		return 0, ""
	}
	if havezi {
		return 3, "混一色 3飜\n"
	}
	return 6, "清一色 6飜\n"

}

// below are yakumans, the return integer value represents how many times of yakuman
func Yakuman_Special(sp win.Seven_Pairs_Win) (int, string) {
	var yakuman_count int
	var msg string
	if sp.TenHo {
		if !sp.Tsumo || sp.SelfWind != 1 {
			log.Println("error for yakuman special")
			return 0, ""
		}
		yakuman_count++
		msg += "天和 役満"
	} else if sp.JiHo {
		if !sp.Tsumo || sp.SelfWind == 1 {
			log.Println("error for yakuman special")
			return 0, ""
		}
		yakuman_count++
		msg += "地和 役満"
	}
	return yakuman_count, msg
}

func OnlyZi(sp win.Seven_Pairs_Win) (int, string) {
	for i := range sp.PairList {
		if sp.PairList[i].Suit != 'z' {
			return 0, ""
		}
	}
	return 1, "字一色 役満\n"
}

func Yakuman_Check(sp win.Seven_Pairs_Win) (int, string) {
	var yakuman_count int
	var msg string
	yakumanChecks := []func(win.Seven_Pairs_Win) (int, string){
		Yakuman_Special,
		OnlyZi,
	}

	for _, checkFunc := range yakumanChecks {
		count, str := checkFunc(sp)
		yakuman_count += count
		msg += str
	}

	return yakuman_count, msg
}

func CalculateYaku(sp win.Seven_Pairs_Win) (int, string) {
	var han int = 2
	var msg string = "七対子 2飜\n"

	// Execute each function and accumulate the results
	funcs := []func(win.Seven_Pairs_Win) (int, string){
		Non_Yakuman_Special,
		MenchinTsumo,
		Tanyao,
		HonOldHead,
		Somete,
	}

	for _, f := range funcs {
		curHan, curMsg := f(sp)
		han += curHan
		msg += curMsg
	}

	dorahan, doramsg := CalculateDora(sp)
	han += dorahan
	msg += doramsg
	return han, msg
}
func CalculateDora(sp win.Seven_Pairs_Win) (int, string) {
	var han int
	var msg string

	tiles := win.ConvertSevenPairsToMap(sp)
	var dora_han int
	for i := range sp.MotedoraSuit {
		var dora_rank uint8
		if sp.MotedoraSuit[i] == 'z' {
			if sp.MotedoraRank[i] >= 5 {
				if sp.MotedoraRank[i] == 7 {
					dora_rank = 5
				} else {
					dora_rank = sp.MotedoraRank[i] + 1
				}
			} else {
				if sp.MotedoraRank[i] == 4 {
					dora_rank = 1
				} else {
					dora_rank = sp.MotedoraRank[i] + 1
				}
			}
		} else {
			if sp.MotedoraRank[i] == 9 {
				dora_rank = 1
			} else {
				dora_rank = sp.MotedoraRank[i] + 1
			}
		}
		dora_han += tiles[sp.MotedoraSuit[i]][dora_rank]
	}
	if dora_han != 0 {
		han += dora_han
		msg += fmt.Sprintf("ドラ %v飜\n", dora_han)
	}
	if sp.Akadora != 0 {
		han += sp.Akadora
		msg += fmt.Sprintf("赤ドラ %v飜\n", sp.Akadora)
	}
	var uradora_han int
	for i := range sp.MotedoraSuit {
		var dora_rank uint8
		if sp.MotedoraSuit[i] == 'z' {
			if sp.MotedoraRank[i] >= 5 {
				if sp.MotedoraRank[i] == 7 {
					dora_rank = 5
				} else {
					dora_rank = sp.MotedoraRank[i] + 1
				}
			} else {
				if sp.MotedoraRank[i] == 4 {
					dora_rank = 1
				} else {
					dora_rank = sp.MotedoraRank[i] + 1
				}
			}
		} else {
			if sp.MotedoraRank[i] == 9 {
				dora_rank = 1
			} else {
				dora_rank = sp.MotedoraRank[i] + 1
			}
		}
		uradora_han += tiles[sp.MotedoraSuit[i]][dora_rank]
	}
	if uradora_han != 0 {
		han += uradora_han
		msg += fmt.Sprintf("裏ドラ %v飜\n", uradora_han)
	}
	return han, msg
}
