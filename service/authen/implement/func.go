package implement

import (
	"time"

	"idev-cms-service/domain"

	"github.com/dgrijalva/jwt-go"
)

func (impl *implementation) generateTokens(data *domain.Users) (accessToken, refreshToken *string, err error) {
	accessToken, err = impl.generateAccessToken(data)
	if err != nil {
		return nil, nil, err
	}
	refreshToken, err = impl.generateRefreshToken(data)
	if err != nil {
		return nil, nil, err
	}
	return accessToken, refreshToken, nil
}

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

func (impl *implementation) generateRefreshToken(data *domain.Users) (token *string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": data.ID,
		"exp": time.Now().Add(time.Minute * time.Duration(impl.Config.RefreshTokenTTL)).Unix(),
	})
	strToken, err := jwtToken.SignedString([]byte(impl.Config.RefreshTokenSigningKey))
	if err != nil {
		return nil, err
	}
	return &strToken, nil
}
