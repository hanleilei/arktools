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

// calCapacity 根据当前切片容量和长度，判断是否需要收缩，并返回新容量和是否变化。
// 收缩策略：
// 1. 小容量（<=64）不收缩。
// 2. 大容量（>2048）且利用率低（容量/长度>=2）时按0.625收缩。
// 3. 中等容量（<=2048）且利用率极低（容量/长度>=4）时减半收缩。
// 4. 其余情况不收缩。
func calCapacity(c, l int) (int, bool) {
	if c <= 64 {
		return c, false
	}
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float64(c) * factor), true
	}
	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}
	return c, false
}

// Shrink 收缩切片容量，避免内存浪费。
// 仅在容量远大于实际长度时触发收缩，返回收缩后的新切片，否则返回原切片。
// 收缩后数据顺序不变，原切片不受影响。
func Shrink[T any](src []T) []T {
	if len(src) == 0 {
		return src
	}
	c, l := cap(src), len(src)
	n, changed := calCapacity(c, l)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}
