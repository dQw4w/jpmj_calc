package main

import (
	//"jpmj_calc/test1"

	"jpmj_calc/hand"
	"log"
	//"strconv"
)

func main() {

	a, err := hand.ConvertStrToHand("011928m2223p055s12734")
	if err != nil {
		log.Println(err)
	}
	log.Println(hand.SortAndReturnAkadora(&a))
	log.Println(a)
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
