package memberships

import (
	"api-music/internal/models/memberships"
	"api-music/pkg/jwt"
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) Login(request memberships.LoginRequest) (string, error) {
	userDetail, err := s.repository.GetUser(request.Email, "", 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error to get user from database")
		return "", err
	}
	if userDetail == nil {
		return "", errors.New("email is not exist")
	}
	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(request.Password))

	if err != nil {
		fmt.Println("test case password not match")
		return "", errors.New("email and password not mathced")
	}
	accessToken, err := jwt.CreateToken(userDetail.ID, userDetail.Username, s.cfg.Service.SecretJWT)

	if err != nil {
		log.Error().Err(err).Msg("failed to create jwt token")
		return "", err
	}

	return accessToken, nil

}
