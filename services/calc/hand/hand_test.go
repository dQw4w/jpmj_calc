package hand

import (
	"reflect"
	"testing"
)

func TestLen(t *testing.T) {
	type args struct {
		hand Hand
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Len(tt.args.hand); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertStrToHand(t *testing.T) {
	type args struct {
		str_hand string
	}
	tests := []struct {
		name    string
		args    args
		want    Hand
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertStrToHand(tt.args.str_hand)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertStrToHand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertStrToHand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		arr  []uint8
		low  int
		high int
	}
	tests := []struct {
		name  string
		args  args
		want  []uint8
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := partition(tt.args.arr, tt.args.low, tt.args.high)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("partition() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		arr  []uint8
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.arr, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSortStart(t *testing.T) {
	type args struct {
		arr []uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSortStart(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSortStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortColor(t *testing.T) {
	type args struct {
		arr []uint8
	}
	tests := []struct {
		name  string
		args  args
		want  []uint8
		want1 uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SortColor(tt.args.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortColor() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SortColor() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSortAndReturnAkadora(t *testing.T) {
	type args struct {
		hand *Hand
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortAndReturnAkadora(tt.args.hand); got != tt.want {
				t.Errorf("SortAndReturnAkadora() = %v, want %v", got, tt.want)
			}
		})
	}
}
