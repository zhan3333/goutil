package util_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	util "github.com/zhan3333/goutil"
)

func TestSlice_Contains(t *testing.T) {
	type fields struct {
		slice []int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "int contains",
			fields: fields{
				slice: []int{1, 2, 3},
			},
			args: args{
				v: 1,
			},
			want: true,
		},
		{
			name: "not contains",
			fields: fields{
				slice: []int{1, 2, 3},
			},
			args: args{
				v: 4,
			},
			want: false,
		},
		{
			name: "empty slice",
			fields: fields{
				slice: []int{},
			},
			args: args{
				v: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := util.NewSlice(tt.fields.slice)
			assert.Equalf(t, tt.want, s.Contains(tt.args.v), "Contains(%v)", tt.args.v)
		})
	}
}

func TestSlice_ContainsCount(t *testing.T) {
	type fields struct {
		slice []string
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "empty slice",
			fields: fields{
				slice: []string{},
			},
			args: args{
				v: "1",
			},
			want: 0,
		},
		{
			name: "contains",
			fields: fields{
				slice: []string{"1", "2", "3"},
			},
			args: args{
				v: "1",
			},
			want: 1,
		},
		{
			name: "contains multiple",
			fields: fields{
				slice: []string{"1", "1", "3"},
			},
			args: args{
				v: "1",
			},
			want: 2,
		},
		{
			name: "not contains",
			fields: fields{
				slice: []string{"1", "1", "3"},
			},
			args: args{
				v: "4",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := util.NewSlice(tt.fields.slice)
			assert.Equalf(t, tt.want, s.ContainsCount(tt.args.v), "ContainsCount(%v)", tt.args.v)
		})
	}
}

func TestSlice_Copy(t *testing.T) {
	type fields struct {
		slice []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "copy",
			fields: fields{
				slice: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := util.NewSlice(tt.fields.slice)
			assert.True(t, util.Equal(s.Slice(), tt.want), "Equal(%v)", tt.want)
		})
	}
}

func TestSlice_Diff(t *testing.T) {
	type fields struct {
		slice []int
	}
	type args struct {
		c2 []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "diff: equal",
			fields: fields{
				slice: []int{1, 2, 3},
			},
			args: args{
				c2: []int{1, 2, 3},
			},
			want: []int{},
		},
		{
			name: "diff1",
			fields: fields{
				slice: []int{1, 2, 3},
			},
			args: args{
				c2: []int{1},
			},
			want: []int{2, 3},
		},
		{
			name: "diff2",
			fields: fields{
				slice: []int{1, 2, 3},
			},
			args: args{
				c2: []int{1, 4},
			},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := util.NewSlice(tt.fields.slice)
			diff := s.Diff(util.NewSlice(tt.args.c2))
			assert.Equal(t, tt.want, diff.Slice())
		})
	}
}

func TestSlice_Each(t *testing.T) {
	type fields struct {
		slice []int
	}
	type args struct {
		f func(int) int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "each",
			fields: fields{
				slice: []int{1, 2, 3},
			},
			args: args{
				f: func(a int) int {
					return a
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := util.NewSlice(tt.fields.slice)
			assert.Equalf(t, tt.want, s.Each(tt.args.f).Slice(), "Each()")
		})
	}
}

func TestSlice_Empty(t *testing.T) {
	type fields struct {
		slice []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty=true",
			fields: fields{
				slice: []int{},
			},
			want: true,
		},
		{
			name: "empty=false",
			fields: fields{
				slice: []int{1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := util.NewSlice(tt.fields.slice)
			assert.Equalf(t, tt.want, s.Empty(), "Empty()")
		})
	}
}

func TestSlice_Filter(t *testing.T) {
	type fields struct {
		slice []int
	}
	type args struct {
		f func(int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "filter",
			fields: fields{
				slice: []int{1, 2, 3, 4},
			},
			args: args{
				f: func(a int) bool {
					return a%2 == 0
				},
			},
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := util.NewSlice(tt.fields.slice)
			assert.Equalf(t, tt.want, s.Filter(tt.args.f).Slice(), "Filter(f)")
		})
	}
}

func TestSlice_First(t *testing.T) {
	type fields struct {
		slice []int
	}
	tests := []struct {
		name   string
		fields fields
		want   *int
	}{
		{
			name: "first",
			fields: fields{
				slice: []int{1, 2, 3},
			},
			want: func() *int {
				var a = 1
				return &a
			}(),
		},
		{
			name: "empty",
			fields: fields{
				slice: []int{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := util.NewSlice(tt.fields.slice)
			if tt.want == nil {
				assert.Nil(t, s.First(), "First()")
			} else {
				assert.Equalf(t, *tt.want, *s.First(), "First()")
			}
		})
	}
}

func TestSlice_Index(t *testing.T) {
	type fields struct {
		slice []int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *int
	}{
		{
			name: "index",
			fields: fields{
				slice: []int{1, 2, 3},
			},
			args: args{
				i: 0,
			},
			want: func() *int {
				var a = 1
				return &a
			}(),
		},
		{
			name: "index out range",
			fields: fields{
				slice: []int{1, 2, 3},
			},
			args: args{
				i: 3,
			},
			want: nil,
		},
		{
			name: "index empty slice",
			fields: fields{
				slice: []int{},
			},
			args: args{
				i: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := util.NewSlice(tt.fields.slice)
			if tt.want == nil {
				assert.Nil(t, s.Index(tt.args.i), "Index()")
			} else {
				assert.Equalf(t, *tt.want, *s.Index(tt.args.i), "Index()")
			}
		})
	}
}

func TestSlice_Last(t *testing.T) {
	type fields struct {
		slice []int
	}
	tests := []struct {
		name   string
		fields fields
		want   *int
	}{
		{
			name: "last",
			fields: fields{
				slice: []int{1, 2, 3},
			},
			want: PInt(3),
		},
		{
			name: "empty",
			fields: fields{
				slice: []int{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := util.NewSlice(tt.fields.slice)
			if tt.want == nil {
				assert.Nil(t, s.Last(), "Last()")
			} else {
				assert.Equalf(t, *tt.want, *s.Last(), "Last()")
			}
		})
	}
}

func TestSlice_Len(t *testing.T) {
	type fields struct {
		slice []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "len1",
			fields: fields{
				slice: []int{1, 2, 3},
			},
			want: 3,
		},
		{
			name: "len2",
			fields: fields{
				slice: []int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := util.NewSlice(tt.fields.slice)
			assert.Equalf(t, tt.want, s.Len(), "Len()")
		})
	}
}

func TestSlice_Map(t *testing.T) {
	type fields struct {
		slice []int
	}
	type args struct {
		f func(int) int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "map",
			fields: fields{
				slice: []int{1, 2, 3},
			},
			args: args{
				f: func(i int) int {
					return i * i
				},
			},
			want: []int{1, 4, 9},
		},
		{
			name: "map empty slice",
			fields: fields{
				slice: []int{},
			},
			args: args{
				f: func(i int) int {
					return i * i
				},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := util.NewSlice(tt.fields.slice)
			assert.Equalf(t, tt.want, s.Map(tt.args.f).Slice(), "Map()")
		})
	}
}

func TestSlice_Merge(t *testing.T) {
	type fields struct {
		slice []int
	}
	type args struct {
		ss []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "merge",
			fields: fields{
				slice: []int{1},
			},
			args: args{
				ss: []int{2},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := util.NewSlice(tt.fields.slice)
			s2 := util.NewSlice(tt.args.ss)
			assert.Equal(t, tt.want, s.Merge(s2).Slice())
		})
	}
}

func TestSlice_Push(t *testing.T) {
	type fields struct {
		slice *util.Slice[int]
	}
	type args struct {
		vs []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *util.Slice[int]
	}{
		{
			name: "push",
			fields: fields{
				slice: util.NewSlice([]int{1, 2, 3}),
			},
			args: args{
				vs: []int{4, 5},
			},
			want: util.NewSlice([]int{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, tt.fields.slice.Push(tt.args.vs...).Equal(tt.want))
		})
	}
}

func TestSlice_Pop(t *testing.T) {
	type fields struct {
		slice *util.Slice[int]
	}
	tests := []struct {
		name   string
		fields fields
		wantA  *util.Slice[int]
		wantB  *int
	}{
		{
			name: "pop",
			fields: fields{
				slice: util.NewSlice([]int{1, 2, 3}),
			},
			wantA: util.NewSlice([]int{1, 2}),
			wantB: PInt(3),
		},
		{
			name: "pop empty slice",
			fields: fields{
				slice: util.NewSlice([]int{}),
			},
			wantA: util.NewSlice([]int{}),
			wantB: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.fields.slice.Pop()
			assert.True(t, tt.fields.slice.Equal(tt.wantA))
			if tt.wantB == nil {
				assert.Nil(t, v)
			} else {
				assert.Equal(t, *tt.wantB, *v)
			}
		})
	}
}

func TestSlice_Reject(t *testing.T) {
	type fields struct {
		slice *util.Slice[int]
	}
	type args struct {
		f func(int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *util.Slice[int]
	}{
		{
			name: "reject",
			fields: fields{
				slice: util.NewSlice([]int{1, 2, 3}),
			},
			args: args{
				f: func(i int) bool {
					return i%2 == 0
				},
			},
			want: util.NewSlice([]int{1, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.slice.Reject(tt.args.f)
			t.Logf("%+v", tt.fields.slice.Slice())
			t.Logf("%+v", tt.want.Slice())
			assert.Equal(t, tt.want.Slice(), tt.fields.slice.Slice())

			assert.True(t, tt.fields.slice.Reject(tt.args.f).Equal(tt.want))
		})
	}
}

func TestSlice_Reset(t *testing.T) {
	type fields struct {
		slice *util.Slice[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   *util.Slice[int]
	}{
		{
			name: "reset",
			fields: fields{
				slice: util.NewSlice([]int{1, 2, 3}),
			},
			want: util.NewSlice([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, tt.fields.slice.Reset().Empty())
		})
	}
}

func TestSlice_Reverse(t *testing.T) {
	type fields struct {
		slice *util.Slice[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   *util.Slice[int]
	}{
		{
			name: "reverse",
			fields: fields{
				slice: util.NewSlice([]int{1, 2, 3}),
			},
			want: util.NewSlice([]int{3, 2, 1}),
		},
		{
			name: "reverse empty slice",
			fields: fields{
				slice: util.NewSlice([]int{}),
			},
			want: util.NewSlice([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, tt.fields.slice.Reverse().Equal(tt.want))
		})
	}
}

func TestSlice_Set(t *testing.T) {
	type fields struct {
		slice *util.Slice[int]
	}
	type args struct {
		vs []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *util.Slice[int]
	}{
		{
			name: "set",
			fields: fields{
				slice: util.NewSlice([]int{}),
			},
			args: args{
				vs: []int{1, 2},
			},
			want: util.NewSlice([]int{1, 2}),
		},
		{
			name: "set no empty slice",
			fields: fields{
				slice: util.NewSlice([]int{1, 2}),
			},
			args: args{
				vs: []int{3, 4},
			},
			want: util.NewSlice([]int{3, 4}),
		},
		{
			name: "set empty slice",
			fields: fields{
				slice: util.NewSlice([]int{1, 2}),
			},
			args: args{
				vs: []int{},
			},
			want: util.NewSlice([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, tt.fields.slice.Set(tt.args.vs).Equal(tt.want))
		})
	}
}

func TestSlice_Slice(t *testing.T) {
	type fields struct {
		slice *util.Slice[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "slice",
			fields: fields{
				slice: util.NewSlice([]int{1, 2, 3}),
			},
			want: []int{1, 2, 3},
		},
		{
			name: "empty slice",
			fields: fields{
				slice: util.NewSlice([]int{}),
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.fields.slice.Slice(), "Slice()")
		})
	}
}

func TestSlice_Unique(t *testing.T) {
	type fields struct {
		slice *util.Slice[int]
	}
	tests := []struct {
		name   string
		fields fields
		want   *util.Slice[int]
	}{
		{
			name: "unique",
			fields: fields{
				slice: util.NewSlice([]int{1, 2, 3, 1, 2, 3}),
			},
			want: util.NewSlice([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, tt.fields.slice.Unique().Equal(tt.want))
		})
	}
}

func PInt(i int) *int {
	return &i
}
