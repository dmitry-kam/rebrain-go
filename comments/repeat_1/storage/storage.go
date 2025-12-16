package storage

import . "repeat/task"

type Storage interface {
	Save(task *Task) error
	GetByID(id int) (*Task, error)
}
