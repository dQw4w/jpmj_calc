package score

import (
	"fmt"
	"math"
)

func calcBscPt(han, fu int) int {
	basicPt := 0
	if han <= 4 {
		basicPt = fu * int(math.Pow(2, float64(han+2)))
		if basicPt > 2000 {
			basicPt = 2000
		}
	} else {
		switch han {
		case 5:
			basicPt = 2000
		case 6, 7:
			basicPt = 3000
		case 8, 9, 10:
			basicPt = 4000
		case 11, 12:
			basicPt = 6000
		default:
			basicPt = 8000
		}
	}
	return basicPt
}
func YakumanRonPoints(yakuman int, oyaka bool) int {
	pt := yakuman * 32000
	if oyaka {
		pt *= 3
		pt /= 2
	}
	return pt
}
func YakumanTsumoPoints(yakuman int, oyaka bool) (int, int) {
	bscpt := yakuman * 8000
	if oyaka {
		return 0, 2 * bscpt
	}
	return 2 * bscpt, bscpt
}
func RonPoints(han, fu int, oyaka bool) int {
	basicPt := calcBscPt(han, fu)
	pt := 0
	if oyaka {
		pt = basicPt * 6
	} else {
		pt = basicPt * 4
	}

	for pt%100 != 0 {
		pt++
	}
	return pt
}

func TsumoPoints(han, fu int, oyaka bool) (int, int) {
	basicPt := calcBscPt(han, fu)
	koPt := 0
	oyaPt := 0

	if oyaka {
		koPt = basicPt * 2
		for koPt%100 != 0 {
			koPt++
		}
	} else {
		koPt = basicPt
		for koPt%100 != 0 {
			koPt++
		}

		oyaPt = basicPt * 2
		for oyaPt%100 != 0 {
			oyaPt++
		}
	}

	return oyaPt, koPt
}
func CalcYakumanPointsString(yakuman int, oyaka, tsumo bool, msg string) string {

	if yakuman > 1 {
		msg += fmt.Sprintf("%v倍", yakuman)
	}
	msg += "役満\n"

	if tsumo {
		var total_pt int

		oya, ko := YakumanTsumoPoints(yakuman, oyaka)
		if oyaka {
			total_pt = 3 * ko
			msg += fmt.Sprintf("%v点(%v オール)\n", total_pt, ko)
		} else {
			total_pt = 2*ko + oya
			msg += fmt.Sprintf("%v点(%v,%v)\n", total_pt, ko, oya)
		}

	} else {
		pt := YakumanRonPoints(yakuman, oyaka)
		msg += fmt.Sprintf("%v点\n", pt)
	}
	return msg
}
func CalcPointsString(han, hu int, oyaka, tsumo bool, msg string) string {
	var hanmsg string
	if han >= 13 {
		hanmsg = "数え役満"
	} else if han >= 11 {
		hanmsg = "三倍満"
	} else if han >= 8 {
		hanmsg = "倍満"
	} else if han >= 6 {
		hanmsg = "跳満"
	} else if han >= 5 {
		hanmsg = "満貫"
	} else if han >= 4 && hu >= 40 {
		hanmsg = "満貫"
	} else if han >= 3 && hu >= 70 {
		hanmsg = "満貫"
	}
	msg += fmt.Sprintf("%v飜%v符 %s\n", han, hu, hanmsg)

	if tsumo {
		var total_pt int

		oya, ko := TsumoPoints(han, hu, oyaka)
		if oyaka {
			total_pt = 3 * ko
			msg += fmt.Sprintf("%v点(%v オール)\n", total_pt, ko)
		} else {
			total_pt = 2*ko + oya
			msg += fmt.Sprintf("%v点(%v,%v)\n", total_pt, ko, oya)
		}

	} else {
		pt := RonPoints(han, hu, oyaka)
		msg += fmt.Sprintf("%v点\n", pt)
	}

	return msg
}
