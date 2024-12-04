package task_1

type Employee struct {
	Name     string
	Position string
	Tasks    string
}

func (e *Employee) AddTask(task string) string {
	e.Tasks += task
	return e.Tasks
}

func (e *Employee) CompleteTasks() *Employee {
	return &Employee{
		Name:     e.Name,
		Position: e.Position,
		Tasks:    e.Tasks,
	}
}
