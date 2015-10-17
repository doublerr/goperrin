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

import "testing"

func TestPerrin(t *testing.T) {
	var perrinList [4]int
	masterList := [4]int{3, 0, 2, 3}

	p := Perrin()

	for i := 0; i < 4; i += 1 {
		perrinList[i] = p()
	}

	if perrinList != masterList {
		t.Errorf("perrinList == %v, want %v", perrinList, masterList)
	}
}

func TestPerrinMax(t *testing.T) {
	var perrinList [10]int
	max := 5

	p := PerrinMax(max)

	for i := 0; i < 10; i += 1 {
		perrinList[i] = p()
	}

	if perrinList[len(perrinList)-1] != max {
		t.Errorf("last element of perrinList[] != %v", max)
	}
}

func TestPerrinReset(t *testing.T) {
	var perrinList [9]int
	masterList := [9]int{3, 0, 2, 3, 2, 5, 3, 0, 2}

	p := PerrinReset(5)

	for i := 0; i < 9; i += 1 {
		perrinList[i] = p()
	}

	if perrinList != masterList {
		t.Errorf("perrinList == %v, want %v", perrinList, masterList)
	}
}
