package leetcode2810

import "testing"

func Test_finalString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{s: "string"},
			want: "rtsng",
		},
		{
			name: "example 2",
			args: args{s: "poiinter"},
			want: "ponter",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalString(tt.args.s); got != tt.want {
				t.Errorf("finalString() = %v, want %v", got, tt.want)
			}
		})
	}
}
