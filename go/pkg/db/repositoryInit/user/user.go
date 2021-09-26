package user

import (
	"go/pkg/db/connection"
	"go/pkg/models/user"
)

var Repository user.UserRepository

func init() {
	Repository = user.ProvideUserRepostiory(connection.DB)
}
