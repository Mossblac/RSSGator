package ext

import "fmt"

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("handler expects a single argument with a username")
	}

	s.Config.CurrentUserName = cmd.Args[0]

	fmt.Printf("User: %v has been set", cmd.Args[0])

	return nil
}
