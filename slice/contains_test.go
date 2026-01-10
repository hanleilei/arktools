// Copyright 2026 hanleilei
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  4,
			name: "dst exist",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  3,
			name: "dst not exist",
		},
		{
			want: false,
			src:  []int{},
			dst:  4,
			name: "length of src is 0",
		},
		{
			want: false,
			dst:  4,
			name: "src nil",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Contains[int](test.src, test.dst))
		})
	}
}

func TestContainsFunc(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  4,
			name: "dst exist",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  3,
			name: "dst not exist",
		},
		{
			want: false,
			src:  []int{},
			dst:  4,
			name: "length of src is 0",
		},
		{
			want: false,
			dst:  4,
			name: "src nil",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsFunc[int](test.src, func(src int) bool {
				return src == test.dst
			}))
		})
	}
}

func TestContainsFunc_Struct(t *testing.T) {
	type S struct{ A, B int }
	src := []S{{1, 2}, {3, 4}, {5, 6}}
	assert.True(t, ContainsFunc(src, func(s S) bool { return s.A == 3 }))
	assert.False(t, ContainsFunc(src, func(s S) bool { return s.B == 7 }))
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  []int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 6},
			name: "exist two ele",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{7, 0},
			name: "not exist the same",
		},
		{
			want: true,
			src:  []int{1, 1, 8},
			dst:  []int{1, 1},
			name: "exist two same ele",
		},
		{
			want: false,
			src:  []int{},
			dst:  []int{1},
			name: "length of src is 0",
		},
		{
			want: false,
			dst:  []int{1},
			name: "src nil",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAny[int](test.src, test.dst))
		})
	}
}

func TestContainsAnyFunc(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  []int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 6},
			name: "exist two ele",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{7, 0},
			name: "not exist the same",
		},
		{
			want: true,
			src:  []int{1, 1, 8},
			dst:  []int{1, 1},
			name: "exist two same ele",
		},
		{
			want: false,
			src:  []int{},
			dst:  []int{1},
			name: "length of src is 0",
		},
		{
			want: false,
			dst:  []int{1},
			name: "src nil",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAnyFunc[int](test.src, test.dst, func(src, dst int) bool {
				return src == dst
			}))
		})
	}
}

func TestContainsAnyFunc_Struct(t *testing.T) {
	type S struct{ A, B int }
	src := []S{{1, 2}, {3, 4}, {5, 6}}
	dst := []S{{7, 8}, {3, 4}}
	assert.True(t, ContainsAnyFunc(src, dst, func(a, b S) bool { return a.A == b.A }))
	dst2 := []S{{7, 8}, {9, 10}}
	assert.False(t, ContainsAnyFunc(src, dst2, func(a, b S) bool { return a.A == b.A }))
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  []int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2},
			name: "src exist one not in dst",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2, 6, 7},
			name: "src not include the whole ele",
		},
		{
			want: false,
			src:  []int{},
			dst:  []int{1},
			name: "length of src is 0",
		},
		{
			want: true,
			src:  nil,
			dst:  []int{},
			name: "src nil dst empty",
		},
		{
			want: true,
			src:  nil,
			name: "src and dst nil",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAll[int](test.src, test.dst))
		})
	}
}

func TestContainsAllFunc(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  []int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2},
			name: "src exist one not in dst",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2, 6, 7},
			name: "src not include the whole ele",
		},
		{
			want: false,
			src:  []int{},
			dst:  []int{1},
			name: "length of src is 0",
		},
		{
			want: true,
			src:  nil,
			dst:  []int{},
			name: "src nil dst empty",
		},
		{
			want: true,
			src:  nil,
			name: "src and dst nil",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAllFunc[int](test.src, test.dst, func(src, dst int) bool {
				return src == dst
			}))
		})
	}
}

func TestContainsAllFunc_Struct(t *testing.T) {
	type S struct{ A, B int }
	src := []S{{1, 2}, {3, 4}, {5, 6}}
	dst := []S{{3, 0}, {1, 0}}
	assert.True(t, ContainsAllFunc(src, dst, func(a, b S) bool { return a.A == b.A }))
	dst2 := []S{{3, 0}, {7, 0}}
	assert.False(t, ContainsAllFunc(src, dst2, func(a, b S) bool { return a.A == b.A }))
}

func TestContains_String(t *testing.T) {
	src := []string{"a", "b", "c"}
	assert.True(t, Contains(src, "b"))
	assert.False(t, Contains(src, "d"))
	assert.False(t, Contains([]string{}, "a"))
	assert.False(t, Contains(nil, "a"))
}

func TestContainsFunc_String(t *testing.T) {
	src := []string{"foo", "bar", "baz"}
	assert.True(t, ContainsFunc(src, func(s string) bool { return s == "bar" }))
	assert.False(t, ContainsFunc(src, func(s string) bool { return s == "qux" }))
	assert.False(t, ContainsFunc([]string{}, func(s string) bool { return s == "foo" }))
	assert.False(t, ContainsFunc(nil, func(s string) bool { return s == "foo" }))
}

func TestContainsAny_String(t *testing.T) {
	src := []string{"a", "b", "c"}
	dst := []string{"d", "b"}
	assert.True(t, ContainsAny(src, dst))
	dst2 := []string{"x", "y"}
	assert.False(t, ContainsAny(src, dst2))
	assert.False(t, ContainsAny([]string{}, []string{"a"}))
	assert.False(t, ContainsAny(nil, []string{"a"}))
}

func TestContainsAnyFunc_String(t *testing.T) {
	src := []string{"foo", "bar", "baz"}
	dst := []string{"baz", "qux"}
	assert.True(t, ContainsAnyFunc(src, dst, func(a, b string) bool { return a == b }))
	dst2 := []string{"qux", "xyz"}
	assert.False(t, ContainsAnyFunc(src, dst2, func(a, b string) bool { return a == b }))
	assert.False(t, ContainsAnyFunc([]string{}, []string{"foo"}, func(a, b string) bool { return a == b }))
	assert.False(t, ContainsAnyFunc(nil, []string{"foo"}, func(a, b string) bool { return a == b }))
}

func TestContainsAll_String(t *testing.T) {
	src := []string{"a", "b", "c"}
	dst := []string{"a", "c"}
	assert.True(t, ContainsAll(src, dst))
	dst2 := []string{"a", "d"}
	assert.False(t, ContainsAll(src, dst2))
	assert.False(t, ContainsAll([]string{}, []string{"a"}))
	assert.True(t, ContainsAll[string](nil, []string{}))
	assert.True(t, ContainsAll[string](nil, nil))
}

func TestContainsAllFunc_String(t *testing.T) {
	src := []string{"foo", "bar", "baz"}
	dst := []string{"foo", "baz"}
	assert.True(t, ContainsAllFunc(src, dst, func(a, b string) bool { return a == b }))
	dst2 := []string{"foo", "qux"}
	assert.False(t, ContainsAllFunc(src, dst2, func(a, b string) bool { return a == b }))
	assert.False(t, ContainsAllFunc([]string{}, []string{"foo"}, func(a, b string) bool { return a == b }))
	assert.True(t, ContainsAllFunc(nil, []string{}, func(a, b string) bool { return a == b }))
	assert.True(t, ContainsAllFunc(nil, nil, func(a, b string) bool { return a == b }))
}

func TestContains_Struct_Empty(t *testing.T) {
	type S struct{ A, B int }
	src := []S{}
	dst := S{1, 2}
	assert.False(t, Contains(src, dst))
	assert.False(t, ContainsFunc(src, func(s S) bool { return s.A == 1 }))
	assert.False(t, ContainsAny(src, []S{dst}))
	assert.False(t, ContainsAnyFunc(src, []S{dst}, func(a, b S) bool { return a.A == b.A }))
	assert.False(t, ContainsAll(src, []S{dst}))
	assert.False(t, ContainsAllFunc(src, []S{dst}, func(a, b S) bool { return a.A == b.A }))
	assert.False(t, ContainsAll[S](nil, []S{dst}))
	assert.True(t, ContainsAll[S](nil, nil)) // 按照语法和实现实际结果为 true
}

func ExampleContains() {
	res := Contains[int]([]int{1, 2, 3}, 3)
	fmt.Println(res)
	// Output:
	// true
}

func ExampleContainsFunc() {
	res := ContainsFunc[int]([]int{1, 2, 3}, func(src int) bool {
		return src == 3
	})
	fmt.Println(res)
	// Output:
	// true
}

func ExampleContainsAll() {
	res := ContainsAll[int]([]int{1, 2, 3}, []int{3, 1})
	fmt.Println(res)
	res = ContainsAll[int]([]int{1, 2, 3}, []int{3, 1, 4})
	fmt.Println(res)
	// Output:
	// true
	// false
}

func ExampleContainsAllFunc() {
	res := ContainsAllFunc[int]([]int{1, 2, 3}, []int{3, 1}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	res = ContainsAllFunc[int]([]int{1, 2, 3}, []int{3, 1, 4}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	// Output:
	// true
	// false
}

func ExampleContainsAny() {
	res := ContainsAny[int]([]int{1, 2, 3}, []int{3, 6})
	fmt.Println(res)
	res = ContainsAny[int]([]int{1, 2, 3}, []int{4, 5, 9})
	fmt.Println(res)
	// Output:
	// true
	// false
}

func ExampleContainsAnyFunc() {
	res := ContainsAnyFunc[int]([]int{1, 2, 3}, []int{3, 1}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	res = ContainsAllFunc[int]([]int{1, 2, 3}, []int{4, 7, 6}, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res)
	// Output:
	// true
	// false
}
