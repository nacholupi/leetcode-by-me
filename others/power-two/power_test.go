package powertwo

import (
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "zero case",
			args: args{n: 0},
			want: false,
		},
		{
			name: "one case",
			args: args{n: 1},
			want: true,
		},
		{
			name: "non power case, 17",
			args: args{n: 17},
			want: false,
		},
		{
			name: "non power case, 3",
			args: args{n: 3},
			want: false,
		},
		{
			name: "power case, 16",
			args: args{n: 16},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("IsPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
