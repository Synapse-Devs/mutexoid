package uselesscomments

// UserService is a service for users // want "redundant comment that restates what is evident from the code"
type UserService struct{}

// GetUser gets the user // want "redundant comment that restates function signature or obvious behavior"
func (s *UserService) GetUser(id string) string {
	return id
}

// ServiceInterface extends the base interface // want "redundant comment that restates what is evident from the code"
type ServiceInterface interface {
	// GetData gets the data // want "redundant comment that restates method signature"
	GetData() string
}

// ValidComment explains complex business logic that is not evident from the code
type ValidService struct{}

// ProcessData performs data validation, transformation and persistence according to business rules
func (s *ValidService) ProcessData(data string) error {
	return nil
}
