package ext

import (
	"context"

	"github.com/Mossblac/RSSGator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.DataBase.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}

}
