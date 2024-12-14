package repository

import (
	"findsafe/account/models/interfaces"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type RepositoryInt struct {
	CertRepositoryInt     interfaces.CertRepository
	OrgRepositoryInt      interfaces.OrgRepository
	ResourceRepositoryInt interfaces.ResourceRepository
	SearchRepositoryInt   interfaces.SearchRepository
	TeamRepositoryInt     interfaces.TeamRepository
	UserRepositoryInt     interfaces.UserRepository
}
