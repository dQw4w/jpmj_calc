package main

import (
	//"jpmj_calc/test1"

	"jpmj_calc/combination"
	"jpmj_calc/win"
	"log"
	//"strconv"
)

func main() {

	a := win.Common_Win{}
	//var err error
	stra, _ := combination.NewStraight('s', 1)
	err := win.AddMenzi(stra, &a)
	if err != nil {
		log.Println(err)
	}
	//a.MenziList[0], err = combination.NewStraight('s', 1)
	log.Printf("%s,%s,%b\n", string(a.MenziList[0].Type), string(a.MenziList[0].Suit), a.MenziList[0].Rank)
	//log.Println(a.MenziList[0])

	/*a, err := hand.ConvertStrToHand("011928m2223p055s12734z12")
	if err != nil {
		log.Println(err)
	}
	log.Println(hand.SortAndReturnAkadora(&a))
	log.Println(a)*/
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
