package firstunique

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "not found empty case",
			args: args{
				s: "",
			},
			want: -1,
		}, {
			name: "one char in string",
			args: args{
				s: "s",
			},
			want: 0,
		}, {
			name: "first char in string is not repeated",
			args: args{
				s: "sa",
			},
			want: 0,
		}, {
			name: "sec char in string is not repeated",
			args: args{
				s: "sNs",
			},
			want: 1,
		}, {
			name: "first char in string is not repeated",
			args: args{
				s: "sabcdefg",
			},
			want: 0,
		}, {
			name: "every char is repeated",
			args: args{
				s: "sabcdefgsabcdefg",
			},
			want: -1,
		}, {
			name: "every char is repeated",
			args: args{
				s: "sabcdefgsabcdefg",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
