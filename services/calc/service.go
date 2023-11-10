package calc

import (
	"log"

	"github.com/dQw4w/jpmj_calc/services/calc/combination"
	"github.com/dQw4w/jpmj_calc/services/calc/hand"
	"github.com/dQw4w/jpmj_calc/services/calc/hu"
	"github.com/dQw4w/jpmj_calc/services/calc/score"
	"github.com/dQw4w/jpmj_calc/services/calc/win"
	"github.com/dQw4w/jpmj_calc/services/calc/yaku_common"
	"github.com/dQw4w/jpmj_calc/services/calc/yaku_sevenpair"
)

type CalculateService struct {
}

func NewCalculateService() *CalculateService {
	new := CalculateService{}
	return &new
}

func (s *CalculateService) Calculate(handstr string, akarehai string, furolist []string, tempwin win.Common_Win) string {
	output := ""
	//handstr := "12340678mps11122z" //TODO:enter the concealed tiles (discluding 暗槓, 和了牌) here
	a, err := hand.ConvertStrToHand(handstr)
	if err != nil {
		log.Panic(err)
	}
	a, _ = hand.AppendOne(a, akarehai) //TODO: enter 和了牌 (winning tile)
	if hand.Len(a)%3 != 2 {
		panic("no!")
	}
	//furolist := []string{ /*"333z", "444z"*/ } //TODO: add furo menzis here (including 暗槓, XnnXs is the format for it, n:rank,s:suit)
	//tempwin := win.CreateEmptyCommon()
	//TODO: modify info that can't be read from string here
	// tempwin.Reach = true
	// //tempwin.TenHo = true
	// //tempwin.JiHo = true
	tempwin.Menchin = true // this is always set to true by default, no need to modify
	// tempwin.Tsumo = true
	// tempwin.SelfWind = 1
	// tempwin.FieldWind = 1
	// //tempwin.Ippatsu = true
	// tempwin.MotedoraRank = append(tempwin.MotedoraRank, 4)
	// tempwin.MotedoraSuit = append(tempwin.MotedoraSuit, 'z')
	// tempwin.UradoraRank = append(tempwin.UradoraRank, 1)
	// tempwin.UradoraSuit = append(tempwin.UradoraSuit, 'z')
	//tempwin.RinShan = true
	//tempwin.ChanKan = true

	for i := range furolist {
		if furolist[i] == "" {
			continue
		}
		newmenzi, _, err1 := combination.ConvertStrToMenzi(furolist[i])
		if err1 == nil {
			tempwin, _ = win.AddMenzi(newmenzi, tempwin)
		}
		if newmenzi.Type != 'C' {
			tempwin.Menchin = false
		}
	}
	tempwin.Akadora = int(hand.SortAndReturnAkadora(&a))
	//log.Println(a.Win_Rank, a.Win_Suit)
	result, valid := win.CreateCommon(a, tempwin)
	if valid {
		maxyakuman := 0
		maxhan := -1
		resulthu := 0
		var maxmsg string
		var resultwin win.Common_Win
		// log.Print("result:")
		output += "result:\n"
		for i := range result {
			//log.Println(win.CommonString(result[i]))
			yakuman, msg := yaku_common.Yakuman_Check(result[i])
			if yakuman > 0 {
				if yakuman > maxyakuman {
					resultwin = result[i]
					maxyakuman = yakuman
					maxmsg = msg
				}
			} else if maxyakuman == 0 {
				han, msg2 := yaku_common.CalculateYaku(result[i])
				if han > maxhan {
					resultwin = result[i]
					maxhan = han
					maxmsg = msg2
					resulthu = hu.CalcHu(result[i])
				}
			}
		}
		// log.Println(win.CommonString(resultwin))
		output += win.CommonString(resultwin) + "\n"
		var oyaka, tsumo bool
		if resultwin.SelfWind == 1 {
			oyaka = true
		}
		tsumo = resultwin.Tsumo
		if maxyakuman > 0 {
			// log.Println(score.CalcYakumanPointsString(maxyakuman, oyaka, tsumo, maxmsg))
			output += score.CalcYakumanPointsString(maxyakuman, oyaka, tsumo, maxmsg) + "\n"
			return output
		} else if maxhan != 0 {

			if maxyakuman == 0 {

				dorahan, doramsg := yaku_common.CalculateDora(resultwin)
				maxhan += dorahan
				maxmsg += doramsg
			}
			log.Println(score.CalcPointsString(maxhan, resulthu, oyaka, tsumo, maxmsg))
			output += score.CalcPointsString(maxhan, resulthu, oyaka, tsumo, maxmsg) + "\n"
			return output
			// return
			//without using local yakus,
			//if a set of winning tiles both satisfy common_win and seven_pairs_win, than the former intepretation is always better
			//since 2peko is 3han and sevenpairs is 2han
			//also, that's the case unless you consider some obscure local yakus
		}

	}
	tempsevenpair := win.ConvertCommonToSeven(tempwin)
	result2, valid2 := win.CreateSevenPair(a, tempsevenpair)
	if valid2 {
		// log.Print("result:")
		output += "result:\n"
		// log.Println(win.SevenPairString(result2))
		output += win.SevenPairString(result2) + "\n"
		var oyaka, tsumo bool
		if result2.SelfWind == 1 {
			oyaka = true
		}
		tsumo = result2.Tsumo
		hu := 25
		yakuman, msg := yaku_sevenpair.Yakuman_Check(result2)
		if yakuman > 0 {
			// log.Println(score.CalcYakumanPointsString(yakuman, oyaka, tsumo, msg))
			output += score.CalcYakumanPointsString(yakuman, oyaka, tsumo, msg) + "\n"
		} else {
			han, msg2 := yaku_sevenpair.CalculateYaku(result2)
			// log.Println(score.CalcPointsString(han, hu, oyaka, tsumo, msg2))
			output += score.CalcPointsString(han, hu, oyaka, tsumo, msg2) + "\n"
		}
		return output

	}
	var tsumo, oyaka, tenho, jiho bool
	if tempwin.SelfWind == 1 {
		oyaka = true
	}
	tsumo = tempwin.Tsumo
	tenho = tempwin.TenHo
	jiho = tempwin.JiHo
	result3, valid3 := win.Create13Orphans(a, tsumo, oyaka, tenho, jiho)

	if valid3 {
		// log.Println("result:")
		output += "result:\n"
		yakuman := 1
		var msg string
		if tenho {
			msg += "天和 役満\n"
			yakuman++
		} else if jiho {
			msg += "地和 役満\n"
			yakuman++
		}
		if result3.Thirteen_Wait {
			yakuman++
			msg += "国士無双十三面待ち 2倍役満\n"
		} else {
			msg += "国士無双 役満\n"
		}

		// log.Println(score.CalcYakumanPointsString(yakuman, oyaka, tsumo, msg))
		output += score.CalcYakumanPointsString(yakuman, oyaka, tsumo, msg) + "\n"
		return output
	}
	// log.Println("Your input is invalid")
	return "Your input is invalid"

}

//some random testcode
/*
hi := "1234m"
	eee := make([]uint8, 5, 5)
	for i := range hi {
		eee[i] = hi[i] - '0'
	}
	log.Println(eee)
*/
/*for i := 0; i < 4-(hand.Len(a)-2)/3; i++ {
	var input string
	fmt.Scanf("%s", &input)
	//input = "X99Xp"
	if input == "" {
		i--
		continue
	}
	log.Println("read:", input)
	newmenzi, _, err1 := combination.ConvertStrToMenzi(input)
	if err1 != nil {
		log.Println("Invalid!")
		i--
	} else {
		tempwin, _ = win.AddMenzi(newmenzi, tempwin)
	}
}*/
