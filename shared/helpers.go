package shared

import (
	"math/rand"
	"time"
)

var (
	source = []byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I',
		'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R',
		'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
)

func CreateRandomValue() string {
	var built []byte

	for x := 0; x < 7; x++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		built = append(built, source[r1.Intn(len(source)-1)])
	}

	return string(built)
}
