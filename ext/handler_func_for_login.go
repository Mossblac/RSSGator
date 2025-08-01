package ext

import (
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username expected")
	}

	s.Config.CurrentUserName = cmd.Args[0]

	err := s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User: %v has been set", cmd.Args[0])

	return nil
}
