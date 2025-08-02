package ext

import (
	"github.com/Mossblac/RSSGator/internal/config"
	"github.com/Mossblac/RSSGator/internal/database"
)

type State struct {
	DataBase *database.Queries
	Config   *config.Config
}
