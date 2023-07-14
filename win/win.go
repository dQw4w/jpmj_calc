package win

import (
	"errors"
	"jpmj_calc/combination"
)

type Common_Win struct {
	MenziList [4]combination.Menzi
	Eye       combination.Pair
}

type Seven_Pairs_Win struct {
	PairList [7]combination.Pair
}
type Thirteen_Orphans struct {
	Repeat_Suit byte
	Repeat_Rank uint8
}

func AddMenzi(menzi combination.Menzi, common *Common_Win) error {
	for i := range common.MenziList {
		if combination.IsEmptyMenzi(common.MenziList[i]) {
			common.MenziList[i] = menzi
			return nil
		}
	}
	return errors.New("Error: Menzilist is already full")
}
func SetPair(pair combination.Pair, common *Common_Win) error {
	common.Eye = pair
	return nil
}
func AddPair(pair combination.Pair, sevenpairwin *Seven_Pairs_Win) error {
	for i := range sevenpairwin.PairList {
		if combination.IsEmptyPair(sevenpairwin.PairList[i]) {
			sevenpairwin.PairList[i] = pair
			return nil
		}
	}
	return errors.New("Error: Pairlist is already full")
}
