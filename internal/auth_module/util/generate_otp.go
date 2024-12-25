package util

import (
	"crypto/rand"
	"english_app/pkg/errs"
	"math/big"
)

func GenerateOTP() (string, errs.MessageErr) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	ret := make([]byte, 6)

	for i := 0; i < 6; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", errs.NewBadRequest("error generate otp")
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil

}
