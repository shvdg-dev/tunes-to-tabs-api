package environments

import (
	tstdb "github.com/shvdg-dev/base-logic/pkg/testable/database"
	inl "github.com/shvdg-dev/tunes-to-tabs-api/internal"
)

// EnvManagement represents the operations related to managing environments.
type EnvManagement interface {
	CreatePostgresEnv() (DbEnvOperations, error)
	CreateDbEnv(config *tstdb.ContainerConfig) (DbEnvOperations, error)
}

// Service is responsible for managing different environments required for integration testing.
type Service struct {
	EnvManagement
	database tstdb.ContainerManagement
}

// NewService instantiates a new Service.
func NewService() EnvManagement {
	return &Service{
		database: tstdb.NewContainerService(),
	}
}

// CreatePostgresEnv creates a database environment for Postgres, with default configurations.
func (s *Service) CreatePostgresEnv() (DbEnvOperations, error) {
	return s.CreateDbEnv(tstdb.NewPostgresContainerConfig())
}

// CreateDbEnv creates an database environment.
func (s *Service) CreateDbEnv(config *tstdb.ContainerConfig) (DbEnvOperations, error) {
	dbContainer, err := s.database.CreateContainer(config)
	if err != nil {
		return nil, err
	}

	tablesService := inl.NewTableService(dbContainer)
	creatorService := inl.NewCreateService(tablesService)
	dropService := inl.NewDropService(tablesService)
	dbEnv := NewDbEnv(dbContainer, creatorService, dropService)

	return dbEnv, nil
}
