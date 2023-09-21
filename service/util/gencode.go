package util

import (
	"crypto/rand"
	"fmt"
	"io"
)

//go:generate mockery --name=GenCode
type GenCode interface {
	Generate(codeType GenCodeType, length int64) (code string)
	GenerateWithPrefix(codeType GenCodeType, length int64, prefix string) (code string)
}

type GenCodeType string

const (
	OTPCode   GenCodeType = "otpCode"
	RefCode   GenCodeType = "refCode"
	TokenCode GenCodeType = "tokenCode"
)

var charNum = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
var charUpper = [...]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var charLower = [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

type genCode struct{}

func NewGenCode() GenCode {
	return &genCode{}
}

func (gc *genCode) Generate(codeType GenCodeType, length int64) (code string) {
	switch codeType {
	case OTPCode:
		return gc.generateOTPCode(int(length))
	case RefCode:
		return gc.generateRefCode(int(length))
	case TokenCode:
		return gc.generateTokenCode(int(length))
	default:
		return ""
	}
}

func (gc *genCode) GenerateWithPrefix(codeType GenCodeType, length int64, prefix string) (code string) {
	switch codeType {
	case OTPCode:
		return fmt.Sprintf("%s-%s", prefix, gc.generateOTPCode(int(length)))
	case RefCode:
		return fmt.Sprintf("%s-%s", prefix, gc.generateRefCode(int(length)))
	case TokenCode:
		return fmt.Sprintf("%s%s", prefix, gc.generateTokenCode(int(length)))
	default:
		return ""
	}
}

func (gc *genCode) generateOTPCode(max int) string {
	b := make([]byte, max)
	n, _ := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return ""
	}
	for i := 0; i < len(b); i++ {
		b[i] = charNum[int(b[i])%len(charNum)]
	}
	return string(b)
}

func (gc *genCode) generateRefCode(max int) string {
	b := make([]byte, max)
	n, _ := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return ""
	}
	for i := 0; i < len(b); i++ {
		b[i] = charUpper[int(b[i])%len(charUpper)]
	}
	return string(b)
}

func (gc *genCode) generateTokenCode(max int) string {
	chars := append(append(charNum[:], charUpper[:]...), charLower[:]...)

	b := make([]byte, max)
	n, _ := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return ""
	}
	for i := 0; i < len(b); i++ {
		b[i] = chars[int(b[i])%len(chars)]
	}
	return string(b)
}
