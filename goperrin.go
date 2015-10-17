/*
Copyright 2015 Ryan Richard

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package goperrin

func Perrin() func() int {
	a, b, c := 3, 0, 2

	return func() int {
		p := a
		a, b, c = b, c, a+b

		return p

	}
}

func PerrinMax(max int) func() int {
	a, b, c := 3, 0, 2

	return func() int {
		if a >= max {
			return a
		}

		a, b, c = b, c, a+b
		return a
	}
}

func PerrinReset(max int) func() int {
	a, b, c := 3, 0, 2

	return func() int {
		if a >= max {
			p := a
			a, b, c = 3, 0, 2
			return p
		}

		p := a
		a, b, c = b, c, a+b
		return p
	}
}
