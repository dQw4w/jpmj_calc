package combination

import (
	"reflect"
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		suit     byte
		rank     uint8
		whattype byte
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		// Valid pair
		{
			name: "Valid pair",
			args: struct {
				suit     byte
				rank     uint8
				whattype byte
			}{
				suit:     's',
				rank:     3,
				whattype: 'P',
			},
			want:    true,
			wantErr: false,
		},
		// Invalid suit
		{
			name: "Invalid suit",
			args: struct {
				suit     byte
				rank     uint8
				whattype byte
			}{
				suit:     'x',
				rank:     3,
				whattype: 'P',
			},
			want:    false,
			wantErr: true,
		},
		// Invalid rank
		{
			name: "Invalid rank",
			args: struct {
				suit     byte
				rank     uint8
				whattype byte
			}{
				suit:     's',
				rank:     12,
				whattype: 'P',
			},
			want:    false,
			wantErr: true,
		},
		// Invalid type
		{
			name: "Invalid type",
			args: struct {
				suit     byte
				rank     uint8
				whattype byte
			}{
				suit:     's',
				rank:     3,
				whattype: 'X',
			},
			want:    false,
			wantErr: true,
		},
		// Valid straight
		{
			name: "Valid straight",
			args: struct {
				suit     byte
				rank     uint8
				whattype byte
			}{
				suit:     's',
				rank:     6,
				whattype: 'S',
			},
			want:    true,
			wantErr: false,
		},
		// Invalid straight with 'z' suit
		{
			name: "Invalid straight with 'z' suit",
			args: struct {
				suit     byte
				rank     uint8
				whattype byte
			}{
				suit:     'z',
				rank:     5,
				whattype: 'S',
			},
			want:    false,
			wantErr: true,
		},
		// Invalid straight with rank > 7
		{
			name: "Invalid straight with rank > 7",
			args: struct {
				suit     byte
				rank     uint8
				whattype byte
			}{
				suit:     's',
				rank:     9,
				whattype: 'S',
			},
			want:    false,
			wantErr: true,
		},
		// Invalid straight (zihai))
		{
			name: "Invalid straight (zihai)",
			args: struct {
				suit     byte
				rank     uint8
				whattype byte
			}{
				suit:     'z',
				rank:     1,
				whattype: 'S',
			},
			want:    false,
			wantErr: true,
		},
	}
	/*for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isValid(tt.args.suit, tt.args.rank, tt.args.whattype)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("isValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}*/
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isValid(tt.args.suit, tt.args.rank, tt.args.whattype)
			if (err != nil) != tt.wantErr {
				t.Errorf("isValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPair(t *testing.T) {
	type args struct {
		suit byte
		rank uint8
	}
	tests := []struct {
		name    string
		args    args
		want    Pair
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPair(tt.args.suit, tt.args.rank)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStraight(t *testing.T) {
	type args struct {
		suit byte
		rank uint8
	}
	tests := []struct {
		name    string
		args    args
		want    Straight
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStraight(tt.args.suit, tt.args.rank)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStraight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTriplet(t *testing.T) {
	type args struct {
		suit byte
		rank uint8
	}
	tests := []struct {
		name    string
		args    args
		want    Triplet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTriplet(tt.args.suit, tt.args.rank)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTriplet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKanzi_closed(t *testing.T) {
	type args struct {
		suit byte
		rank uint8
	}
	tests := []struct {
		name    string
		args    args
		want    Kanzi_closed
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewKanzi_closed(tt.args.suit, tt.args.rank)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKanzi_closed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKanzi_closed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKanzi_open(t *testing.T) {
	type args struct {
		suit byte
		rank uint8
	}
	tests := []struct {
		name    string
		args    args
		want    Kanzi_open
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewKanzi_open(tt.args.suit, tt.args.rank)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKanzi_open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKanzi_open() = %v, want %v", got, tt.want)
			}
		})
	}
}
