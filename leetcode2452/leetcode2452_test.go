package leetcode2452

import (
	"reflect"
	"testing"
)

func Test_twoEditWords(t *testing.T) {
	type args struct {
		queries    []string
		dictionary []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example 1",
			args: args{
				queries:    []string{"word", "note", "ants", "wood"},
				dictionary: []string{"wood", "joke", "moat"},
			},
			want: []string{"word", "note", "wood"},
		},
		{
			name: "example 2",
			args: args{
				queries:    []string{"yes"},
				dictionary: []string{"not"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoEditWords(tt.args.queries, tt.args.dictionary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoEditWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
