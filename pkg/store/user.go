package store

type User interface {
	// Save user
	Save()
	// Get User
	Get()
}

type user struct {
}

func NewStore() User {
	return &user{}
}

func (s *user) Save() {}
func (s *user) Get()  {}
