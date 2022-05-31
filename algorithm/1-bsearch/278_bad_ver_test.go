package bsearch

import "testing"

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "four",
			args: args{
				n: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
