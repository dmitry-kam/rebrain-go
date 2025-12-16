package task

type Task struct {
	ID        int
	Title     string
	completed bool
}

func (t *Task) Complete() {
	t.completed = true
}

func (t *Task) IsCompleted() bool {
	return t.completed
}
