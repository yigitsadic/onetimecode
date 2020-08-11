package shared

import (
	"math/rand"
	"time"
)

func CreateRandomValue() string {
	var built []byte
	var source []byte
	for x := byte('A'); x <= byte('Z'); x++ {
		source = append(source, x)
	}

	for x := byte('0'); x <= byte('9'); x++ {
		source = append(source, x)
	}

	for x := 0; x < 7; x++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		built = append(built, source[r1.Intn(len(source)-1)])
	}

	return string(built)
}
