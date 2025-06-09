package todo

import "time"

type TodoList []Todo

type Todo struct {
	priority    string
	description string
	createdOn   time.Time
	completedOn time.Time
}

func New(description string) *Todo {
	return &Todo{
		description: description,
		createdOn:   time.Now(),
	}
}

func (t *Todo) IsComplete() bool {
	return false
}

func (t *Todo) MarkComplete() {

}

func (t *Todo) MarkCompleteOn(completedOn time.Time) {

}

func (t *Todo) ContextTags() {

}

func (t *Todo) ProjectTags() {

}
