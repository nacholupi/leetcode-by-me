package str

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty case",
			args: args{
				s: "",
				t: "",
			},
			want: false,
		},
		{
			name: "diff length case",
			args: args{
				s: "abc",
				t: "abcd",
			},
			want: false,
		},
		{
			name: "anagram case",
			args: args{
				s: "abc",
				t: "cba",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
