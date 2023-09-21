package util

import (
	"crypto/md5"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

//go:generate mockery --name=Encrypt
type Encrypt interface {
	MD5(str string) string
	HashPassword(password string) string
	CheckPasswordHash(password, hash string) bool
}

type encrypt struct{}

func NewEncrypt() (e Encrypt) {
	return &encrypt{}
}

func (e *encrypt) MD5(str string) string {
	h := md5.New()
	_, err := io.WriteString(h, str)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (e *encrypt) HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (e *encrypt) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
