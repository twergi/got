package slices

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZipSeq(t *testing.T) {
	type args struct {
		v1 []int
		v2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "returns smallest 3",
			args: args{
				v1: []int{1, 3, 5},
				v2: []int{2, 4, 6, 8},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "returns smallest 0",
			args: args{
				v1: []int{},
				v2: []int{2, 4, 6, 8},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := slices.Collect(ZipSeq(tt.args.v1, tt.args.v2))

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSliceSeq(t *testing.T) {
	type args struct {
		vs []int
		n  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "4, 3",
			args: args{
				vs: []int{0, 1, 2, 3},
				n:  3,
			},
			want: [][]int{
				{0, 1, 2},
				{3},
			},
		},
		{
			name: "4, 2",
			args: args{
				vs: []int{0, 1, 2, 3},
				n:  2,
			},
			want: [][]int{
				{0, 1},
				{2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := slices.Collect(SliceSeq(slices.Values(tt.args.vs), tt.args.n))

			assert.Equal(t, tt.want, got)
		})
	}
}
