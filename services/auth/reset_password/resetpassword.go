package auth

import (
	"errors"
	"fmt"

	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	util "github.com/greetinc/greet-util/s"
)

func (s *verifyService) ResetPassword(req dto.Reset) error {

	encryp := util.EncryptPasswordAfterReset(&req)
	if encryp != nil {
		return errors.New("Invalid password")
	}

	// Validate token and get user ID
	userID, err := s.Repo.ValidatePasswordResetToken(req.Token)
	if err != nil {
		return errors.New("Invalid or expired token")
	}
	fmt.Println("Received Token:", req.Token)

	// Update user password
	err = s.Repo.UpdateUserPassword(userID, req.NewPassword)
	if err != nil {
		return errors.New("Error updating password")
	}

	return nil
}
