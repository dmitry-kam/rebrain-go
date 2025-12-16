package task

import (
	"errors"
	"fmt"
)

var ErrTaskNotFound = errors.New("task not found")

var ErrPositiveId = errors.New("use positive id")

type TaskError struct {
	TaskID int
	Cause  error
}

func (e *TaskError) Error() string {
	return fmt.Sprintf("task error (ID=%d): %s", e.TaskID, e.Cause)
}

func (e *TaskError) Unwrap() error {
	return e.Cause
}
