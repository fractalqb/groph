// Copyright 2022 Marcus Perlick
// This file is part of Go module git.fractalqb.de/fractalqb/groph
//
// groph is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// groph is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with groph.  If not, see <http://www.gnu.org/licenses/>.

package internal

func Slice[T any](s []T, l int) []T {
	if s == nil || cap(s) < l {
		return make([]T, l)
	}
	return s[:l]
}

// func CpSlice[T any](s []T, l int) []T {
// 	if s == nil || cap(s) < l {
// 		s2 := make([]T, l)
// 		copy(s2, s)
// 		return s2
// 	}
// 	i := len(s)
// 	s = s[:l]
// 	var zero T
// 	for i < l {
// 		s[i] = zero
// 	}
// 	return s
// }
