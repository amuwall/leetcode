package leetcode2129

import "testing"

func Test_capitalizeTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{title: "capiTalIze tHe titLe"},
			want: "Capitalize The Title",
		},
		{
			name: "example 2",
			args: args{title: "First leTTeR of EACH Word"},
			want: "First Letter of Each Word",
		},
		{
			name: "example 3",
			args: args{title: "i lOve leetcode"},
			want: "i Love Leetcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := capitalizeTitle(tt.args.title); got != tt.want {
				t.Errorf("capitalizeTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
