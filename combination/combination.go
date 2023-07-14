package combination

import (
	"errors"
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
