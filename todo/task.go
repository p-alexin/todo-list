package todo

import "time"

type Task struct {
	Title     string
	Text      string
	Status    bool
	Createdat time.Time
	DoneAt    *time.Time
}

func NewTask(title, text string) Task {
	return Task{
		Title:     title,
		Text:      text,
		Status:    false,
		Createdat: time.Now(),
		DoneAt:    nil,
	}
}
