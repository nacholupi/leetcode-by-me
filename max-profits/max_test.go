package maxprofits

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero case",
			args: args{prices: []int{}},
			want: 0,
		},
		{
			name: "pos profits case",
			args: args{prices: []int{2, 7}},
			want: 5,
		},
		{
			name: "neg profits case",
			args: args{prices: []int{7, 2}},
			want: 0,
		},
		{
			name: "pos profits case",
			args: args{prices: []int{7, 2, 10}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
