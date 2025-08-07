package ext

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {

	err := s.DataBase.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting users table: %v", err)
	}

	fmt.Printf("reset successful\n")
	return nil
}
