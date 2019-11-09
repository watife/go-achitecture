package todo

// Repository defines actions on todo package
type Repository interface {
	Create(todo *Model) (*Model, error)
	FindByID(id string) (*Model, error)
	FindAll() ([]*Model, error)
	// Update(id string, todo *Model) (*Model, error)
	// Delete(id string) error
}
