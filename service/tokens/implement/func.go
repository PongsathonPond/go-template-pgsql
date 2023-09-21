package implement

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"idev-cms-service/domain"

	"github.com/dgrijalva/jwt-go"
)

func (impl *implementation) generateAccessToken(data *domain.Users) (token *string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":        data.ID,
		"first_name": data.FirstName,
		"last_name":  data.LastName,
		"email":      data.Email,
		"exp":        time.Now().Add(time.Minute * time.Duration(impl.Config.AccessTokenTTL)).Unix(),
	})
	strToken, err := jwtToken.SignedString([]byte(impl.Config.AccessTokenSigningKey))
	if err != nil {
		return nil, err
	}
	return &strToken, nil
}

func (impl *implementation) md5(str string) string {
	h := md5.New()
	_, err := io.WriteString(h, str)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
