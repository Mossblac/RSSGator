package ext

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Mossblac/RSSGator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username expected")
	}

	name := cmd.Args[0]
	newId := uuid.New()
	currentTime := time.Now()

	argsCUP := database.CreateUserParams{
		ID:        newId,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      name,
	}

	UD, err := s.DataBase.CreateUser(context.Background(), argsCUP)
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			switch pgErr.Code.Name() {
			case "unique_violation":
				return fmt.Errorf("username already exists: %v", pgErr)
			default:
				return fmt.Errorf("create user error: %v", pgErr)

			}
		}
	}
	s.Config.CurrentUserName = name

	err = s.Config.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("user created:\nID:%v\nCreatedat:%v\nUpdatedat:%v\nUser Name:%v\n", UD.ID, UD.CreatedAt, UD.UpdatedAt, UD.Name)

	return nil
}

/*type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}*/

/*func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error)*/
