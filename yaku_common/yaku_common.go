package yaku_common

import (
	"jpmj_calc/combination"
	"jpmj_calc/win"
	"log"
)

func Non_Yakuman_Special(cw win.Common_Win) (int, string) {
	var han int
	var msg string
	if cw.DoubleReach {
		han += 2
		msg += "ダブル立直 2飜\n"
	} else if cw.Reach {
		han++
		msg += "立直 1飜\n"
	}
	if cw.Ippatsu {
		han++
		msg += "一発 1飜\n"
	}
	if cw.ChanKan {
		han++
		msg += "搶槓 1飜\n"
	}
	if cw.RinShan {
		han++
		msg += "嶺上開花 1飜\n"
	}
	if cw.HaiTei {
		han++
		msg += "海底摸月 1飜\n"
	}
	if cw.HoTei {
		han++
		msg += "河底撈魚 1飜\n"
	}
	return han, msg
}
func MenchinTsumo(cw win.Common_Win) (int, string) {
	if !(cw.Menchin && cw.Tsumo) {
		return 0, ""
	}
	return 1, "門前清自摸和 1飜\n"
}
func Tanyao(cw win.Common_Win) (int, string) {
	if cw.Eye.Suit == 'z' || cw.Eye.Rank == 1 || cw.Eye.Rank == 9 {
		return 0, ""
	}
	for i := range cw.MenziList {
		menzi := cw.MenziList[i]
		if menzi.Type == 'S' {
			if menzi.Rank == 1 || menzi.Rank == 7 {
				return 0, ""
			}
		} else {
			if menzi.Suit == 'z' || menzi.Rank == 1 || menzi.Rank == 9 {
				return 0, ""
			}
		}
	}
	return 1, "断么九　1飜\n"
}

func Yakuhai_Selfwind(cw win.Common_Win) (int, string) {
	for i := range cw.MenziList {
		if cw.MenziList[i].Suit == 'z' && cw.MenziList[i].Rank == cw.SelfWind {
			var msg string
			switch cw.SelfWind {
			case 1:
				msg = "役牌 自風東 1飜\n"
			case 2:
				msg = "役牌 自風南 1飜\n"
			case 3:
				msg = "役牌 自風西 1飜\n"
			case 4:
				msg = "役牌 自風北 1飜\n"
			default:
				return 0, ""
			}
			return 1, msg
		}
	}
	return 0, ""
}
func Yakuhai_Fieldwind(cw win.Common_Win) (int, string) {
	for i := range cw.MenziList {
		if cw.MenziList[i].Suit == 'z' && cw.MenziList[i].Rank == cw.FieldWind {
			var msg string
			switch cw.FieldWind {
			case 1:
				msg = "役牌 場風東 1飜\n"
			case 2:
				msg = "役牌 場風南 1飜\n"
			case 3:
				msg = "役牌 場風西 1飜\n"
			case 4:
				msg = "役牌 場風北 1飜\n"
			default:
				return 0, ""
			}
			return 1, msg
		}
	}
	return 0, ""
}
func Yakuhai_Sangen(cw win.Common_Win) (int, string) {
	var han int
	var msg string
	//var first bool = true
	for i := range cw.MenziList {
		menzi := cw.MenziList[i]
		if menzi.Suit == 'z' {
			switch menzi.Rank {
			case 5, 6, 7:
				han++

				switch menzi.Rank {
				case 5:
					msg += "役牌　白　1飜\n"

				case 6:
					msg += "役牌　發　1飜\n"

				case 7:
					msg += "役牌　中　1飜\n"
				}
			}
		}
	} //for end
	return han, msg
}

func Pinhu(cw win.Common_Win) (int, string) { // unfinished

	if !cw.Menchin {
		return 0, ""
	}
	for i := range cw.MenziList {
		if cw.MenziList[i].Type != 'S' {
			return 0, ""
		}
	}
	//TODO:if eye is yakuhai, false
	if cw.Eye.Suit == 'z' { //Eye is yakuhai
		if cw.Eye.Rank == cw.SelfWind || cw.Eye.Rank == cw.FieldWind || cw.Eye.Rank >= 5 {
			return 0, ""
		}
	}
	//TODO:聽牌型
	if cw.Win_Com_IDX == 4 { //單騎
		return 0, ""
	}
	if cw.Win_Tile_IDX == 1 { //坎張
		return 0, ""
	}
	//邊張
	if cw.Win_Tile_IDX == 0 && cw.MenziList[cw.Win_Com_IDX].Rank == 7 {
		return 0, ""
	}
	if cw.Win_Tile_IDX == 2 && cw.MenziList[cw.Win_Com_IDX].Rank == 1 {
		return 0, ""
	}
	return 1, "平和 1飜\n"
}
func OnePekoandTwoPeko(cw win.Common_Win) (int, string) {
	if !cw.Menchin {
		return 0, ""
	}
	var same_straights int
	var han int
	var msg string
	var used_j int
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			if j == used_j {
				continue
			}
			if cw.MenziList[i].Type == 'S' && combination.SameMenzi(cw.MenziList[i], cw.MenziList[j]) {
				same_straights++
				used_j = j
				break
			}
		}
	}
	switch same_straights {
	case 0:
		han = 0
		msg = ""
	case 1:
		han = 1
		msg = "一盃口 1飜\n"
	case 2:
		han = 3
		msg = "二盃口 3飜\n"
	default:
		log.Println("Something is wrong with peko")
	}
	return han, msg
}

//TODO:dora,akarora

// 2hanyaku
func ThreeSameTrp(cw win.Common_Win) (int, string) {
	//check if three trp up
	var nonzi_trp_count int
	for i := range cw.MenziList {
		if cw.MenziList[i].Type != 'S' && cw.MenziList[i].Suit != 'z' {
			nonzi_trp_count++
		}
	}
	if nonzi_trp_count < 3 {
		return 0, ""
	}
	for i := 0; i < 4; i++ {
		var indexs []int
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			indexs = append(indexs, j)

		}
		menzi1, menzi2, menzi3 := cw.MenziList[indexs[0]], cw.MenziList[indexs[1]], cw.MenziList[indexs[2]]

		if menzi1.Type == 'S' || menzi2.Type == 'S' || menzi3.Type == 'S' {
			continue
		}
		if menzi1.Suit == 'z' || menzi2.Suit == 'z' || menzi3.Suit == 'z' {
			continue
		}
		if menzi1.Rank == menzi2.Rank && menzi2.Rank == menzi3.Rank {
			return 2, "三色同刻 2飜\n"
		}
	}
	return 0, ""
	//check
}
func ThreeSameStra(cw win.Common_Win) (int, string) {
	//check if three trp up
	var stra_count int
	for i := range cw.MenziList {
		if cw.MenziList[i].Type == 'S' {
			stra_count++
		}
	}
	if stra_count < 3 {
		return 0, ""
	}
	for i := 0; i < 4; i++ {
		var indexs []int
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			indexs = append(indexs, j)
		}
		menzi1, menzi2, menzi3 := cw.MenziList[indexs[0]], cw.MenziList[indexs[1]], cw.MenziList[indexs[2]]

		if menzi1.Type != 'S' || menzi2.Type != 'S' || menzi3.Type != 'S' {
			continue
		}
		if menzi1.Suit == menzi2.Suit || menzi2.Suit == menzi3.Suit || menzi3.Suit == menzi1.Suit {
			continue
		}
		if menzi1.Rank == menzi2.Rank && menzi2.Rank == menzi3.Rank {
			if cw.Menchin {
				return 2, "三色同順 2飜\n"
			}
			return 1, "三色同順 1飜\n"
		}
	}
	return 0, ""
	//check
}
func ThreeKanzi(cw win.Common_Win) (int, string) {
	var kanzi_count int
	for i := range cw.MenziList {
		if cw.MenziList[i].Type == 'C' || cw.MenziList[i].Type == 'O' {
			kanzi_count++
		}
	}
	if kanzi_count == 3 {
		return 2, "三槓子 2飜\n"
	}
	return 0, ""
}
func Toitoi(cw win.Common_Win) (int, string) {
	for i := range cw.MenziList {
		if cw.MenziList[i].Type == 'S' {
			return 0, ""
		}
	}
	return 2, "対々和 2飜\n"
}
func ThreeConcealedTrp(cw win.Common_Win) (int, string) {
	var concealed_trp_count int
	for i := range cw.MenziList {
		if cw.MenziList[i].Type != 'S' && !cw.MenziList[i].Furo {
			concealed_trp_count++
		}
	}
	if concealed_trp_count == 3 {
		return 2, "三暗刻 2飜\n"
	}
	return 0, ""
}
func HonOldHead(cw win.Common_Win) (int, string) {
	if han, _ := Toitoi(cw); han != 2 {
		return 0, ""
	}
	var havezi bool = false
	for i := range cw.MenziList {
		if cw.MenziList[i].Suit == 'z' {
			havezi = true
		} else if cw.MenziList[i].Rank == 1 {
			continue
		} else if cw.MenziList[i].Rank == 9 {
			continue
		} else {
			return 0, ""
		}
	}
	if havezi {
		return 2, "混老頭 2飜\n"
	}
	return 0, ""
}

func SmallSangen(cw win.Common_Win) (int, string) {
	havesangen := [3]bool{false, false, false} // white,green,red
	if cw.Eye.Suit == 'z' && cw.Eye.Rank >= 5 {
		havesangen[cw.Eye.Rank-5] = true
	} else {
		return 0, ""
	}
	for i := range cw.MenziList {
		if cw.MenziList[i].Suit == 'z' && cw.MenziList[i].Rank >= 5 {
			havesangen[cw.MenziList[i].Rank-5] = true
		}
	}

	for i := 0; i < 3; i++ {
		if !havesangen[i] {
			return 0, ""
		}
	}
	return 2, "小三元  2飜\n"
	/*return 1, "大三元 役満\n"
	if cw.Eye.Suit != 'z' || cw.Eye.Rank < 5 {
		return 0, ""
	}
	var menzi_sangen_count int
	for i := range cw.MenziList {
		if cw.MenziList[i].Suit == 'z' && cw.MenziList[i].Rank >= 5 {
			menzi_sangen_count++
		}
	}*/
	/*if menzi_sangen_count == 2 {
		return 2, "小三元  2飜\n"
	}*/
	//return 0, ""
}
func ChanTa(cw win.Common_Win) (int, string) {
	if han, _ := HonOldHead(cw); han == 2 {
		return 0, ""
	}
	var havezi bool = false
	for i := range cw.MenziList {
		if cw.MenziList[i].Suit == 'z' {
			havezi = true
			continue
		} else if cw.MenziList[i].Type == 'S' {
			if cw.MenziList[i].Rank == 1 || cw.MenziList[i].Rank == 7 {
				continue
			}
			return 0, ""
		} else {
			if cw.MenziList[i].Rank == 1 || cw.MenziList[i].Rank == 9 {
				continue
			}
			return 0, ""
		}
	}
	if havezi {
		if cw.Menchin {
			return 2, "混全帯么九 2飜\n"
		} else {
			return 1, "混全帯么九 1飜\n"
		}
	} else {
		if cw.Menchin {
			return 3, "純全帯么九 3飜"
		} else {
			return 2, "純全帯么九 2飜"
		}
	}
}
func OneDragon(cw win.Common_Win) (int, string) {
	var stra_count int
	for i := range cw.MenziList {
		if cw.MenziList[i].Type == 'S' {
			stra_count++
		}
	}
	if stra_count < 3 {
		return 0, ""
	}
	for i := 0; i < 4; i++ {
		var indexs []int
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			indexs = append(indexs, j)
		}
		menzi1, menzi2, menzi3 := cw.MenziList[indexs[0]], cw.MenziList[indexs[1]], cw.MenziList[indexs[2]]

		if menzi1.Type != 'S' || menzi2.Type != 'S' || menzi3.Type != 'S' {
			continue
		}
		if menzi1.Suit != menzi2.Suit || menzi2.Suit != menzi3.Suit || menzi3.Suit != menzi1.Suit {
			continue
		}
		if menzi1.Rank == menzi2.Rank || menzi2.Rank == menzi3.Rank || menzi3.Rank == menzi1.Rank {
			continue
		}
		if menzi1.Rank%3 == 0 && menzi2.Rank%3 == 0 && menzi3.Rank%3 == 0 {
			if menzi1.Rank+menzi2.Rank+menzi3.Rank == 12 {
				if cw.Menchin {
					return 2, "一気通貫 2飜\n"
				}
				return 1, "一気通貫 1飜\n"
			}
		}
	}
	return 0, ""
}
func Somete(cw win.Common_Win) (int, string) {
	var havezi bool = false
	var suit byte = 'x'
	for i := range cw.MenziList {
		if cw.MenziList[i].Suit == 'z' {
			havezi = true
		} else {
			if suit == 'x' {
				suit = cw.MenziList[i].Suit
			} else if suit != cw.MenziList[i].Suit {
				return 0, ""
			}
		}
	}
	if suit == 'x' {
		return 0, ""
	}
	if havezi {
		if cw.Menchin {
			return 3, "混一色 3飜\n"
		} else {
			return 2, "混一色 2飜\n"
		}
	}
	if cw.Menchin {
		return 6, "清一色 6飜\n"
	} else {
		return 5, "清一色 5飜\n"
	}
}

// below are yakumans, the return integer value represents how many times of yakuman
// TODO: tenho jiho
func Yakuman_Special(cw win.Common_Win) (int, string) {
	var yakuman_count int
	var msg string
	if cw.TenHo {
		if !cw.Tsumo || cw.SelfWind != 1 {
			log.Println("error for yakuman special")
			return 0, ""
		}
		yakuman_count++
		msg += "天和 役満"
	} else if cw.JiHo {
		if !cw.Tsumo || cw.SelfWind == 1 {
			log.Println("error for yakuman special")
			return 0, ""
		}
		yakuman_count++
		msg += "地和 役満"
	}
	return yakuman_count, msg
}
func BigSangen(cw win.Common_Win) (int, string) { // TODO: also modify small ,using similar logic
	havesangen := [3]bool{false, false, false} // white,green,red

	for i := range cw.MenziList {
		if cw.MenziList[i].Suit == 'z' && cw.MenziList[i].Rank >= 5 {
			havesangen[cw.MenziList[i].Rank-5] = true
		}
	}

	for i := 0; i < 3; i++ {
		if !havesangen[i] {
			return 0, ""
		}
	}
	return 1, "大三元 役満\n"
}
func FourConcealedTrp(cw win.Common_Win) (int, string) {
	for i := range cw.MenziList {
		if cw.MenziList[i].Type == 'S' {
			return 0, ""
		}
		if cw.MenziList[i].Furo {
			return 0, ""
		}
	}
	if cw.Win_Com_IDX == 4 { //pair
		return 2, "四暗刻単騎 二倍役満\n"
	}
	return 1, "四暗刻 役満\n"
}
func OnlyZi(cw win.Common_Win) (int, string) {
	if cw.Eye.Suit != 'z' {
		return 0, ""
	}
	for i := range cw.MenziList {
		if cw.MenziList[i].Suit != 'z' {
			return 0, ""
		}
	}
	return 1, "字一色 役満\n"
}
func OnlyGreen(cw win.Common_Win) (int, string) {
	if cw.Eye.Suit == 'z' {
		if cw.Eye.Rank != 6 {
			return 0, ""
		}
	} else if cw.Eye.Suit == 's' {
		if cw.Eye.Rank != 3 && cw.Eye.Rank%2 == 1 {
			return 0, ""
		}
	} else {
		return 0, ""
	}
	for i := range cw.MenziList {
		menzi := cw.MenziList[i]
		if menzi.Suit == 'z' {
			if menzi.Rank == 6 {
				continue
			}
		} else if menzi.Suit == 's' {
			if menzi.Rank == 3 || menzi.Rank%2 == 0 {
				continue
			}
		}
		return 0, ""
	}
	return 1, "緑一色 役満\n"
}
func OnlyOld(cw win.Common_Win) (int, string) {
	if !(cw.Eye.Suit != 'z' && (cw.Eye.Rank == 1 || cw.Eye.Rank == 9)) {
		return 0, ""
	}
	for i := range cw.MenziList {
		menzi := cw.MenziList[i]
		if menzi.Suit == 'z' || menzi.Type == 'S' {
			return 0, ""
		} else if menzi.Rank != 1 && menzi.Rank != 9 {
			return 0, ""
		}
	}
	return 1, "清老頭 役満\n"
}
func FourWinds(cw win.Common_Win) (int, string) {
	havewind := [4]bool{false, false, false, false} // white,green,red
	var isEye bool = false
	for i := range cw.MenziList {
		if cw.MenziList[i].Suit == 'z' && cw.MenziList[i].Rank <= 4 {
			havewind[cw.MenziList[i].Rank-1] = true
		}
	}
	if cw.Eye.Suit == 'z' && cw.Eye.Rank <= 4 {
		isEye = true
		havewind[cw.Eye.Rank-1] = true
	}
	for i := 0; i < 4; i++ {
		if !havewind[i] {
			return 0, ""
		}
	}
	if isEye {
		return 1, "小四喜 役満\n"
	}
	return 2, "大四喜 二倍役満"
}
func FourKanzi(cw win.Common_Win) (int, string) {
	var kanzi_count int
	for i := range cw.MenziList {
		if cw.MenziList[i].Type == 'C' || cw.MenziList[i].Type == 'O' {
			kanzi_count++
		}
	}
	if kanzi_count == 4 {
		return 1, "四槓子 役満\n"
	}
	return 0, ""
}
func NineGates(cw win.Common_Win) (int, string) {
	if !cw.Menchin {
		return 0, ""
	}
	if han, _ := Somete(cw); han < 5 { //one suit check
		return 0, ""
	}
	tile_count := make(map[uint8]int)
	for i := range cw.MenziList {
		switch cw.MenziList[i].Type {
		case 'C', 'O':
			return 0, ""
		case 'S':
			for j := cw.MenziList[i].Rank; j < cw.MenziList[i].Rank+3; j++ {
				tile_count[j]++
			}

		case 'T':
			tile_count[cw.MenziList[i].Rank] += 3

		default:
			return 0, ""
		}
		for i := 1; i <= 9; i++ {
			if i == 1 || i == 9 {
				if tile_count[uint8(i)] < 3 {
					return 0, ""
				}
			} else {
				if tile_count[uint8(i)] < 1 {
					return 0, ""
				}
			}
		}
	}
	_, winrank := win.GetWinningTile(cw)
	tile_count[winrank]--
	var pure bool = false
	if winrank == 1 || winrank == 9 {
		if tile_count[winrank] == 3 {
			pure = true
		}
	} else {
		if tile_count[winrank] == 1 {
			pure = true
		}
	}
	if pure {
		return 2, "純正九蓮宝燈 二倍役満\n"
	}
	return 1, "九蓮宝燈 役満\n"
}

func Yakuman_Check(cw win.Common_Win) (int, string) {
	var yakuman_count int
	var msg string
	yakumanChecks := []func(cw win.Common_Win) (int, string){
		Yakuman_Special,
		BigSangen,
		FourConcealedTrp,
		OnlyZi,
		OnlyGreen,
		OnlyOld,
		FourWinds,
		FourKanzi,
		NineGates,
	}

	for _, checkFunc := range yakumanChecks {
		count, str := checkFunc(cw)
		yakuman_count += count
		msg += str
	}

	return yakuman_count, msg
}
func CalculateYaku(cw win.Common_Win) (int, string) { // todo
	var han int
	var msg string

	// Execute each function and accumulate the results
	funcs := []func(win.Common_Win) (int, string){
		Non_Yakuman_Special,
		MenchinTsumo,
		Tanyao,
		Yakuhai_Selfwind,
		Yakuhai_Fieldwind,
		Yakuhai_Sangen,
		Pinhu,
		OnePekoandTwoPeko,
		ThreeSameTrp,
		ThreeSameStra,
		ThreeKanzi,
		Toitoi,
		ThreeConcealedTrp,
		HonOldHead,
		SmallSangen,
		ChanTa,
		OneDragon,
		Somete,
	}

	for _, f := range funcs {
		curHan, curMsg := f(cw)
		han += curHan
		msg += curMsg
	}

	return han, msg
}
