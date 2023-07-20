package hand

import (
	"errors"
	"fmt"
)

type Hand struct {
	Man      []uint8
	Pin      []uint8
	Sou      []uint8
	Zi       []uint8
	Win_Suit byte
	Win_Rank uint8
}

func Len(hand Hand) int {
	return len(hand.Man) + len(hand.Pin) + len(hand.Sou) + len(hand.Zi)
}
func AllLen(hand Hand) (int, int, int, int) {
	return len(hand.Man), len(hand.Pin), len(hand.Sou), len(hand.Zi)
}
func Copy(src Hand) Hand {
	dst := Hand{}
	dst.Man = make([]uint8, len(src.Man))
	dst.Pin = make([]uint8, len(src.Pin))
	dst.Sou = make([]uint8, len(src.Sou))
	dst.Zi = make([]uint8, len(src.Zi))
	copy(dst.Man, src.Man)
	copy(dst.Pin, src.Pin)
	copy(dst.Sou, src.Sou)
	copy(dst.Zi, src.Zi)
	dst.Win_Suit = src.Win_Suit
	dst.Win_Rank = src.Win_Rank
	return dst
}
func IsEmptyHand(hand Hand) bool {
	if Len(hand) == 0 {
		return true
	}
	return false
}

func ConvertStrToHand(str_hand string) (Hand, error) {

	var states = [5]byte{'m', 'p', 's', 'z', 'f'} // f represents finish
	state_idx := 0
	out := Hand{}

	for i := range str_hand {

		curbyte := str_hand[i]
		curint := curbyte - '0'
		curstate := states[state_idx]
		var maxrank uint8
		if curstate == 'z' {
			maxrank = 7
		} else {
			maxrank = 9
		}
		if curint <= maxrank {
			switch curstate {
			case 'm':
				out.Man = append(out.Man, curint)
			case 'p':
				out.Pin = append(out.Pin, curint)
			case 's':
				out.Sou = append(out.Sou, curint)
			case 'z':
				if curint != 0 {
					out.Zi = append(out.Zi, curint)
				} else {
					msg := fmt.Sprintf("Invalid input at %s", string(curbyte))
					return Hand{}, errors.New(msg)
				}
			case 'f':
				return Hand{}, errors.New("Extra letters behind z")
			}
			continue
		}
		if curbyte == curstate {
			state_idx++
		} else {
			msg := fmt.Sprintf("Invalid input at %s", string(curbyte))
			return Hand{}, errors.New(msg)
		}

	}
	return out, nil
}
func AppendOne(hd Hand, str_mj string) (Hand, error) {
	var rk uint8 = str_mj[0] - '0'
	if len(str_mj) != 2 || rk > 9 {
		return hd, errors.New("Error: invalid input for AppendOne func")
	}
	switch str_mj[1] {
	case 'm':
		hd.Man = append(hd.Man, rk)
	case 'p':
		hd.Pin = append(hd.Pin, rk)
	case 's':
		hd.Sou = append(hd.Sou, rk)
	case 'z':
		if rk > 7 || rk == 0 {
			return hd, errors.New("Error: invalid input for AppendOne func")
		}
		hd.Zi = append(hd.Zi, rk)
		//log.Println("hd zi after:", hd.Zi)
	}
	hd.Win_Suit = str_mj[1]
	hd.Win_Rank = rk
	return hd, nil
}

// quick sort funcs cited from https://blog.boot.dev/golang/quick-sort-golang/
func partition(arr []uint8, low, high int) ([]uint8, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	//fmt.Println("111w")
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
func quickSort(arr []uint8, low, high int) []uint8 {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []uint8) []uint8 {
	return quickSort(arr, 0, (len(arr) - 1))
}
func SortColor(arr []uint8) ([]uint8, uint8) {
	var akadora uint8
	for i := range arr {
		if arr[i] == 0 {
			arr[i] = 5
			akadora++
		}
	}
	//log.Println(arr)
	return quickSortStart(arr), akadora
}
func SortAndReturnAkadora(hand *Hand) uint8 {
	var akaMan, akaPin, akaSou uint8

	hand.Man, akaMan = SortColor(hand.Man)

	hand.Pin, akaPin = SortColor(hand.Pin)

	hand.Sou, akaSou = SortColor(hand.Sou)

	hand.Zi, _ = SortColor(hand.Zi)

	return akaMan + akaPin + akaSou
}
func IsValidHandNum(hand Hand) bool {
	if Len(hand)%3 != 2 {
		return false
	}

	manLen := len(hand.Man)
	pinLen := len(hand.Pin)
	souLen := len(hand.Sou)
	ziLen := len(hand.Zi)

	if (manLen%3 == 0 && pinLen%3 == 0 && souLen%3 == 0) || (manLen%3 == 0 && pinLen%3 == 0 && ziLen%3 == 0) ||
		(manLen%3 == 0 && souLen%3 == 0 && ziLen%3 == 0) || (pinLen%3 == 0 && souLen%3 == 0 && ziLen%3 == 0) {
		return true
	}

	return false

}
