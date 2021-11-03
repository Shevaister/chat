package social

import (
	"go/pkg/db/connection"
	"go/pkg/repositories/social"
)

var Repository social.SocialRepository

func init() {
	Repository = social.ProvideSocialRepostiory(connection.DB)
}
