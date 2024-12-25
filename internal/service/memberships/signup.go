package memberships

import (
	"api-music/internal/models/memberships"
	"errors"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(request memberships.SignUpRequest) error {
	existingUser, err := s.repository.GetUser(request.Email, request.Username, 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("erros get user from database")

		return err
	}
	if existingUser != nil {
		return errors.New("email or username exits")
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("erros hash password")
		return err
	}
	model := memberships.User{
		Email:     request.Email,
		Username:  request.Username,
		Password:  string(pass),
		CreatedBy: request.Email,
		UpdatedBy: request.Email,
	}

	return s.repository.CreateUser(model)
}
