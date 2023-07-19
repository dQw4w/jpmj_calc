package score

import (
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
