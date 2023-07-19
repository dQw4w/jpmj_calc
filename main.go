package main

import (
	//"jpmj_calc/test1"

	"fmt"
	"jpmj_calc/combination"
	"jpmj_calc/hand"
	"jpmj_calc/hu"
	"jpmj_calc/score"
	"jpmj_calc/win"
	"jpmj_calc/yaku_common"
	"log"
	//"strconv"
)

/*
	func rm(a []int, i int) {
		a = append(a[:i], a[i+1:]...)
		log.Println(a)
		return
	}
*/
func main() {

	/*d := win.CreateEmptyCommon()
	fmt.Println(win.HaveEye(d))
	fmt.Println((win.Menzi_Count(d)))
	return*/
	/*c := win.CreateEmptyCommon()
	log.Println(c)
	return
	a := win.Common_Win{}
	//var err error
	stra, _ := combination.NewStraight('s', 1)
	strb, _ := combination.NewStraight('m', 1)
	strc, _ := combination.NewStraight('p', 1)
	trp1, _ := combination.NewTriplet('z', 1)
	prr, _ := combination.NewPair('z', 2)

	var err error
	a, err = win.AddMenzi(stra, a)
	if err != nil {
		log.Println(err)
	}
	a, err = win.AddMenzi(strb, a)
	if err != nil {
		log.Println(err)
	}
	a, err = win.AddMenzi(strc, a)
	if err != nil {
		log.Println(err)
	}
	a, err = win.AddMenzi(trp1, a)
	if err != nil {
		log.Println(err)
	}
	a, err = win.SetPair(prr, a)
	if err != nil {
		log.Println(err)
	}
	//a.MenziList[0], err = combination.NewStraight('s', 1)

	b, _ := win.CopyCommon(a)
	b.MenziList[0].Rank = 5
	for i := range a.MenziList {
		log.Printf("%s,%s,%v\n", string(a.MenziList[i].Type), string(a.MenziList[i].Suit), a.MenziList[i].Rank)
	}
	log.Printf("%s,%v\n", string(a.Eye.Suit), a.Eye.Rank)
	for i := range a.MenziList {
		log.Printf("%s,%s,%v\n", string(b.MenziList[i].Type), string(b.MenziList[i].Suit), b.MenziList[i].Rank)
	}
	log.Printf("%s,%v\n", string(b.Eye.Suit), b.Eye.Rank)
	//log.Println(a.MenziList[0])*/

	//a, err := hand.ConvertStrToHand("12122m460ps222333z")
	handstr := "99m11199p789s222z"

	a, err := hand.ConvertStrToHand(handstr /*"19m19p119s123456z"*/) //test different inputs here
	if err != nil {
		log.Panic(err)

	}
	a, _ = hand.AppendOne(a, "9p")
	if hand.Len(a)%3 != 2 {
		panic("no!")
	}

	tempwin := win.CreateEmptyCommon()
	tempwin.Menchin = true
	tempwin.Reach = true
	tempwin.Tsumo = true
	tempwin.SelfWind = 2
	tempwin.FieldWind = 1
	//tempwin.Ippatsu = true
	tempwin.Motedora_rank = append(tempwin.Motedora_rank, 5)
	tempwin.Motedora_suit = append(tempwin.Motedora_suit, 'p')
	tempwin.Uradora_rank = append(tempwin.Uradora_rank, 5)
	tempwin.Uradora_suit = append(tempwin.Uradora_suit, 's')
	tempwin.Akadora = 1
	tempwin.RinShan = true
	//tempwin.ChanKan = true
	//log.Println((hand.Len(a) - 2) / 3)
	for i := 0; i < 4-(hand.Len(a)-2)/3; i++ {
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
	}

	log.Println(hand.SortAndReturnAkadora(&a))
	log.Println(a)
	log.Println("tempwin:", tempwin)
	result, valid := win.CreateCommon(a, tempwin)
	log.Println("result plain:", result)
	if valid {
		maxyakuman := 0
		maxhan := -1
		resulthu := 0
		var maxmsg string
		var hanmsg string
		var resultwin win.Common_Win
		log.Print("result:")
		for i := range result {
			log.Println(win.CommonString(result[i]))
			//log.Printf("HaveEye:%t\nMenziCount:%v\n", win.HaveEye(result[i]), win.Menzi_Count(result[i]))
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
		if maxhan != 0 {

			dorahan, doramsg := yaku_common.CalculateDora(resultwin)
			maxhan += dorahan
			maxmsg += doramsg
			log.Println(maxmsg)
			if maxhan >= 13 {
				hanmsg = "数え役満"
			} else if maxhan >= 11 {
				hanmsg = "三倍満"
			} else if maxhan >= 8 {
				hanmsg = "倍満"
			} else if maxhan >= 6 {
				hanmsg = "跳満"
			} else if maxhan >= 5 {
				hanmsg = "満貫"
			} else if maxhan >= 4 && resulthu >= 40 {
				hanmsg = "満貫"
			} else if maxhan >= 3 && resulthu >= 70 {
				hanmsg = "満貫"
			}
			if maxyakuman > 0 {
				log.Printf("%v倍役満\n", maxyakuman)
			} else {
				log.Printf("%v飜%v符 %s\n", maxhan, resulthu, hanmsg)
			}
			oya := false
			if tempwin.SelfWind == 1 {
				oya = true
			}
			if tempwin.Tsumo {
				oyapt, kopt := score.TsumoPoints(maxhan, resulthu, oya)
				if oya {
					log.Printf("%v点 オール\n", kopt)
				} else {
					log.Printf("(%v,%v)点\n", kopt, oyapt)
				}
			} else {
				log.Printf("%v点", score.RonPoints(maxhan, resulthu, oya))
			}
			return
		} else {
			log.Println("false for common")
		}
	}

	result2, valid2 := win.CreateSevenPair(a)
	if valid2 {
		log.Print("result:")

		log.Println(win.SevenPairString(result2))

		//combination.MenziS
	} else {
		log.Println("false for 7 pairs")
	}
	result3, valid3 := win.Create13Orphans(a)
	if valid3 {
		log.Println("result:", result3)
	} else {
		log.Println("false for 13 orphans")
	}
	/*b, err := combination.NewStraight('s', 6)
	if err != nil {
		log.Println(err)
		//log.Panic(err)
	}
	log.Println(b.Rank)*/
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
