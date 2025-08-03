package ext

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	err := s.DataBase.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting database")
	}
	fmt.Printf("users reset successfully\n")
	return nil
}
