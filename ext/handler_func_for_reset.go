package ext

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {

	err := s.DataBase.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting users table: %v", err)
	}

	err = s.DataBase.ResetFeedFollowsSequence(context.Background())
	if err != nil {
		return err
	}

	err = s.DataBase.ResetFeedsSequence(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("reset successful\n")
	return nil
}
