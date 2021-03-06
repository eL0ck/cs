package processor

import (
	"math/rand"
	"testing"
)

var leftCases = map[string][]int{
	" aaaa":        {4, 0},
	"a aaa":        {4, 1},
	"aaaa ":        {4, 4},
	" 12345678901": {11, 11},
}

func TestFindNearbySpaceLeft(t *testing.T) {
	for k, v := range leftCases {
		space, _ := findNearbySpace(&fileJob{Content: []byte(k)}, v[0], 10)

		if space != v[1] {
			t.Error("Expected", v[1], "got", space)
		}
	}
}

func TestFindNearbySpaceLeftChangedTrue(t *testing.T) {
	space, changed := findNearbySpace(&fileJob{Content: []byte(` aaaa`)}, 4, 10)

	if space != 0 || changed != true {
		t.Error("Expected 0 and changed true got", space, changed)
	}
}

func TestFindNearbySpaceLeftChangedFalse(t *testing.T) {
	space, changed := findNearbySpace(&fileJob{Content: []byte(` aaaa`)}, 4, 1)

	if space != 4 || changed != false {
		t.Error("Expected 4 and changed true got", space, changed)
	}
}

var rightCases = map[string][]int{
	"aaaa ":        {0, 4},
	"aaa a":        {0, 3},
	" aaaa":        {0, 0},
	"12345678901 ": {0, 0},
}

func TestFindNearbySpaceRight(t *testing.T) {
	for k, v := range rightCases {
		space, _ := findNearbySpace(&fileJob{Content: []byte(k)}, v[0], 10)

		if space != v[1] {
			t.Error("Expected", v[1], "got", space)
		}
	}
}

func TestFindNearbySpaceRightChangedTrue(t *testing.T) {
	space, changed := findNearbySpace(&fileJob{Content: []byte(`aaaa `)}, 0, 10)

	if space != 4 || changed != true {
		t.Error("Expected 4 and changed true got", space, changed)
	}
}

func TestFindNearbySpaceRightChangedFalse(t *testing.T) {
	space, changed := findNearbySpace(&fileJob{Content: []byte(`aaaa `)}, 0, 1)

	if space != 0 || changed != false {
		t.Error("Expected 0 and changed true got", space, changed)
	}
}

// Try to find bugs by fuzzing the input to all sorts of random things
func TestFindNearbySpaceFuzzy(t *testing.T) {
	for i := 0; i < 100000; i++ {
		findNearbySpace(&fileJob{Content: []byte(randStringBytes(1000))}, rand.Intn(1000), rand.Intn(10000))
	}
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890~!@#$%^&*()_+{}|:<>?                        "

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
