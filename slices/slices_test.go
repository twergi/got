package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	type args struct {
		v []int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1, 0",
			args: args{
				v: []int{1},
				n: 0,
			},
			want: nil,
		},
		{
			name: "0, 1",
			args: args{
				v: []int{},
				n: 1,
			},
			want: nil,
		},
		{
			name: "4, 3",
			args: args{
				v: []int{1, 2, 3, 4},
				n: 3,
			},
			want: [][]int{
				{1, 2, 3},
				{4},
			},
		},
		{
			name: "4, 1",
			args: args{
				v: []int{1, 2, 3, 4},
				n: 1,
			},
			want: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
		{
			name: "10, 2",
			args: args{
				v: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				n: 2,
			},
			want: [][]int{
				{0, 1},
				{2, 3},
				{4, 5},
				{6, 7},
				{8, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Slice(tt.args.v, tt.args.n)

			assert.Equal(t, tt.want, got)
		})
	}
}
