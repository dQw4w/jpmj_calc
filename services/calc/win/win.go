package win

import (
	"errors"

	"github.com/dQw4w/jpmj_calc/services/calc/hand"

	"github.com/dQw4w/jpmj_calc/services/calc/combination"
)

type Wind uint8

const (
	EAST Wind = iota + 1
	SOUTH
	WEST
	NORTH
)

type Common_Win struct { //一般型
	MenziList  [4]combination.Menzi
	Eye        combination.Pair
	WinComIDX  int // 4 represents eye
	WinTileIDX int
	Tsumo      bool
	Menchin    bool
	SelfWind   uint8 //1234 ESWN
	FieldWind  uint8 //1234

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

	Akadora      int
	MotedoraSuit []byte
	MotedoraRank []uint8

	UradoraSuit []byte
	UradoraRank []uint8
}

type Seven_Pairs_Win struct { //七對子型
	PairList  [7]combination.Pair
	WinComIDX int
	Tsumo     bool
	Menchin   bool
	SelfWind  uint8 //1234 ESWN
	FieldWind uint8 //1234

	//specials
	Reach       bool
	DoubleReach bool // Reach and Double Reach only one is true
	RinShan     bool
	HaiTei      bool
	HoTei       bool
	Ippatsu     bool

	//yakuman special
	TenHo bool
	JiHo  bool

	Akadora      int
	MotedoraSuit []byte
	MotedoraRank []uint8

	UradoraSuit []byte
	UradoraRank []uint8
}
type Thirteen_Orphans struct { //國士無雙型
	Repeat_Suit   byte
	Repeat_Rank   uint8
	Thirteen_Wait bool

	Tsumo bool
	Oyaka bool
	TenHo bool
	JiHo  bool
}

func CreateEmptyCommon() Common_Win {
	New := Common_Win{}
	//New.MenziList = [4]combination.Menzi{}
	return New
}
func CreateEmptySevenPair() Seven_Pairs_Win {
	New := Seven_Pairs_Win{}
	return New
}
func AddMenzi(menzi combination.Menzi, common Common_Win) (Common_Win, error) {
	for i := range common.MenziList {
		if combination.IsEmptyMenzi(common.MenziList[i]) {
			common.MenziList[i] = menzi
			return common, nil
		}
	}
	return common, errors.New("error: Menzilist is already full")
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
			if i != 0 && combination.SamePair(sevenpairwin.PairList[i-1], pair) {
				return sevenpairwin, errors.New("error: repeated pair")
			}
			sevenpairwin.PairList[i] = pair
			return sevenpairwin, nil
		}
	}
	return sevenpairwin, errors.New("error: pairlist is already full")
}

type NewHandAndRemovedMenzis struct {
	NewHand       hand.Hand
	RemovedMenzis []combination.Menzi
}

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

			//NewWin := Common_Win{}
			NewWin, _ := CopyCommon(cw)
			//log.Println("Newwin:", NewWin)
			NewWin, _ = SetPair(pair, NewWin)
			possibles := []NewHandAndRemovedMenzis{ /*NewHandAndRemovedMenzis*/ {NewHand: newhd}}

			for i := 0; i < 4-Menzi_Count(cw); i++ {
				cases := len(possibles)
				for j := 0; j < cases; j++ {
					newPossibilities, _ := RemoveMenzi(possibles[0])
					possibles = append(possibles[1:], newPossibilities...)
				}
			}

			for i := range possibles {
				NewNewWin, _ := CopyCommon(NewWin)
				for j := range possibles[i].RemovedMenzis {
					NewNewWin, _ = AddMenzi(possibles[i].RemovedMenzis[j], NewNewWin)
				}
				result = append(result, NewNewWin)
			}

		}
	}
	canCreate := false
	len_result := len(result)
	if len_result != 0 {
		canCreate = true
	}

	for i := 0; i < len_result; i++ { // set the menzi including the win tile to furo
		count := 0
		for j := range result[i].MenziList {
			if inmenzi, idx := combination.InMenziandIndex(hd.Win_Suit, hd.Win_Rank, result[i].MenziList[j]); inmenzi {
				count++
				if count > 1 {
					win_copy, _ := CopyCommon(result[i])
					win_copy.WinComIDX = j
					win_copy.WinTileIDX = idx
					if win_copy.Tsumo {
						win_copy.MenziList[j].Furo = false
					} else {
						win_copy.MenziList[j].Furo = true
					}
					result = append(result, win_copy)
				} else {
					result[i].WinComIDX = j
					result[i].WinTileIDX = idx
					if result[i].Tsumo {
						result[i].MenziList[j].Furo = false
					} else {
						result[i].MenziList[j].Furo = true
					}
				}
			}
		}
		if inpair, idx := combination.InPairandIndex(hd.Win_Suit, hd.Win_Rank, result[i].Eye); inpair {
			count++
			if count > 1 {
				win_copy, _ := CopyCommon(result[i])
				win_copy.WinComIDX = 4
				win_copy.WinTileIDX = idx
				if win_copy.Tsumo {
					win_copy.Eye.Furo = false
				} else {
					win_copy.Eye.Furo = true
				}
				result = append(result, win_copy)
			} else {
				result[i].WinComIDX = 4
				result[i].WinTileIDX = idx
				if result[i].Tsumo {
					result[i].Eye.Furo = false
				} else {
					result[i].Eye.Furo = true
				}
			}
		}
	}

	return result, canCreate
}

func CopyCommon(src Common_Win) (Common_Win, error) {
	dst := Common_Win{}
	dst.MenziList = [4]combination.Menzi{}
	var err2 error
	dst.Eye, _ = combination.NewPair(src.Eye.Suit, src.Eye.Rank, src.Eye.Furo)

	dst.WinComIDX = src.WinComIDX
	dst.WinTileIDX = src.WinTileIDX
	dst.Tsumo = src.Tsumo
	dst.Menchin = src.Menchin
	dst.SelfWind = src.SelfWind
	dst.FieldWind = src.FieldWind
	dst.Reach = src.Reach
	dst.DoubleReach = src.DoubleReach
	dst.ChanKan = src.ChanKan
	dst.RinShan = src.RinShan
	dst.HaiTei = src.HaiTei
	dst.HoTei = src.HoTei
	dst.Ippatsu = src.Ippatsu
	dst.TenHo = src.TenHo
	dst.JiHo = src.JiHo
	dst.Akadora = src.Akadora
	dst.MotedoraSuit = src.MotedoraSuit
	dst.MotedoraRank = src.MotedoraRank
	dst.UradoraSuit = src.UradoraSuit
	dst.UradoraRank = src.UradoraRank

	for i := range dst.MenziList {
		dst.MenziList[i], err2 = combination.New_Menzi(src.MenziList[i].Type, src.MenziList[i].Suit, src.MenziList[i].Rank, src.MenziList[i].Furo)
		if err2 != nil {
			//panic(err2)
			return dst, err2
		}
	}

	return dst, nil
}
func RemovePair(hd hand.Hand, idx int) (hand.Hand, combination.Pair, bool) { // check if the tile of index = idx in hd can be removed as a pair with another tile
	var mLen, pLen, sLen, zLen = hand.AllLen(hd)

	if idx < mLen {
		if len(hd.Man)%3 != 2 {
			return hand.Hand{}, combination.Pair{}, false
		}
		if mLen > idx+1 && hd.Man[idx] == hd.Man[idx+1] {
			pair, err := combination.NewPair('m', hd.Man[idx], false)

			if err == nil {
				newhd := hand.Copy(hd)

				newhd.Man = append(newhd.Man[:idx], newhd.Man[idx+2:]...) //remove elements of index = idx, idx+1

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
func RemoveMenzi(hdmz NewHandAndRemovedMenzis) ([]NewHandAndRemovedMenzis, bool) { // check if the tile of index = idx in hd can be removed as a menzi with other tiles
	//returns all posibilities
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
func CreateSevenPair(hd hand.Hand, sp Seven_Pairs_Win) (Seven_Pairs_Win, bool) {

	result := CopySevenPair(sp)
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

	for i := range result.PairList {
		if hd.Win_Rank == result.PairList[i].Rank && hd.Win_Suit == result.PairList[i].Suit {
			result.WinComIDX = i
			break
		}
	}
	return result, true

}
func CopySevenPair(sp Seven_Pairs_Win) Seven_Pairs_Win {
	new := Seven_Pairs_Win{
		WinComIDX:    sp.WinComIDX,
		Tsumo:        sp.Tsumo,
		Menchin:      sp.Menchin,
		SelfWind:     sp.SelfWind,
		FieldWind:    sp.FieldWind,
		Reach:        sp.Reach,
		DoubleReach:  sp.DoubleReach,
		RinShan:      sp.RinShan,
		HaiTei:       sp.HaiTei,
		HoTei:        sp.HoTei,
		Ippatsu:      sp.Ippatsu,
		TenHo:        sp.TenHo,
		JiHo:         sp.JiHo,
		Akadora:      sp.Akadora,
		MotedoraSuit: make([]byte, len(sp.MotedoraSuit)),
		MotedoraRank: make([]uint8, len(sp.MotedoraRank)),
		UradoraSuit:  make([]byte, len(sp.UradoraSuit)),
		UradoraRank:  make([]uint8, len(sp.UradoraRank)),
	}

	copy(new.MotedoraSuit, sp.MotedoraSuit)
	copy(new.MotedoraRank, sp.MotedoraRank)
	copy(new.UradoraSuit, sp.UradoraSuit)
	copy(new.UradoraRank, sp.UradoraRank)

	new.PairList = [7]combination.Pair{}
	copy(new.PairList[:], sp.PairList[:])

	return new
}

func OnlyYao_Slice(tiles []uint8) bool {
	for i := range tiles {
		if tiles[i] != 1 && tiles[i] != 9 {
			return false
		}
	}
	return true
}
func Create13Orphans(hd hand.Hand, tsumo, oyaka, tenho, jiho bool) (Thirteen_Orphans, bool) {
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
	if result.Repeat_Rank == hd.Win_Rank && result.Repeat_Suit == hd.Win_Suit {
		result.Thirteen_Wait = true
	}
	result.Tsumo = tsumo
	result.Oyaka = oyaka
	result.TenHo = tenho
	result.JiHo = jiho
	return result, true
}

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
	if cw.WinComIDX == 4 {
		suit = cw.Eye.Suit
		rank = cw.Eye.Rank
	} else {
		suit = cw.MenziList[cw.WinComIDX].Suit
		var addIDX uint8 = 0
		if cw.MenziList[cw.WinComIDX].Type == 'S' {
			addIDX = uint8(cw.WinTileIDX)
		}
		rank = cw.MenziList[cw.WinComIDX].Rank + addIDX
	}
	return suit, rank
}

func ConvertWinToMap(cw Common_Win) map[byte](map[uint8]int) {
	//log.Println("wwwwwwwwwwww")
	result := make(map[byte](map[uint8]int))
	//log.Println("qqqqqqqqqqq")
	//log.Println("init:", result)
	result['m'] = make(map[uint8]int)
	result['p'] = make(map[uint8]int)

	result['s'] = make(map[uint8]int)

	result['z'] = make(map[uint8]int)

	for i := range cw.MenziList {
		//log.Println("iiiiiiiiiii")
		menzi := cw.MenziList[i]
		//log.Println("popopopopo")
		switch menzi.Type {
		case 'S':
			for i := 0; i < 3; i++ {
				result[menzi.Suit][menzi.Rank+uint8(i)]++
			}
		case 'T':
			result[menzi.Suit][menzi.Rank] += 3
		case 'C', 'O':
			result[menzi.Suit][menzi.Rank] += 4
		}
		//log.Println("vvvvvvvvvvvvvvvv")
	}
	result[cw.Eye.Suit][cw.Eye.Rank] += 2
	return result
}
func ConvertSevenPairsToMap(sp Seven_Pairs_Win) map[byte](map[uint8]int) {
	result := make(map[byte](map[uint8]int))
	result['m'] = make(map[uint8]int)
	result['p'] = make(map[uint8]int)

	result['s'] = make(map[uint8]int)

	result['z'] = make(map[uint8]int)
	for i := range sp.PairList {
		pair := sp.PairList[i]

		result[pair.Suit][pair.Rank] += 2
	}

	return result
}
func ConvertCommonToSeven(cw Common_Win) Seven_Pairs_Win {
	//TODO: copy the elements that Common_Win and Seven_Pairs_Win have in common, from cw into a new Seven_Pairs_Win, and return the Seven_Pairs_Win

	sp := Seven_Pairs_Win{
		WinComIDX:    cw.WinComIDX,
		Tsumo:        cw.Tsumo,
		Menchin:      cw.Menchin,
		SelfWind:     cw.SelfWind,
		FieldWind:    cw.FieldWind,
		Reach:        cw.Reach,
		DoubleReach:  cw.DoubleReach,
		RinShan:      cw.RinShan,
		HaiTei:       cw.HaiTei,
		HoTei:        cw.HoTei,
		Ippatsu:      cw.Ippatsu,
		TenHo:        cw.TenHo,
		JiHo:         cw.JiHo,
		Akadora:      cw.Akadora,
		MotedoraSuit: make([]byte, len(cw.MotedoraSuit)),
		MotedoraRank: make([]uint8, len(cw.MotedoraRank)),
		UradoraSuit:  make([]byte, len(cw.UradoraSuit)),
		UradoraRank:  make([]uint8, len(cw.UradoraRank)),
	}

	copy(sp.MotedoraSuit, cw.MotedoraSuit)
	copy(sp.MotedoraRank, cw.MotedoraRank)
	copy(sp.UradoraSuit, cw.UradoraSuit)
	copy(sp.UradoraRank, cw.UradoraRank)

	return sp

}
