package references

// Operations represents operations related to references.
type Operations interface {
	DataOperations
}

// Service is responsible for managing references.
type Service struct {
	DataOperations
}

// NewService instantiates a new Service.
func NewService(database DataOperations) Operations {
	return &Service{DataOperations: database}
}
