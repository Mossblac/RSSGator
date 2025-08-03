package ext

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {
	users, err := s.DataBase.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error obtaining users list")
	}

	for _, user := range users {
		if user == s.Config.CurrentUserName {
			fmt.Printf("%v  (current)\n", user)
		} else {
			fmt.Printf("%v\n", user)
		}
	}

	return nil

}
