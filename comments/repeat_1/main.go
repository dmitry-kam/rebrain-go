package main

import (
	"errors"
	"fmt"
	. "repeat/mixed"
	. "repeat/storage"
	. "repeat/task"
)

func PrintValue(value Mixed) {
	switch v := value.(type) {
	case *Task:
		fmt.Printf("Task: ID=%d, Title=%s, Completed=%t\n",
			v.ID, v.Title, v.IsCompleted())
	case int:
		fmt.Printf("Int: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Bool: %t\n", v)
	default:
		fmt.Printf("Unknown: %T\n", v)
	}
}

func PrintAnyValue(value any) {
	switch v := value.(type) {
	case int:
		fmt.Printf("Int: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Bool: %t\n", v)
	default:
		fmt.Printf("Unknown: %T\n", v)
	}
}

func safeSave(s Storage, task *Task) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v: %w", r, ErrPositiveId)
		}
	}()
	return s.Save(task)
}

func main() {
	memoryStorage := NewMemoryStorage()

	task1 := &Task{Title: "First task", ID: 1}
	task1.Complete()
	task2 := &Task{Title: "Second task", ID: 2}
	task3 := &Task{Title: "Third task", ID: 3}
	task4 := &Task{Title: "Non-existent task", ID: -4}

	if err := safeSave(memoryStorage, task1); err != nil {
		fmt.Println(err)
	}
	if err := safeSave(memoryStorage, task2); err != nil {
		fmt.Println(err)
	}
	if err := safeSave(memoryStorage, task3); err != nil {
		fmt.Println(err)
	}
	if err := safeSave(memoryStorage, task4); err != nil {
		fmt.Println(err)
	}

	foundedTask, err := memoryStorage.GetByID(14)
	if errors.Is(err, ErrTaskNotFound) {
		fmt.Printf("TaskID NotFound ошибка %v\n", err)
	}

	if err != nil {
		var taskErr *TaskError
		if errors.As(err, &taskErr) {
			fmt.Printf("Ошибка с задачей ID=%d: %v\n", taskErr.TaskID, taskErr.Cause)
		}
	}

	foundedTask, err = memoryStorage.GetByID(3)
	if err != nil {
		if errors.Is(err, ErrTaskNotFound) {
			fmt.Printf("TaskID NotFound ошибка %v\n", err)
		}
		return
	}
	fmt.Println("#######################")
	PrintValue(foundedTask.ID)
	PrintValue(foundedTask.Title)
	PrintAnyValue(foundedTask.IsCompleted())
	fmt.Println()
	PrintValue(foundedTask)
	fmt.Println("#######################")
	fmt.Println()

	///////////////////////////////////////////////

	fileStorage := NewFileStorage("tasks")

	task3.Complete()
	if err := safeSave(fileStorage, task1); err != nil {
		fmt.Println(err)
	}
	if err := safeSave(fileStorage, task2); err != nil {
		fmt.Println(err)
	}
	if err := safeSave(fileStorage, task3); err != nil {
		fmt.Println(err)
	}

	foundedTask, err = fileStorage.GetByID(18)
	if errors.Is(err, ErrTaskNotFound) {
		fmt.Printf("TaskID NotFound ошибка %v\n", err)
	}

	foundedTask, err = memoryStorage.GetByID(2)
	if err != nil {
		if errors.Is(err, ErrTaskNotFound) {
			fmt.Printf("TaskID NotFound ошибка %v\n", err)
		}
		return
	}

	fmt.Println("#######################")
	PrintValue(foundedTask.ID)
	PrintAnyValue(foundedTask.Title)
	PrintValue(foundedTask.IsCompleted())
	fmt.Println("#######################")
	fmt.Println()

	//////////////////////////////////////

	foundedTask, err = memoryStorage.GetByID(3)
	if err != nil {
		if errors.Is(err, ErrTaskNotFound) {
			fmt.Printf("TaskID NotFound ошибка %v\n", err)
		}
		return
	}
	fmt.Println("#######################")
	PrintAnyValue(foundedTask.ID)
	PrintValue(foundedTask.Title)
	PrintValue(foundedTask.IsCompleted())
	fmt.Println("#######################")
}
