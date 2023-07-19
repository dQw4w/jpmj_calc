package win

import (
	"errors"
	"jpmj_calc/combination"
	"jpmj_calc/hand"
	"log"
)

type Common_Win struct {
	MenziList    [4]combination.Menzi
	Eye          combination.Pair
	Win_Com_IDX  int // 4 represents eye
	Win_Tile_IDX int
	Tsumo        bool
	Menchin      bool
	SelfWind     uint8 //1234 ESWN
	FieldWind    uint8 //1234

	//special
	Reach       bool
	DoubleReach bool // Reach and Double Reach only one is true
	ChanKan     bool
	RinShan     bool
	HaiTei      bool
	HoTei       bool
	Ippatsu     bool

	//yakuman special
	TenHo bool
	JiHo  bool

	Akadora       int
	Motedora_suit byte
	Motedora_rank uint8

	Uradora_suit byte
	Uradora_rank uint8
}

type Seven_Pairs_Win struct {
	PairList [7]combination.Pair
}
type Thirteen_Orphans struct {
	Repeat_Suit byte
	Repeat_Rank uint8
}

func CreateEmptyCommon() Common_Win {
	New := Common_Win{}
	//New.MenziList = [4]combination.Menzi{}
	return New
}
func AddMenzi(menzi combination.Menzi, common Common_Win) (Common_Win, error) {
	for i := range common.MenziList {
		if combination.IsEmptyMenzi(common.MenziList[i]) {
			common.MenziList[i] = menzi
			return common, nil
		}
	}
	return common, errors.New("Error: Menzilist is already full")
}
func SetPair(pair combination.Pair, common Common_Win) (Common_Win, error) {
	common.Eye = pair
	return common, nil
}
func HaveEye(common Common_Win) bool {
	return !combination.IsEmptyPair(common.Eye)
}
func Menzi_Count(common Common_Win) int {
	count := 0
	for i := range common.MenziList {
		if combination.IsEmptyMenzi(common.MenziList[i]) {
			break
		}
		count++
	}
	return count
}
func AddPair(pair combination.Pair, sevenpairwin Seven_Pairs_Win) (Seven_Pairs_Win, error) {
	for i := range sevenpairwin.PairList {
		if combination.IsEmptyPair(sevenpairwin.PairList[i]) {
			sevenpairwin.PairList[i] = pair
			return sevenpairwin, nil
		}
	}
	return sevenpairwin, errors.New("Error: Pairlist is already full")
}

type NewHandAndRemovedMenzis struct {
	NewHand       hand.Hand
	RemovedMenzis []combination.Menzi
}

// TODO: set the menzi including the win tile to furo
func CreateCommon(hd hand.Hand, cw Common_Win) ([]Common_Win, bool) {
	var result []Common_Win
	if !hand.IsValidHandNum(hd) {
		//log.Println("hd len is invalid")
		return result, false
	}
	//menzi_count := Menzi_Count(cw)

	for i := 0; i < hand.Len(hd); i++ { // checks from the first tile to the second-to-last
		newhd, pair, valid := RemovePair(hd, i)
		if valid {
			//log.Println(newhd)
			//log.Println(pair)

			NewWin := Common_Win{}
			NewWin, _ = CopyCommon(cw)
			log.Println("Newwin:", NewWin)
			NewWin, _ = SetPair(pair, NewWin)
			possibles := []NewHandAndRemovedMenzis{NewHandAndRemovedMenzis{NewHand: newhd}}
			//log.Println(NewWin)
			/*hel, _ := combination.NewStraight('s', 1)
			possibles[0].RemovedMenzis = append(possibles[0].RemovedMenzis, hel)
			log.Println("hi")
			log.Println(possibles)*/
			for i := 0; i < 4-Menzi_Count(cw); i++ {
				cases := len(possibles)
				for j := 0; j < cases; j++ {
					//log.Println("case to remove:", possibles[0])
					newPossibilities, _ := RemoveMenzi(possibles[0])
					possibles = append(possibles[1:], newPossibilities...)
				}
			}
			//log.Println("possibles:", possibles)

			for i := range possibles {
				NewNewWin, _ := CopyCommon(NewWin)
				//log.Println("i:", i)
				for j := range possibles[i].RemovedMenzis {
					NewNewWin, _ = AddMenzi(possibles[i].RemovedMenzis[j], NewNewWin)
				}
				result = append(result, NewNewWin)
			}

		}
	}
	canCreate := false
	if len(result) != 0 {
		canCreate = true
	}
	return result, canCreate
}

func CopyCommon(src Common_Win) (Common_Win, error) {
	dst := Common_Win{}
	dst.MenziList = [4]combination.Menzi{}
	var err2 error
	dst.Eye, _ = combination.NewPair(src.Eye.Suit, src.Eye.Rank, src.Eye.Furo)
	dst.Tsumo = src.Tsumo
	dst.Menchin = src.Menchin
	//log.Println(dst.Eye)
	/*if err1 != nil {
		//panic(err1)
		return Common_Win{}, err1
	}*/
	for i := range dst.MenziList {
		dst.MenziList[i], err2 = combination.New_Menzi(src.MenziList[i].Type, src.MenziList[i].Suit, src.MenziList[i].Rank, src.MenziList[i].Furo)
		if err2 != nil {
			//panic(err2)
			return dst, err2
		}
	}
	//log.Print("hi")
	//log.Println(dst)
	return dst, nil
}
func RemovePair(hd hand.Hand, idx int) (hand.Hand, combination.Pair, bool) {
	var mLen, pLen, sLen, zLen = hand.AllLen(hd)
	//log.Print("hd before rmpair func:")
	//log.Println(hd)
	//log.Printf("idx=%v\n", idx)
	if idx < mLen {
		if len(hd.Man)%3 != 2 {
			return hand.Hand{}, combination.Pair{}, false
		}
		if mLen > idx+1 && hd.Man[idx] == hd.Man[idx+1] {
			pair, err := combination.NewPair('m', hd.Man[idx], false)

			//log.Printf("newpairm = %v\n", hd.Man[idx])
			if err == nil {
				newhd := hand.Copy(hd)
				//fmt.Printf("newhdbefore:%v\n", newhd)
				//newhd.Man = make([]uint8, 0, 0)
				newhd.Man = append(newhd.Man[:idx], newhd.Man[idx+2:]...) //remove elements of index = idx, idx+1
				//log.Print("hd after rmpair func:")
				//log.Println(hd)
				return newhd, pair, true
			}
		}
	} else if idx < mLen+pLen {
		if len(hd.Pin)%3 != 2 {
			return hand.Hand{}, combination.Pair{}, false
		}
		pIdx := idx - mLen // Adjusted index within the p section
		//log.Printf("pidx=%v\n", pIdx)
		if pLen > pIdx+1 && hd.Pin[pIdx] == hd.Pin[pIdx+1] {
			pair, err := combination.NewPair('p', hd.Pin[pIdx], false)
			//log.Printf("newpairp = %v\n", hd.Man[pIdx])

			if err == nil {
				newhd := hand.Copy(hd)
				newhd.Pin = append(newhd.Pin[:pIdx], newhd.Pin[pIdx+2:]...) // Remove elements of index pIdx and pIdx+1
				return newhd, pair, true
			}
		}
	} else if idx < mLen+pLen+sLen {
		if len(hd.Sou)%3 != 2 {
			return hand.Hand{}, combination.Pair{}, false
		}
		sIdx := idx - (mLen + pLen) // Adjusted index within the s section
		//log.Printf("sidx=%v\n", sIdx)

		if sLen > sIdx+1 && hd.Sou[sIdx] == hd.Sou[sIdx+1] {
			pair, err := combination.NewPair('s', hd.Sou[sIdx], false)
			//log.Printf("newpairs = %v\n", hd.Man[sIdx])
			if err == nil {
				newhd := hand.Copy(hd)
				newhd.Sou = append(newhd.Sou[:sIdx], newhd.Sou[sIdx+2:]...) // Remove elements of index sIdx and sIdx+1
				return newhd, pair, true
			}
		}
	} else if idx < mLen+pLen+sLen+zLen {
		if len(hd.Zi)%3 != 2 {
			return hand.Hand{}, combination.Pair{}, false
		}

		zIdx := idx - (mLen + pLen + sLen) // Adjusted index within the z section
		//log.Printf("zidx=%v\n", zIdx)

		if zLen > zIdx+1 && hd.Zi[zIdx] == hd.Zi[int(zIdx+1)] {
			pair, err := combination.NewPair('z', hd.Zi[zIdx], false)
			//log.Printf("newpairz = %v\n", hd.Zi[zIdx])

			if err == nil {
				newhd := hand.Copy(hd)
				newhd.Zi = append(newhd.Zi[:zIdx], newhd.Zi[zIdx+2:]...) // Remove elements of index zIdx and zIdx+1
				return newhd, pair, true
			}
		}
	}
	//log.Println("Invalid input for RemovePair func")
	return hand.Hand{}, combination.Pair{}, false

}
func RemoveElements(slice []uint8, a, b, c int) []uint8 {
	newslice := []uint8{}
	for i := range slice {
		if i == a || i == b || i == c {
			continue
		}
		newslice = append(newslice, slice[i])
	}
	return newslice
}
func RemoveMenzi(hdmz NewHandAndRemovedMenzis) ([]NewHandAndRemovedMenzis, bool) {
	hd := hdmz.NewHand
	mzlist := hdmz.RemovedMenzis
	var possibleReturns []NewHandAndRemovedMenzis

	if len(hd.Man) != 0 {
		if hd.Man[0] == hd.Man[1] && hd.Man[1] == hd.Man[2] {
			trp, err := combination.NewTriplet('m', hd.Man[0], false)
			if err == nil {
				newhd := hand.Copy(hd)
				newhd.Man = newhd.Man[3:] // Remove elements of index = 0, 1, 2
				newMenziList := make([]combination.Menzi, len(mzlist))
				copy(newMenziList, mzlist)
				newMenziList = append(newMenziList, trp)
				possibleReturns = append(possibleReturns, NewHandAndRemovedMenzis{NewHand: newhd, RemovedMenzis: newMenziList})
			}
		}
		var idx1, idx2 int
		for i := range hd.Man {
			if hd.Man[i]-hd.Man[0] == 1 {
				idx1 = i
			} else if hd.Man[i]-hd.Man[0] == 2 {
				idx2 = i
				break
			}
		}

		if idx1 != 0 && idx2 != 0 {
			stra, err := combination.NewStraight('m', hd.Man[0], false)
			if err == nil {
				//log.Println("yay!")
				newhd := hand.Copy(hd)
				newhd.Man = RemoveElements(newhd.Man, 0, idx1, idx2) // Remove elements of index = 0, 1, 2
				newMenziList := make([]combination.Menzi, len(mzlist))
				copy(newMenziList, mzlist)
				newMenziList = append(newMenziList, stra)
				possibleReturns = append(possibleReturns, NewHandAndRemovedMenzis{NewHand: newhd, RemovedMenzis: newMenziList})
			}
		}
		//TODO: finish Pin,Sou,Zi accoring to Man (the straight part of the following code is wrong)
	} else if len(hd.Pin) != 0 {
		if hd.Pin[0] == hd.Pin[1] && hd.Pin[1] == hd.Pin[2] {
			trp, err := combination.NewTriplet('p', hd.Pin[0], false)
			if err == nil {
				newhd := hand.Copy(hd)
				newhd.Pin = newhd.Pin[3:] // Remove elements of index = 0, 1, 2
				newMenziList := make([]combination.Menzi, len(mzlist))
				copy(newMenziList, mzlist)
				newMenziList = append(newMenziList, trp)
				possibleReturns = append(possibleReturns, NewHandAndRemovedMenzis{NewHand: newhd, RemovedMenzis: newMenziList})
			}
		}

		var idx1, idx2 int
		for i := range hd.Pin {
			if hd.Pin[i]-hd.Pin[0] == 1 {
				idx1 = i
			} else if hd.Pin[i]-hd.Pin[0] == 2 {
				idx2 = i
				break
			}
		}

		if idx1 != 0 && idx2 != 0 {
			stra, err := combination.NewStraight('p', hd.Pin[0], false)
			if err == nil {
				newhd := hand.Copy(hd)
				newhd.Pin = RemoveElements(newhd.Pin, 0, idx1, idx2) // Remove elements of index = 0, 1, 2
				newMenziList := make([]combination.Menzi, len(mzlist))
				copy(newMenziList, mzlist)
				newMenziList = append(newMenziList, stra)
				possibleReturns = append(possibleReturns, NewHandAndRemovedMenzis{NewHand: newhd, RemovedMenzis: newMenziList})
			}
		}
	} else if len(hd.Sou) != 0 {
		if hd.Sou[0] == hd.Sou[1] && hd.Sou[1] == hd.Sou[2] {
			trp, err := combination.NewTriplet('s', hd.Sou[0], false)
			if err == nil {
				newhd := hand.Copy(hd)
				newhd.Sou = newhd.Sou[3:] // Remove elements of index = 0, 1, 2
				newMenziList := make([]combination.Menzi, len(mzlist))
				copy(newMenziList, mzlist)
				newMenziList = append(newMenziList, trp)
				possibleReturns = append(possibleReturns, NewHandAndRemovedMenzis{NewHand: newhd, RemovedMenzis: newMenziList})
			}
		}

		var idx1, idx2 int
		for i := range hd.Sou {
			if hd.Sou[i]-hd.Sou[0] == 1 {
				idx1 = i
			} else if hd.Sou[i]-hd.Sou[0] == 2 {
				idx2 = i
				break
			}
		}

		if idx1 != 0 && idx2 != 0 {
			stra, err := combination.NewStraight('s', hd.Sou[0], false)
			if err == nil {
				newhd := hand.Copy(hd)
				newhd.Sou = RemoveElements(newhd.Sou, 0, idx1, idx2) // Remove elements of index = 0, 1, 2
				newMenziList := make([]combination.Menzi, len(mzlist))
				copy(newMenziList, mzlist)
				newMenziList = append(newMenziList, stra)
				possibleReturns = append(possibleReturns, NewHandAndRemovedMenzis{NewHand: newhd, RemovedMenzis: newMenziList})
			}
		}
	} else if len(hd.Zi) != 0 {
		if hd.Zi[0] == hd.Zi[1] && hd.Zi[1] == hd.Zi[2] {
			trp, err := combination.NewTriplet('z', hd.Zi[0], false)
			if err == nil {
				newhd := hand.Copy(hd)
				newhd.Zi = newhd.Zi[3:] // Remove elements of index = 0, 1, 2
				newMenziList := make([]combination.Menzi, len(mzlist))
				copy(newMenziList, mzlist)
				newMenziList = append(newMenziList, trp)
				possibleReturns = append(possibleReturns, NewHandAndRemovedMenzis{NewHand: newhd, RemovedMenzis: newMenziList})
			}
		}
	}

	canRemove := false
	if len(possibleReturns) != 0 {
		canRemove = true
		//log.Println(possibleReturns)
	}

	return possibleReturns, canRemove
}
func is3dup(hd hand.Hand) bool {
	for i := 0; i < len(hd.Man)-2; i++ {
		if hd.Man[i] == hd.Man[i+1] && hd.Man[i+1] == hd.Man[i+2] {
			return true
		}
	}
	for i := 0; i < len(hd.Pin)-2; i++ {
		if hd.Pin[i] == hd.Pin[i+1] && hd.Pin[i+1] == hd.Pin[i+2] {
			return true
		}
	}

	for i := 0; i < len(hd.Sou)-2; i++ {
		if hd.Sou[i] == hd.Sou[i+1] && hd.Sou[i+1] == hd.Sou[i+2] {
			return true
		}
	}

	for i := 0; i < len(hd.Zi)-2; i++ {
		if hd.Zi[i] == hd.Zi[i+1] && hd.Zi[i+1] == hd.Zi[i+2] {
			return true
		}
	}
	return false

}
func CreateSevenPair(hd hand.Hand) (Seven_Pairs_Win, bool) {

	var result Seven_Pairs_Win
	if is3dup(hd) {
		return Seven_Pairs_Win{}, false
	}
	for i := 0; i < len(hd.Man); i += 2 {
		if hd.Man[i] == hd.Man[i+1] {
			pair, _ := combination.NewPair('m', hd.Man[i], false)
			result, _ = AddPair(pair, result)
		} else {
			return Seven_Pairs_Win{}, false
		}
	}
	for i := 0; i < len(hd.Pin); i += 2 {
		if hd.Pin[i] == hd.Pin[i+1] {
			pair, _ := combination.NewPair('p', hd.Pin[i], false)
			result, _ = AddPair(pair, result)
		} else {
			return Seven_Pairs_Win{}, false
		}
	}

	for i := 0; i < len(hd.Sou); i += 2 {
		if hd.Sou[i] == hd.Sou[i+1] {
			pair, _ := combination.NewPair('s', hd.Sou[i], false)
			result, _ = AddPair(pair, result)
		} else {
			return Seven_Pairs_Win{}, false
		}
	}

	for i := 0; i < len(hd.Zi); i += 2 {
		if hd.Zi[i] == hd.Zi[i+1] {
			pair, _ := combination.NewPair('z', hd.Zi[i], false)
			result, _ = AddPair(pair, result)
		} else {
			return Seven_Pairs_Win{}, false
		}
	}

	return result, true

}
func OnlyYao_Slice(tiles []uint8) bool {
	for i := range tiles {
		if tiles[i] != 1 && tiles[i] != 9 {
			return false
		}
	}
	return true
}
func Create13Orphans(hd hand.Hand) (Thirteen_Orphans, bool) {
	if hand.Len(hd) != 14 {
		return Thirteen_Orphans{}, false
	}
	var result Thirteen_Orphans
	if !(OnlyYao_Slice(hd.Man) && OnlyYao_Slice(hd.Pin) && OnlyYao_Slice(hd.Sou)) {
		return Thirteen_Orphans{}, false
	}

	repeat_count := 0
	for i := 0; i < len(hd.Man)-1; i++ {
		if repeat_count > 1 {
			return Thirteen_Orphans{}, false
		}
		if hd.Man[i] == hd.Man[i+1] {
			result.Repeat_Suit = 'm'
			result.Repeat_Rank = hd.Man[i]
			repeat_count++
		}
	}
	for i := 0; i < len(hd.Pin)-1; i++ {
		if repeat_count > 1 {
			return Thirteen_Orphans{}, false
		}
		if hd.Pin[i] == hd.Pin[i+1] {
			result.Repeat_Suit = 'p'
			result.Repeat_Rank = hd.Pin[i]
			repeat_count++
		}
	}

	for i := 0; i < len(hd.Sou)-1; i++ {
		if repeat_count > 1 {
			return Thirteen_Orphans{}, false
		}
		if hd.Sou[i] == hd.Sou[i+1] {
			result.Repeat_Suit = 's'
			result.Repeat_Rank = hd.Sou[i]
			repeat_count++
		}
	}

	for i := 0; i < len(hd.Zi)-1; i++ {
		if repeat_count > 1 {
			return Thirteen_Orphans{}, false
		}
		if hd.Zi[i] == hd.Zi[i+1] {
			result.Repeat_Suit = 'z'
			result.Repeat_Rank = hd.Zi[i]
			repeat_count++
		}
	}
	return result, true
}

/*
	func RemoveMenzi(hdmz NewHandAndRemovedMenzis) ([]NewHandAndRemovedMenzis, bool) {
		hd := hdmz.NewHand
		mzlist := hdmz.RemovedMenzis
		var possibleReturns [](NewHandAndRemovedMenzis)
		if len(hd.Man) != 0 {

			if hd.Man[0] == hd.Man[1] && hd.Man[1] == hd.Man[2] { // line 184
				trp, err := combination.NewTriplet('m', hd.Man[0])
				if err == nil {
					newhd := hand.Copy(hd)
					newhd.Man = newhd.Man[3:]
					newMenziList := []combination.Menzi{}
					copy(newMenziList, mzlist)
					newMenziList = append(newMenziList, trp)
					possibleReturns = append(possibleReturns, NewHandAndRemovedMenzis{NewHand: newhd, RemovedMenzis: newMenziList})
				}
			} // line 194
			if hd.Man[1]-hd.Man[0] == 1 && hd.Man[2]-hd.Man[1] == 1 {
				stra, err := combination.NewStraight('m', hd.Man[0])
				if err == nil {
					newhd := hand.Copy(hd)
					newhd.Man = newhd.Man[3:]
					newMenziList := []combination.Menzi{}
					copy(newMenziList, mzlist)
					newMenziList = append(newMenziList, stra)
					possibleReturns = append(possibleReturns, NewHandAndRemovedMenzis{NewHand: newhd, RemovedMenzis: newMenziList})
				}
			}
			// TODO:modify the code below for Pin, Sou, Zi accoring to the if statement from Man
		} else if len(hd.Pin) != 0 {
			if hd.Pin[0] == hd.Pin[1] && hd.Pin[1] == hd.Pin[2] {
				trp, err := combination.NewTriplet('p', hd.Pin[0])
				if err == nil {
					newhd := hand.Copy(hd)
					newhd.Pin = newhd.Pin[3:] // Remove elements of index = 0, 1, 2
					possibleReturns = append(possibleReturns, NewHandAndRemovedMenzi{NewHand: newhd, RemovedMenzi: trp})
				}
			}

			if hd.Pin[1]-hd.Pin[0] == 1 && hd.Pin[2]-hd.Pin[1] == 1 {
				stra, err := combination.NewStraight('p', hd.Pin[0])
				if err == nil {
					newhd := hand.Copy(hd)
					newhd.Pin = newhd.Pin[3:] // Remove elements of index = 0, 1, 2
					possibleReturns = append(possibleReturns, NewHandAndRemovedMenzi{NewHand: newhd, RemovedMenzi: stra})
				}
			}
		} else if len(hd.Sou) != 0 {
			if hd.Sou[0] == hd.Sou[1] && hd.Sou[1] == hd.Sou[2] {
				trp, err := combination.NewTriplet('s', hd.Sou[0])
				if err == nil {
					newhd := hand.Copy(hd)
					newhd.Sou = newhd.Sou[3:] // Remove elements of index = 0, 1, 2
					possibleReturns = append(possibleReturns, NewHandAndRemovedMenzi{NewHand: newhd, RemovedMenzi: trp})
				}
			}

			if hd.Sou[1]-hd.Sou[0] == 1 && hd.Sou[2]-hd.Sou[1] == 1 {
				stra, err := combination.NewStraight('s', hd.Sou[0])
				if err == nil {
					newhd := hand.Copy(hd)
					newhd.Sou = newhd.Sou[3:] // Remove elements of index = 0, 1, 2
					possibleReturns = append(possibleReturns, NewHandAndRemovedMenzi{NewHand: newhd, RemovedMenzi: stra})
				}
			}
		} else if len(hd.Zi) != 0 {
			if hd.Zi[0] == hd.Zi[1] && hd.Zi[1] == hd.Zi[2] {
				trp, err := combination.NewTriplet('s', hd.Zi[0])
				if err == nil {
					newhd := hand.Copy(hd)
					newhd.Zi = newhd.Zi[3:] // Remove elements of index = 0, 1, 2
					possibleReturns = append(possibleReturns, NewHandAndRemovedMenzi{NewHand: newhd, RemovedMenzi: trp})
				}
			}
		}
		//log.Println("Invalid input for RemovePair func")

		canRemove := false
		if len(possibleReturns) != 0 {
			canRemove = true
		}
		return possibleReturns, canRemove

}
*/
func CommonString(c Common_Win) string {
	var output string
	output += "面子:"
	for i := range c.MenziList {
		output += combination.MenziString(c.MenziList[i])
		output += " "
	}
	output += "\n雀頭:"
	output += combination.PairString(c.Eye)
	return output

}

func SevenPairString(s Seven_Pairs_Win) string {
	var output string
	output += "對子:"
	for i := range s.PairList {
		output += combination.PairString(s.PairList[i])
		output += " "
	}
	return output
}

func GetWinningTile(cw Common_Win) (byte, uint8) {
	var suit byte
	var rank uint8
	if cw.Win_Com_IDX == 4 {
		suit = cw.Eye.Suit
		rank = cw.Eye.Rank
	} else {
		suit = cw.MenziList[cw.Win_Com_IDX].Suit
		var addIDX uint8 = 0
		if cw.MenziList[cw.Win_Com_IDX].Type == 'S' {
			addIDX = uint8(cw.Win_Tile_IDX)
		}
		rank = cw.MenziList[cw.Win_Com_IDX].Rank + addIDX
	}
	return suit, rank
}
