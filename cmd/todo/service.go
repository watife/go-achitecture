package todo

// Service interface
type Service interface {
	CreateTodo(todo *Model) (*Model, error)
	FindTodoByID(id string) (*Model, error)
	FindAllTodos() ([]*Model, error)
}

type service struct {
	repo Repository
}

// NewService returns todo struct
func NewService(repo Repository) Service {
	return &service{
		repo,
	}
}

func (s *service) CreateTodo(todo *Model) (*Model, error) {
	todo.Status = "open"
	return s.repo.Create(todo)
}

func (s *service) FindTodoByID(id string) (*Model, error) {
	return s.repo.FindByID(id)
}

func (s *service) FindAllTodos() ([]*Model, error) {
	return s.repo.FindAll()
}
