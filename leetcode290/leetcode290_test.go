package leetcode290

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				pattern: "abba",
				s:       "dog cat cat dog",
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				pattern: "abba",
				s:       "dog cat cat fish",
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				pattern: "aaaa",
				s:       "dog cat cat dog",
			},
			want: false,
		},
		{
			name: "example 4",
			args: args{
				pattern: "abba",
				s:       "dog dog dog dog",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
