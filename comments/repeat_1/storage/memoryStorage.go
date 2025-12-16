package storage

import (
	. "repeat/task"
)

type MemoryStorage struct {
	tasks map[int]*Task
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{tasks: make(map[int]*Task)}
}

func (m *MemoryStorage) Save(task *Task) error {
	if task.ID < 0 {
		panic("invalid task ID")
	}
	m.tasks[task.ID] = task
	return nil
}

func (m *MemoryStorage) GetByID(id int) (*Task, error) {
	task, ok := m.tasks[id]
	if !ok {
		return nil, &TaskError{
			TaskID: id,
			Cause:  ErrTaskNotFound,
		}
	}
	return task, nil
}
