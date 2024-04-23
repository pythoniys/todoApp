package service

import "github.com/pythoniys/todoApp/pkg/repository"

type Authorisation interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorisation
	TodoList
	TodoItem
}

func NewService(repo *repository.Repo) *Service {
	return &Service{}
}
