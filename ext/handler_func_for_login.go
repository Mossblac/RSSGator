package ext

import (
	"context"
	"database/sql"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username expected")
	}

	_, err := s.DataBase.GetUser(context.Background(), cmd.Args[0])
	if err == sql.ErrNoRows {
		return fmt.Errorf("username not registered")
	}

	s.Config.CurrentUserName = cmd.Args[0]

	err = s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User: %v has been set\n", cmd.Args[0])

	return nil
}

//func (q *Queries) GetUser(ctx context.Context, name string) (User, error)
