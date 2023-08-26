package util

import (
	"crypto/rand"
	"math/big"
)

func RandInt64(min, max int64) (int64, error) {
	// calculate the max we will be using
	bg := big.NewInt(max - min)

	// get big.Int between 0 and bg
	// in this case 0 to 20
	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		return 0, err
	}

	// add n to min to support the passed in range
	return n.Int64() + min, nil
}
