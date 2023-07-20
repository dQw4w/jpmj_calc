package hu

import (
	"jpmj_calc/combination"
	"jpmj_calc/win"
)

func CalcHu(cw win.Common_Win) int {
	hu := 20
	for i := range cw.MenziList { // dealing with trp,kanzi
		menzi := cw.MenziList[i]
		if menzi.Type != 'S' {
			trphu := 2
			if combination.IsYaoMenzi(menzi) {
				trphu *= 2
			}
			if !menzi.Furo {
				trphu *= 2
			}
			if menzi.Type == 'C' || menzi.Type == 'O' {
				trphu *= 4
			}
			hu += trphu
		}
	}
	if cw.Eye.Suit == 'z' { // eye yakuhai
		if cw.Eye.Rank > 4 {
			hu += 2
		} else {
			if cw.Eye.Rank == cw.SelfWind {
				hu += 2
			}
			if cw.Eye.Rank == cw.FieldWind {
				hu += 2
			}
		}
	}
	if cw.Win_Tile_IDX == 1 { //middle win
		hu += 2
	}
	if cw.Win_Tile_IDX == 0 { //"7"89
		if cw.Win_Com_IDX != 4 {
			if cw.MenziList[cw.Win_Com_IDX].Type == 'S' && cw.MenziList[cw.Win_Com_IDX].Rank == 7 {
				hu += 2
			}
		}
	}
	if cw.Win_Tile_IDX == 2 { //12"3"
		if cw.MenziList[cw.Win_Com_IDX].Type == 'S' && cw.MenziList[cw.Win_Com_IDX].Rank == 1 {
			hu += 2
		}
	}
	if cw.Win_Tile_IDX == 4 { //tanki
		hu += 2
	}
	if cw.Tsumo && !cw.Menchin { //tsumo
		hu += 2
	}
	if cw.Menchin && !cw.Tsumo { //menchin ron
		hu += 10
	}
	if !cw.Menchin && hu < 30 { //at least 30 if not menchin pinfu tsumo
		hu = 30
	}
	//log.Println("originalhu:", hu)
	for hu%10 != 0 {
		hu++
	}
	return hu
}
