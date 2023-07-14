package combination

import (
	"errors"
)

type Pair struct {
	Suit byte
	Rank uint8
}

type Menzi struct {
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

func NewPair(suit byte, rank uint8) (Pair, error) {

	if valid, err := isValid(suit, rank, 'P'); !valid {
		return Pair{}, err
	}

	return Pair{
		Suit: suit,
		Rank: rank,
	}, nil
}

func NewStraight(suit byte, rank uint8) (Straight, error) {

	if valid, err := isValid(suit, rank, 'S'); !valid {
		return Straight{}, err
	}

	return Straight{
		Suit: suit,
		Rank: rank,
	}, nil
}

func NewTriplet(suit byte, rank uint8) (Triplet, error) {

	if valid, err := isValid(suit, rank, 'T'); !valid {
		return Triplet{}, err
	}

	return Triplet{
		Suit: suit,
		Rank: rank,
	}, nil
}
func NewKanzi_closed(suit byte, rank uint8) (Kanzi_closed, error) {

	if valid, err := isValid(suit, rank, 'C'); !valid {
		return Kanzi_closed{}, err
	}

	return Kanzi_closed{
		Suit: suit,
		Rank: rank,
	}, nil
}
func NewKanzi_open(suit byte, rank uint8) (Kanzi_open, error) {

	if valid, err := isValid(suit, rank, 'O'); !valid {
		return Kanzi_open{}, err
	}

	return Kanzi_open{
		Suit: suit,
		Rank: rank,
	}, nil
}
