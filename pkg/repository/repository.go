package repository

type Authorisation interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repo struct {
	Authorisation
	TodoList
	TodoItem
}

func NewRepo() *Repo {
	return &Repo{}
}
