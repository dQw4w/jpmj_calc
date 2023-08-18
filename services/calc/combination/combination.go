package combination

import (
	"errors"
	"fmt"
)

type Pair struct {
	Suit byte  //花色 m = 萬 p = 筒 s = 索 z = 字
	Rank uint8 //牌的大小 字牌 : 1234=東南西北 567=白發中
	Furo bool  //明true暗false
}

type Menzi struct {
	Type byte // straight:S,triplet:T,kanzi_closed:C,kanzi_open:O 順子 刻子 暗槓 明槓
	Suit byte
	Rank uint8 //the rank of the smallest tile of a Menzi ex. 123m => rank = 1
	Furo bool
}

func isValid(suit byte, rank uint8, whattype byte) (bool, error) {
	switch suit { //check
	case 'm', 'p', 's':
		if rank > 9 {
			return false, errors.New("Invaid rank")
		}
	case 'z':
		if rank > 7 {
			return false, errors.New("Invalid rank")
		}
	default:
		return false, errors.New("Invalid suit")
	}

	switch whattype {
	case 'S': // straight
		if suit == 'z' {
			return false, errors.New("Invalid suit")
		}
		if rank > 7 {
			return false, errors.New("Invalid rank")
		}
		return true, nil
	case 'P', 'T', 'O', 'C': //pair, triplet, kanzi(open and closed)
		return true, nil
	default:
		return false, errors.New("Invalid type")
	}
	//return false, errors.New("Something went wrong with isValid func")
}
func IsEmptyMenzi(menzi Menzi) bool {
	if menzi.Type == 0 {
		return true
	}
	return false
}

func IsEmptyPair(pair Pair) bool {
	if pair.Suit == 0 {
		return true
	}
	return false
}

func NewPair(suit byte, rank uint8, furo bool) (Pair, error) {

	if valid, err := isValid(suit, rank, 'P'); !valid {
		return Pair{}, err
	}

	return Pair{
		Suit: suit,
		Rank: rank,
		Furo: furo,
	}, nil
}

func NewStraight(suit byte, rank uint8, furo bool) (Menzi, error) {

	if valid, err := isValid(suit, rank, 'S'); !valid {
		return Menzi{}, err
	}

	return Menzi{
		Type: 'S',
		Suit: suit,
		Rank: rank,
		Furo: furo,
	}, nil
}

func NewTriplet(suit byte, rank uint8, furo bool) (Menzi, error) {

	if valid, err := isValid(suit, rank, 'T'); !valid {
		return Menzi{}, err
	}

	return Menzi{
		Type: 'T',
		Suit: suit,
		Rank: rank,
		Furo: furo,
	}, nil
}

func NewKanzi_closed(suit byte, rank uint8, furo bool) (Menzi, error) {

	if valid, err := isValid(suit, rank, 'C'); !valid {
		return Menzi{}, err
	}

	return Menzi{
		Type: 'C',
		Suit: suit,
		Rank: rank,
		Furo: furo,
	}, nil
}
func NewKanzi_open(suit byte, rank uint8, furo bool) (Menzi, error) {

	if valid, err := isValid(suit, rank, 'O'); !valid {
		return Menzi{}, err
	}

	return Menzi{
		Type: 'O',
		Suit: suit,
		Rank: rank,
		Furo: furo,
	}, nil
}
func New_Menzi(tp byte, suit byte, rank uint8, furo bool) (Menzi, error) {
	//log.Print("tp:")
	//log.Println(tp)
	switch tp {
	case 'S':
		return NewStraight(suit, rank, furo)
	case 'T':
		return NewTriplet(suit, rank, furo)
	case 'C':
		if furo {
			return Menzi{}, errors.New("Invalid input for New_Menzi func")
		}
		return NewKanzi_closed(suit, rank, furo)
	case 'O':
		if !furo {
			return Menzi{}, errors.New("Invalid input for New_Menzi func")
		}
		return NewKanzi_open(suit, rank, furo)
	default:
		return Menzi{}, errors.New("Invalid input for New_Menzi func")
	}
}
func PairString(p Pair) string { // for printing and debugging
	var output, suitstr string
	switch p.Suit {
	case 'm':
		suitstr = "萬"
	case 'p':
		suitstr = "筒"
	case 's':
		suitstr = "索"
	case 'z':
		switch p.Rank {
		case 1:
			suitstr = "東"
		case 2:
			suitstr = "南"
		case 3:
			suitstr = "西"
		case 4:
			suitstr = "北"
		case 5:
			suitstr = "白"
		case 6:
			suitstr = "發"
		case 7:
			suitstr = "中"
		default:
			suitstr = ""
		}
	default:
		suitstr = ""
	}

	if p.Suit == 'z' {
		output = fmt.Sprintf("%s%s", suitstr, suitstr)
	} else {
		output = fmt.Sprintf("%v%v%s", p.Rank, p.Rank, suitstr)
	}

	return output
}
func MenziString(m Menzi) string { // for printing and debugging
	var output, suitstr string
	switch m.Suit {
	case 'm':
		suitstr = "萬"
	case 'p':
		suitstr = "筒"
	case 's':
		suitstr = "索"
	case 'z':
		switch m.Rank {
		case 1:
			suitstr = "東"
		case 2:
			suitstr = "南"
		case 3:
			suitstr = "西"
		case 4:
			suitstr = "北"
		case 5:
			suitstr = "白"
		case 6:
			suitstr = "發"
		case 7:
			suitstr = "中"
		default:
			suitstr = ""
		}
	default:
		suitstr = ""
	}
	switch m.Type {
	case 'S':
		output = fmt.Sprintf("%v%v%v%s", m.Rank, m.Rank+1, m.Rank+2, suitstr)
	case 'T':
		if m.Suit == 'z' {
			output = fmt.Sprintf("%s%s%s", suitstr, suitstr, suitstr)
		} else {
			output = fmt.Sprintf("%v%v%v%s", m.Rank, m.Rank, m.Rank, suitstr)
		}
	case 'C':
		if m.Suit == 'z' {
			output = fmt.Sprintf("X%s%sX", suitstr, suitstr)
		} else {
			output = fmt.Sprintf("X%v%vX%s", m.Rank, m.Rank, suitstr)
		}
	case 'O':
		if m.Suit == 'z' {
			output = fmt.Sprintf("%s%s%s%s", suitstr, suitstr, suitstr, suitstr)
		} else {
			output = fmt.Sprintf("%v%v%v%v%s", m.Rank, m.Rank, m.Rank, m.Rank, suitstr)
		}
	}
	return output
}
func ConvertStrToMenzi(str string) (Menzi, int /*akadora*/, error) {
	var akadora int
	length := len(str)
	ranks := []uint8{}
	for i := 0; i < length-1; i++ {
		if str[i] == '0' { // converts 0 to 5, and add to akadora count
			akadora++
			ranks = append(ranks, 5)
		} else if str[i]-'0' <= 9 {
			ranks = append(ranks, uint8(str[i]-'0'))
		} else if str[i] == 'X' {
			ranks = append(ranks, 69) //69 represents closed kanzi
		} else {
			return Menzi{}, 0, errors.New("Invalid input for ConvertStrToMenzi func")
		}
	}
	if length == 4 {
		switch str[3] {
		case 'm', 'p', 's', 'z':
			if ranks[0] == ranks[1] && ranks[1] == ranks[2] {
				trp, err := NewTriplet(str[3], ranks[0], true)
				if err == nil {
					return trp, 0, nil
				}
			} else if ranks[2]-ranks[1] == 1 && ranks[1]-ranks[0] == 1 {
				stra, err := NewStraight(str[3], ranks[0], true)
				if err == nil {
					return stra, 0, nil
				}
			}

		}
	} else if length == 5 {
		switch str[4] {
		case 'm', 'p', 's', 'z':
			if ranks[0] == 69 && ranks[3] == 69 && ranks[1] == ranks[2] {
				closekan, err := NewKanzi_closed(str[4], ranks[1], false)
				if err == nil {
					return closekan, 0, nil
				}
			} else if ranks[0] == ranks[1] && ranks[2] == ranks[3] && ranks[1] == ranks[2] {
				openkan, err := NewKanzi_open(str[4], ranks[1], true)
				if err == nil {
					return openkan, 0, nil
				}
			}
		}
	}
	return Menzi{}, 0, errors.New("Invalid input for ConvertStrToMenzi func")

}
func SameMenzi(a Menzi, b Menzi) bool {
	if a.Type != b.Type {
		return false
	}
	if a.Suit != b.Suit {
		return false
	}
	if a.Rank != b.Rank {
		return false
	}
	return true
}
func SamePair(a Pair, b Pair) bool {
	return (a.Rank == b.Rank && a.Suit == b.Suit)
}
func InMenziandIndex(suit byte, rank uint8, m Menzi) (bool, int) {
	if m.Suit != suit {
		return false, -1
	}
	if m.Type == 'S' {
		idx := int(rank - m.Rank)
		if idx < 3 && idx >= 0 {
			return true, idx
		}
		return false, -1
	}
	if rank == m.Rank {
		return true, 0
	}
	return false, -1
}
func InPairandIndex(suit byte, rank uint8, p Pair) (bool, int) {
	if p.Suit != suit {
		return false, -1
	}
	if rank == p.Rank {
		return true, 0
	}
	return false, -1
}
func IsYaoMenzi(m Menzi) bool { // YaoTile : tile rank == 1,9 or suit = 'z', YaoMenzi: a menzi includes at least one YaoTile
	if m.Suit == 'z' {
		return true
	} else if m.Type == 'S' {
		return m.Rank == 1 || m.Rank == 7
	}
	return m.Rank == 1 || m.Rank == 9
}
func IsYaoPair(p Pair) bool {
	return (p.Suit == 'z' || p.Rank == 1 || p.Rank == 9)
}
