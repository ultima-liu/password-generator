package pkg

import (
	"bytes"
	"github.com/ultima-liu/password-generator/internal"
)

type PasswordGenerator struct {
	
}

func New() *PasswordGenerator{
	return &PasswordGenerator{}
}

func (pg *PasswordGenerator) DigitOnly(length int) string{
	var buf bytes.Buffer
	digits := "0123456789"
	for l := 0; l < length; l++ {

		buf.WriteByte(digits[internal.RandNumber(10)])
	}
	return buf.String()
}
func (pg *PasswordGenerator) CharOnly(length int) string{
	var buf bytes.Buffer
	digits := "abcdefghijklmnopqrstuvwxyz"
	for l := 0; l < length; l++ {
		buf.WriteByte(digits[internal.RandNumber(26)])
	}
	return buf.String()
}