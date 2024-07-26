package tabs

// Operations represent all operations related to tabs.
type Operations interface {
	DatabaseOperations
}

// Service is responsible for managing and retrieving tabs.
type Service struct {
	DatabaseOperations
}

// NewService instantiates a Service.
func NewService(database DatabaseOperations) Operations {
	return &Service{DatabaseOperations: database}
}