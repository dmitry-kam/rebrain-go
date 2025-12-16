package storage

import (
	"fmt"
	"os"
	"path/filepath"
	. "repeat/task"
)

const txtExt = ".txt"

type FileStorage struct {
	*MemoryStorage
	fileName string
}

func NewFileStorage(name string) *FileStorage {
	return &FileStorage{
		MemoryStorage: NewMemoryStorage(),
		fileName:      name,
	}
}

func (m *FileStorage) GetFileName() string {
	return m.fileName + txtExt
}

func (m *FileStorage) Save(task *Task) error {
	if task.ID < 0 {
		panic("invalid task ID")
	}
	m.tasks[task.ID] = task
	if err := m.writeToFile(); err != nil {
		fmt.Println("failed to save in file!")
		return err
	}

	return nil
}

func (m *FileStorage) writeToFile() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	wFilePath := filepath.Join(currentDir, m.GetFileName())
	fW, err := os.OpenFile(wFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer closeFile(fW)

	for _, task := range m.tasks {
		formattedLine := fmt.Sprintf("%d %s %t\n", task.ID, task.Title, task.IsCompleted())
		_, err := fW.WriteString(formattedLine)
		if err != nil {
			return err
		}
	}
	return nil
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		fmt.Println(err)
		return
	}
}
