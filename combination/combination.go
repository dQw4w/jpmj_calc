package combination

import (
	"errors"
	"fmt"
)

type Pair struct {
	Suit byte
	Rank uint8
}

type Menzi struct {
	Type byte // straight:S,triplet:T,kanzi_closed:C,kanzi_open:O
	Suit byte
	Rank uint8
}
type Straight struct {
	Menzi
	Suit byte
	Rank uint8
}
type Triplet struct {
	Menzi
	Suit byte
	Rank uint8
}
type Kanzi_closed struct {
	Menzi
	Suit byte
	Rank uint8
}
type Kanzi_open struct {
	Menzi
	Suit byte
	Rank uint8
}

func isValid(suit byte, rank uint8, whattype byte) (bool, error) {
	switch suit {
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

func NewPair(suit byte, rank uint8) (Pair, error) {

	if valid, err := isValid(suit, rank, 'P'); !valid {
		return Pair{}, err
	}

	return Pair{
		Suit: suit,
		Rank: rank,
	}, nil
}

func NewStraight(suit byte, rank uint8) (Menzi, error) {

	if valid, err := isValid(suit, rank, 'S'); !valid {
		return Menzi{}, err
	}

	return Menzi{
		Type: 'S',
		Suit: suit,
		Rank: rank,
	}, nil
}

func NewTriplet(suit byte, rank uint8) (Menzi, error) {

	if valid, err := isValid(suit, rank, 'T'); !valid {
		return Menzi{}, err
	}

	return Menzi{
		Type: 'T',
		Suit: suit,
		Rank: rank,
	}, nil
}

func NewKanzi_closed(suit byte, rank uint8) (Menzi, error) {

	if valid, err := isValid(suit, rank, 'C'); !valid {
		return Menzi{}, err
	}

	return Menzi{
		Type: 'C',
		Suit: suit,
		Rank: rank,
	}, nil
}
func NewKanzi_open(suit byte, rank uint8) (Menzi, error) {

	if valid, err := isValid(suit, rank, 'O'); !valid {
		return Menzi{}, err
	}

	return Menzi{
		Type: 'O',
		Suit: suit,
		Rank: rank,
	}, nil
}
func New_Menzi(tp byte, suit byte, rank uint8) (Menzi, error) {
	//log.Print("tp:")
	//log.Println(tp)
	switch tp {
	case 'S':
		return NewStraight(suit, rank)
	case 'T':
		return NewTriplet(suit, rank)
	case 'C':
		return NewKanzi_closed(suit, rank)
	case 'O':
		return NewKanzi_open(suit, rank)
	default:
		return Menzi{}, errors.New("Invalid input for New_Menzi func")
	}
}
func PairString(p Pair) string {
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
func MenziString(m Menzi) string {
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
