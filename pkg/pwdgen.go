package pkg

import (
	"bytes"
	"fmt"
	 "github.com/ultima-liu/password-generator/test"
)

type PasswordGenerator struct {
	
}

func New() *PasswordGenerator{
	fmt.Printf("sdsd", 1111)
	return &PasswordGenerator{}
}

func (pg *PasswordGenerator) DigitOnly(length int) string{
	var buf bytes.Buffer
	digits := "0123456789"
	for l := 0; l < length; l++ {
		buf.WriteByte(digits[test.RandNumber(10)])
	}
	return buf.String()
}
func (pg *PasswordGenerator) CharOnly(length int) string{
	var buf bytes.Buffer
	digits := "abcdefghijklmnopqrstuvwxyz"
	for l := 0; l < length; l++ {
		buf.WriteByte(digits[test.RandNumber(26)])
	}
	return buf.String()
}