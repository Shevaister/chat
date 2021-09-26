package social

import (
	"go/pkg/db/connection"
	"go/pkg/models/social"
)

var Repository social.SocialRepository

func init() {
	Repository = social.ProvideSocialRepostiory(connection.DB)
}
