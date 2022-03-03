package todo

type Service interface {
	GetAll() ([]Todo, error)
	GetById(id int) (Todo, error)
	AddTodo(request Request) (Todo, error)
	UpdateTodo(id int, request Request) (Todo, error)
	ToggleIsCompleteTodo(id int) (Todo, error)
	DeleteTodo(id int) (Todo, error)
}

type service struct {
	repository Repository
}

func CreateNewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAll() ([]Todo, error) {
	return s.repository.GetAll()
}

func (s *service) GetById(id int) (Todo, error) {
	todo, err := s.repository.GetById(id)

	return todo, err
}

func (s *service) AddTodo(request Request) (Todo, error) {
	var todo Todo
	todo.Title = request.Title
	return s.repository.AddTodo(todo)
}

func (s *service) UpdateTodo(id int, request Request) (Todo, error) {
	todo, err := s.repository.GetById(id)
	if err != nil {
		return todo, err
	}

	todo.Title = request.Title

	return s.repository.UpdateTodo(todo)
}

func (s *service) ToggleIsCompleteTodo(id int) (Todo, error) {
	todo, err := s.repository.GetById(id)
	if err != nil {
		return todo, err
	}

	todo.IsCompleted = !todo.IsCompleted

	return s.repository.UpdateTodo(todo)
}

func (s *service) DeleteTodo(id int) (Todo, error) {
	todo, err := s.repository.GetById(id)
	if err != nil {
		return todo, err
	}

	return s.repository.DeleteTodo(todo)
}
