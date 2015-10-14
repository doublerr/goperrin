package goperrin

import "testing"

func TestPerrin(t *testing.T) {
	var perrinList [4]int
	masterList := [4]int{3, 0, 2, 3}

	p := Perrin()

	index := 0
	for i := 3; i < 4; i = p() {
		if index == 4 {
			break
		}
		perrinList[index] = i
		index += 1
	}

	if perrinList != masterList {
		t.Errorf("perrinList == %v, want %v", perrinList, masterList)
	}
}

func TestPerrinMax(t *testing.T) {
	var perrinList [6]int
	max := 5

	p := PerrinMax(max)

	index := 0
	for i := 3; i < 6; i = p() {
		if index == 6 {
			break
		}
		perrinList[index] = i
		index += 1
	}

	if perrinList[len(perrinList)-1] != max {
		t.Errorf("last element of perrinList[] != %v", max)
	}
}

func TestPerrinReset(t *testing.T) {
	var perrinList [9]int
	masterList := [9]int{3, 0, 2, 3, 2, 5, 3, 0, 2}

	p := PerrinReset(5)

	index := 0
	for i := 3; i < 9; i = p() {
		if index == 9 {
			break
		}
		perrinList[index] = i
		index += 1
	}

	if perrinList != masterList {
		t.Errorf("perrinList == %v, want %v", perrinList, masterList)
	}
}
