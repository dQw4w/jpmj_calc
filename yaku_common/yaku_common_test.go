package yaku_common

import (
	"jpmj_calc/combination"
	"jpmj_calc/win"
	"testing"
)

func TestNon_Yakuman_Special(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Non_Yakuman_Special(tt.args.cw)
			if got != tt.want {
				t.Errorf("Non_Yakuman_Special() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Non_Yakuman_Special() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMenchinTsumo(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MenchinTsumo(tt.args.cw)
			if got != tt.want {
				t.Errorf("MenchinTsumo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MenchinTsumo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTanyao(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Tanyao(tt.args.cw)
			if got != tt.want {
				t.Errorf("Tanyao() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Tanyao() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestYakuhai_Selfwind(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Yakuhai_Selfwind(tt.args.cw)
			if got != tt.want {
				t.Errorf("Yakuhai_Selfwind() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Yakuhai_Selfwind() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestYakuhai_Fieldwind(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Yakuhai_Fieldwind(tt.args.cw)
			if got != tt.want {
				t.Errorf("Yakuhai_Fieldwind() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Yakuhai_Fieldwind() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestYakuhai_Sangen(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Yakuhai_Sangen(tt.args.cw)
			if got != tt.want {
				t.Errorf("Yakuhai_Sangen() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Yakuhai_Sangen() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPinhu(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Pinhu(tt.args.cw)
			if got != tt.want {
				t.Errorf("Pinhu() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Pinhu() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOnePekoandTwoPeko(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := OnePekoandTwoPeko(tt.args.cw)
			if got != tt.want {
				t.Errorf("OnePekoandTwoPeko() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("OnePekoandTwoPeko() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestThreeSameTrp(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ThreeSameTrp(tt.args.cw)
			if got != tt.want {
				t.Errorf("ThreeSameTrp() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ThreeSameTrp() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestThreeSameStra(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ThreeSameStra(tt.args.cw)
			if got != tt.want {
				t.Errorf("ThreeSameStra() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ThreeSameStra() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestThreeKanzi(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ThreeKanzi(tt.args.cw)
			if got != tt.want {
				t.Errorf("ThreeKanzi() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ThreeKanzi() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToitoi(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Toitoi(tt.args.cw)
			if got != tt.want {
				t.Errorf("Toitoi() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Toitoi() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestThreeConcealedTrp(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ThreeConcealedTrp(tt.args.cw)
			if got != tt.want {
				t.Errorf("ThreeConcealedTrp() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ThreeConcealedTrp() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHonOldHead(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := HonOldHead(tt.args.cw)
			if got != tt.want {
				t.Errorf("HonOldHead() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HonOldHead() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSmallSangen(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SmallSangen(tt.args.cw)
			if got != tt.want {
				t.Errorf("SmallSangen() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SmallSangen() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestChanTa(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ChanTa(tt.args.cw)
			if got != tt.want {
				t.Errorf("ChanTa() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ChanTa() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOneDragon(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := OneDragon(tt.args.cw)
			if got != tt.want {
				t.Errorf("OneDragon() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("OneDragon() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSomete(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Somete(tt.args.cw)
			if got != tt.want {
				t.Errorf("Somete() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Somete() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestYakuman_Special(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Yakuman_Special(tt.args.cw)
			if got != tt.want {
				t.Errorf("Yakuman_Special() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Yakuman_Special() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBigSangen(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BigSangen(tt.args.cw)
			if got != tt.want {
				t.Errorf("BigSangen() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BigSangen() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFourConcealedTrp(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FourConcealedTrp(tt.args.cw)
			if got != tt.want {
				t.Errorf("FourConcealedTrp() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FourConcealedTrp() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOnlyZi(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := OnlyZi(tt.args.cw)
			if got != tt.want {
				t.Errorf("OnlyZi() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("OnlyZi() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOnlyGreen(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := OnlyGreen(tt.args.cw)
			if got != tt.want {
				t.Errorf("OnlyGreen() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("OnlyGreen() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOnlyOld(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := OnlyOld(tt.args.cw)
			if got != tt.want {
				t.Errorf("OnlyOld() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("OnlyOld() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFourWinds(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FourWinds(tt.args.cw)
			if got != tt.want {
				t.Errorf("FourWinds() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FourWinds() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFourKanzi(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FourKanzi(tt.args.cw)
			if got != tt.want {
				t.Errorf("FourKanzi() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FourKanzi() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNineGates(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NineGates(tt.args.cw)
			if got != tt.want {
				t.Errorf("NineGates() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("NineGates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestYakuman_Check(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Yakuman_Check(tt.args.cw)
			if got != tt.want {
				t.Errorf("Yakuman_Check() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Yakuman_Check() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCalculateYaku(t *testing.T) {
	type args struct {
		cw win.Common_Win
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		{
			name: "摸斷三色",
			args: args{
				cw: win.Common_Win{
					MenziList: [4]combination.Menzi{
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 2,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'p',
							Rank: 2,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 's',
							Rank: 2,
							Furo: false,
						},
						combination.Menzi{
							Type: 'T',
							Suit: 'm',
							Rank: 6,
							Furo: false,
						},
					},
					Eye: combination.Pair{
						Suit: 's',
						Rank: 8,
						Furo: false,
					},
					Tsumo:     true,
					Menchin:   true,
					SelfWind:  1,
					FieldWind: 1,
				},
			},
			want:  4,
			want1: "",
		},
		{
			name: "摸斷三色(副露)",
			args: args{
				cw: win.Common_Win{
					MenziList: [4]combination.Menzi{
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 2,
							Furo: true,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'p',
							Rank: 2,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 's',
							Rank: 2,
							Furo: false,
						},
						combination.Menzi{
							Type: 'T',
							Suit: 'm',
							Rank: 6,
							Furo: false,
						},
					},
					Eye: combination.Pair{
						Suit: 's',
						Rank: 8,
						Furo: false,
					},
					Tsumo:     true,
					Menchin:   false,
					SelfWind:  1,
					FieldWind: 1,
				},
			},
			want:  2,
			want1: "",
		},

		{
			name: "平和一peko",
			args: args{
				cw: win.Common_Win{
					MenziList: [4]combination.Menzi{
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 2,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 2,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 's',
							Rank: 3,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'p',
							Rank: 6,
							Furo: false,
						},
					},
					Eye: combination.Pair{
						Suit: 'z',
						Rank: 2,
						Furo: false,
					},
					//Tsumo:     true,
					Menchin:      true,
					SelfWind:     1,
					FieldWind:    1,
					Win_Com_IDX:  0,
					Win_Tile_IDX: 0,
				},
			},
			want:  2,
			want1: "",
		},
		{
			name: "平和fail(yakuhai eye)",
			args: args{
				cw: win.Common_Win{
					MenziList: [4]combination.Menzi{
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 1,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 2,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 's',
							Rank: 3,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'p',
							Rank: 6,
							Furo: false,
						},
					},
					Eye: combination.Pair{
						Suit: 'z',
						Rank: 2,
						Furo: false,
					},
					//Tsumo:     true,
					Menchin:      true,
					SelfWind:     1,
					FieldWind:    2,
					Win_Com_IDX:  0,
					Win_Tile_IDX: 0,
				},
			},
			want:  0,
			want1: "",
		},
		{
			name: "平和fail(middle)",
			args: args{
				cw: win.Common_Win{
					MenziList: [4]combination.Menzi{
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 1,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 2,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 's',
							Rank: 3,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'p',
							Rank: 6,
							Furo: false,
						},
					},
					Eye: combination.Pair{
						Suit: 'z',
						Rank: 3,
						Furo: false,
					},
					//Tsumo:     true,
					Menchin:      true,
					SelfWind:     1,
					FieldWind:    2,
					Win_Com_IDX:  0,
					Win_Tile_IDX: 1,
				},
			},
			want:  0,
			want1: "",
		},
		{
			name: "平和fail(123)",
			args: args{
				cw: win.Common_Win{
					MenziList: [4]combination.Menzi{
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 1,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 2,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 's',
							Rank: 3,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'p',
							Rank: 6,
							Furo: false,
						},
					},
					Eye: combination.Pair{
						Suit: 'z',
						Rank: 3,
						Furo: false,
					},
					//Tsumo:     true,
					Menchin:      true,
					SelfWind:     1,
					FieldWind:    2,
					Win_Com_IDX:  0,
					Win_Tile_IDX: 2,
				},
			},
			want:  0,
			want1: "",
		},
		{
			name: "平和fail(789)",
			args: args{
				cw: win.Common_Win{
					MenziList: [4]combination.Menzi{
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 7,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 2,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 's',
							Rank: 3,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'p',
							Rank: 6,
							Furo: false,
						},
					},
					Eye: combination.Pair{
						Suit: 'z',
						Rank: 3,
						Furo: false,
					},
					//Tsumo:     true,
					Menchin:      true,
					SelfWind:     1,
					FieldWind:    2,
					Win_Com_IDX:  0,
					Win_Tile_IDX: 0,
				},
			},
			want:  0,
			want1: "",
		},
		{
			name: "平和fail(eye)",
			args: args{
				cw: win.Common_Win{
					MenziList: [4]combination.Menzi{
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 6,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 2,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 's',
							Rank: 3,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'p',
							Rank: 4,
							Furo: false,
						},
					},
					Eye: combination.Pair{
						Suit: 'z',
						Rank: 3,
						Furo: false,
					},
					//Tsumo:     true,
					Menchin:      true,
					SelfWind:     1,
					FieldWind:    2,
					Win_Com_IDX:  4,
					Win_Tile_IDX: 0,
				},
			},
			want:  0,
			want1: "",
		},
		{
			name: "threetrpoldthreeconcealedyakuhaithreekan",
			args: args{
				cw: win.Common_Win{
					MenziList: [4]combination.Menzi{
						combination.Menzi{
							Type: 'O',
							Suit: 'm',
							Rank: 1,
							Furo: true,
						},
						combination.Menzi{
							Type: 'T',
							Suit: 'p',
							Rank: 1,
							Furo: false,
						},
						combination.Menzi{
							Type: 'C',
							Suit: 's',
							Rank: 1,
							Furo: false,
						},
						combination.Menzi{
							Type: 'C',
							Suit: 'z',
							Rank: 3,
							Furo: false,
						},
					},
					Eye: combination.Pair{
						Suit: 'z',
						Rank: 5,
						Furo: false,
					},
					Tsumo:        true,
					Menchin:      false,
					SelfWind:     1,
					FieldWind:    3,
					Win_Com_IDX:  1,
					Win_Tile_IDX: 0,
				},
			},
			want:  11,
			want1: "",
		},

		{
			name: "混一一氣",
			args: args{
				cw: win.Common_Win{
					MenziList: [4]combination.Menzi{
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 1,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 4,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 7,
							Furo: false,
						},
						combination.Menzi{
							Type: 'T',
							Suit: 'z',
							Rank: 3,
							Furo: false,
						},
					},
					Eye: combination.Pair{
						Suit: 'z',
						Rank: 6,
						Furo: false,
					},
					Tsumo:        true,
					Menchin:      true,
					SelfWind:     1,
					FieldWind:    3,
					Win_Com_IDX:  1,
					Win_Tile_IDX: 0,
				},
			},
			want:  7,
			want1: "",
		},
		{
			name: "清一全代",
			args: args{
				cw: win.Common_Win{
					MenziList: [4]combination.Menzi{
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 1,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 1,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 7,
							Furo: false,
						},
						combination.Menzi{
							Type: 'S',
							Suit: 'm',
							Rank: 7,
							Furo: false,
						},
					},
					Eye: combination.Pair{
						Suit: 'z',
						Rank: 1,
						Furo: false,
					},
					Tsumo:        true,
					Menchin:      true,
					DoubleReach:  true,
					RinShan:      true,
					HaiTei:       true,
					ChanKan:      true,
					SelfWind:     1,
					FieldWind:    3,
					Win_Com_IDX:  1,
					Win_Tile_IDX: 2,
				},
			},
			want:  1 + 6 + 3 + 3,
			want1: "",
		},

		// TODO: Add test cases.
		// TODO: test cases are type Common_Win struct :
		/*MenziList    [4]combination.Menzi
		Eye          combination.Pair
		Win_Com_IDX  int // 4 represents eye
		Win_Tile_IDX int
		Tsumo        bool
		Menchin      bool
		SelfWind     uint8 //1234 ESWN
		FieldWind    uint8 //1234

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

		Akadora       int
		Motedora_suit byte
		Motedora_rank uint8

		Uradora_suit byte
		Uradora_rank uint8*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CalculateYaku(tt.args.cw)
			if got != tt.want {
				t.Errorf("CalculateYaku() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CalculateYaku() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
