package ransom

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty case",
			args: args{
				ransomNote: "",
				magazine:   "",
			},
			want: false,
		},
		{
			name: "single case",
			args: args{
				ransomNote: "a",
				magazine:   "a",
			},
			want: true,
		},
		{
			name: "magazine with less characters than note case",
			args: args{
				ransomNote: "aa",
				magazine:   "a",
			},
			want: false,
		},
		{
			name: "can be constructed case",
			args: args{
				ransomNote: "b",
				magazine:   "ab",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
